var converter = new showdown.Converter();

$(document).ready(function(){  
    $('#prompt').keypress(function(event){        
        var keycode = event.which || event.keyCode;
        if((keycode == 10 || keycode == 13) && event.ctrlKey){
            send(event);
            return false;
        }
    });       
    autosize($('#prompt'));    

    $("#run").click(function(event) {
        send(event);
    });

    $(".assistant, .model").click(function() {
        var changeType = $(this).hasClass("assistant") ? "assistant" : "model";
        var clickedElement = $(this); // Store the reference to the clicked element
        $.ajax({
            url: "/change/" + changeType,
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify({[changeType]: clickedElement.text()}),
            success: function() {
                $("#" + changeType + "_select").text(clickedElement.text()); // Use the stored reference
            }
        });
    });
    

    $("#document, #settings, #history, #clear, #image").click(function() {
        var action = $(this).attr('id');
        if(action === 'clear') {
            clear();
        } else {
            console.log('/' + action);
            window.location.href = '/' + action;
        }
    });

    $(window).keypress(function(event) {
        if (event.metaKey) {
            switch(event.key) {
                case 'c': copySelectedText(); break;
                case 'v': pasteFromClipboard($('#prompt')[0]); break;
                case 'a': $('#prompt').select(); break;
                case 'q': closeWindow(); break;
            }
            event.preventDefault();
        }
    });
});  

function send(e){
    e.preventDefault();
    var prompt = $("#prompt").val().trimEnd();
    $("#prompt").val("");
    autosize.update($("#prompt"));

    $("#printout").append(
        "<div class='prompt-message'>" + 
        "<div style='white-space: pre-wrap;'>" +
        encodeHtml(prompt) +
        "</div>" +
        "<span class='message-loader js-loading spinner-border'></span>" +
        "</div>"             
    );        

    $('html, body').animate({scrollTop: $(document).height()}, 'smooth');
    ask(prompt);          
    $(".js-logo").addClass("active");
}

async function ask(prompt, action="/ask") {    
    var id = Math.random().toString(36).substring(2,7);
    var outId = "result-" + id;
    
    $("#printout").append(
        "<div class='px-3 py-3'>" + 
        "<div id='" + outId +
        "' class='whitespace'>" +         
        "</div>" +
        "</div>" 
    );  
    // using fetch because $.ajax doesn't support streaming
    response = await fetch("/ask", {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({input: prompt}),
    });

    // process the streaming output from the HTTP call
    decoder = new TextDecoder();
    reader = response.body.getReader();
    let output = "";
    while (true) {
        const { done, value } = await reader.read();
        if (done) break;
        output = output + decoder.decode(value);
        $("#"+outId).text(output);
        $('.printout-wrapper').scrollTop($('.printout-wrapper')[0].scrollHeight);
    }

    // streaming ends, remove the spinner, make the text HTML and apply highlighting
    $(".js-loading").removeClass("spinner-border");        
    $("#" + outId).html(converter.makeHtml(($("#"+outId).text())));
    $('.printout-wrapper').scrollTop($('.printout-wrapper')[0].scrollHeight);
    hljs.highlightAll();
    $("#" + outId).removeClass('whitespace');
}

function encodeHtml(str) {
    return str.replace(/[&<>"']/g, function(char) {
      switch (char) {
        case '&': return '&amp;';
        case '<': return '&lt;';
        case '>': return '&gt;';
        case '"': return '&quot;';
        case "'": return '&apos;';
      }
    });
}

async function copySelectedText() {
    try {
      await navigator.clipboard.writeText(window.getSelection().toString());
      console.log('Text successfully copied to clipboard');
    } catch (err) {
      console.error('Failed to copy text: ', err);
    }
}

async function pasteFromClipboard(textarea) {
    try {
        const text = await navigator.clipboard.readText();
        let cursorPosition = textarea.selectionStart;
        $(textarea).val($(textarea).val().substring(0, cursorPosition) + text + $(textarea).val().substring(cursorPosition));
        textarea.selectionStart = cursorPosition + text.length;
        textarea.selectionEnd = textarea.selectionStart;      
    } catch (err) {
        console.error('Failed to read clipboard contents: ', err);
    }
}

function clear() {
    $.get("/document/clear", function() {
        $("#content").text("");
        $("#printout").text("");
    });
}
