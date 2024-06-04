$(document).ready(function(){  

    $("#back").click(function() {
        window.location.href = '/';
    });

    $("#add").click(function() {
        var model = $("#modeltag").val();
        addModel(model);
    });    

    $("#remove").click(function() {
        var model = $("#modeltag").val();
        removeModel(model);
    });   

    $(window).keypress(function(event) {
        if (event.metaKey && event.key === 'c') {
          copySelectedText();
          event.preventDefault();
        }

        if (event.metaKey && event.key === 'q') {
            closeWindow();
            event.preventDefault();
        }                        
    });
});  

async function addModel(model) {
    response = await fetch("/models/pull", {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({model: model}),
    });           

    decoder = new TextDecoder();
    reader = response.body.getReader();
    while (true) {
        const { done, value } = await reader.read();
        if (done) break;
        $("#modeloutput").html(decoder.decode(value));
    }
}

async function removeModel(model) {
    response = await fetch("/models/remove", {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({model: model}),
    });           

    decoder = new TextDecoder();
    reader = response.body.getReader();
    while (true) {
        const { done, value } = await reader.read();
        if (done) break;
        $("#modeloutput").html(decoder.decode(value));
    }
}

async function copySelectedText() {
    try {
      const text = document.getSelection().toString();
      await navigator.clipboard.writeText(text);
      console.log('Text successfully copied to clipboard');
    } catch (err) {
      console.error('Failed to copy text: ', err);
    }
}
