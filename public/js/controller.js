$(function() {
    var gridster = $(".gridster ul").gridster({
        widget_margins: [10, 10],
        widget_base_dimensions: [140, 140]
    }).data('gridster');

    var data = {
        "Id": "X",
        "Row": "1",
        "Col": "1",
        "Sizex": "1",
        "Sizey": "1",
        "Content": "Default block"
    };

    $('#add-block').click(function() {
        $.ajax({
            url: "/",
            contentType: "application/json",
            data: JSON.stringify(data),
            method: "POST",
            dataType: "json"
        }).success(function(data) {
            if(data.Id) {
                gridster.add_widget('<li class="new">The HTML of the widget...</li>');
            }
            else {
                //TODO Show error
                console.log("Error!");
            }
        });
    });
});
