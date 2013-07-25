$(function() {
    console.log('Yes');
    var data = {
        "Id": "3",
        "Row": "1",
        "Col": "1",
        "Sizex": "1",
        "Sizey": "1",
        "Content": "Hola"
    };

    $('#add-block').click(function() {
        $.ajax({
            url: "/",
            contentType: "application/json",
            data: JSON.stringify(data),
            method: "POST",
            dataType: "json"
        });
    });
});
