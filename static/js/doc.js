$(document).ready(function(){  
  $("#back").click(function() {
      window.history.back();
  });

  $("#add").click(function() {
    var file = $("#file").val();
    $("#add").prop('disabled', true);
    $("#load").prop('disabled', true);
    add(file);
});   

  $("#load").click(function() {
      var file = $("#file").val();
      $("#add").prop('disabled', true);
      $("#load").prop('disabled', true); 
      upload(file);
  });   

  $("#clear").click(function() {
    clear();
  });   

  $(window).keypress(function(event) {
    if (event.metaKey) {
        switch(event.key) {
            case 'c':
                copySelectedText();
                break;
            case 'v':
                pasteFromClipboard($("#file"));
                break;
            case 'a':
                $("#file").select();
                break;
            case 'q':
                closeWindow();
                break;
        }
        event.preventDefault();
    }                        
  });  
});  

async function upload(file) {
  $.ajax({
      url: "/document/load",
      type: "POST",
      contentType: "application/json",
      data: JSON.stringify({file: file}),
      success: function(response) {
          $("#content").html(response);
          $("#add").prop('disabled', false);
          $("#load").prop('disabled', false);
      },
      error: function(err) {
          console.error('Failed to upload file: ', err);
      }
  });           
}

async function add(file) {
    $.ajax({
        url: "/document/add",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({file: file}),
        success: function(response) {
            $("#content").html(response);
            $("#add").prop('disabled', false);
            $("#load").prop('disabled', false);            
        },
        error: function(err) {
            console.error('Failed to upload file: ', err);
        }
    });           
  }
  

async function pasteFromClipboard($textarea) {
  try {
      const text = await navigator.clipboard.readText();
      let cursorPosition = $textarea[0].selectionStart;
      let currentVal = $textarea.val();
      $textarea.val(currentVal.substring(0, cursorPosition) + text + currentVal.substring(cursorPosition));
      $textarea[0].selectionStart = cursorPosition + text.length;
      $textarea[0].selectionEnd = cursorPosition + text.length;      
  } catch (err) {
      console.error('Failed to read clipboard contents: ', err);
  }
}

function clear() {
  $.get("/document/clear", function() {
      $("#content").text("");
  }).fail(function(err) {
      console.error('Failed to clear content: ', err);
  });        
}
