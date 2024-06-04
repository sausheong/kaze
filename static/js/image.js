var converter = new showdown.Converter();

$(document).ready(function(){  
    $('#prompt').keypress(function(event){        
        var keycode = event.which || event.keyCode;
        if((keycode == 10 || keycode == 13) && event.ctrlKey){
            send($('#prompt').val());
            return false;
        }
    });         

    $("#run").click(function() {
        send($('#prompt').val());
    });

    $(".size, .model").click(function() {
        var changeType = $(this).hasClass("size") ? "size" : "model";
        var clickedElement = $(this); // Store the reference to the clicked element
        $.ajax({
            url: "/image/change/" + changeType,
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify({[changeType]: clickedElement.text()}),
            success: function() {
                $("#" + changeType + "_select").text(clickedElement.text()); // Use the stored reference
            }
        });
    });

    $("#back").click(function() {
        window.location.href="/"
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

function send(prompt){
    // add a  bootstrap spinner to button with id 'run'
    $('#runtxt').addClass('spinner-border spinner-border-sm');
    $('#runtxt').text("");
    $.ajax({
        url: "/image/texttoimage",
        type: "POST",
        data: JSON.stringify({prompt: prompt}),
        success: function(response) {
            $('#imgout').attr('src', 'data:image/png;base64,' + response);
            $('#runtxt').removeClass('spinner-border spinner-border-sm');
            $('#runtxt').text("▶ run");
        },
        error: function(err) {
            console.error('failed to generate image: ', err);
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
