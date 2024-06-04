var converter = new showdown.Converter();

$(document).ready(function(){  
    $(window).keypress(function(event) {
        if (event.metaKey) {
            switch(event.key) {
                case 'c':
                    copySelectedText();
                    break;
                case 'q':
                    closeWindow();
                    break;
            }
            event.preventDefault();
        }                        
    });
    $("#history").html(converter.makeHtml($("#history").text()));    
    hljs.highlightAll();    
    $('.history-wrapper').scrollTop($('.history-wrapper').prop("scrollHeight"));     
});  

async function copySelectedText() {
    const text = document.getSelection().toString();
    try {
      await navigator.clipboard.writeText(text);
    } catch (err) {
      console.error('Failed to copy text: ', err);
    }
}

$("#back").click(function() {
    window.history.back();
});
