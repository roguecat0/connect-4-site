columns = [0, 0, 0, 0, 0, 0, 0]
turn = 1

$(document).ready(function () {
    $(".column").on("click", function () {
        play_move(this.id[1])
    })
})

function play_move(eid) {
    sum = columns[eid]
    if (sum > 5) {
        return
    }

    let column_row = `#c${eid}r${sum}`
    if (turn % 2 != 0) {
        $(column_row).css("background-color", "red")
        $("#whosturn").text("Red's turn")
    } else {
        $(column_row).css("background-color", "yellow")
        $("#whosturn").text("Yellow's turn")
    }
    columns[eid]++
    turn++
}