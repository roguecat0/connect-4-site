columns = [0, 0, 0, 0, 0, 0, 0]
turn = 1
moves = ""

const GameState = {
    Draw: 0,
    Win: 1,
    Ongoing: 2
}
// response json
response = {
    "cpu_move": 0,
    "game_state_pre": GameState.Ongoing,
    "game_state_post": GameState.Win
}

$(document).ready(function () {
    $(".column").on("click", function () {
        play_move(this.id[1])
        $.ajax({
            url: "/move/"+ moves,
            success: function (response) {
                // $("#result").html(response)
                let res = JSON.parse(response)
                handel_response(res)
            },
        })
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
    moves = moves + eid
    console.log(moves)
    columns[eid]++
    turn++
}
function handel_response(response) {
    console.log(GameState.Ongoing)
    console.log(response)
    if (response["game_state_pre"] != GameState.Ongoing) {
        show_winner(response["game_state_pre"])
        return
    }

    play_move(response["cpu_move"])

    if (response["game_state_post"] != GameState.Ongoing) {
        show_winner(response["game_state_post"])
    }
}
function show_winner(state) {
    if (state === GameState.Draw) {
        $("#whosturn").text("Draw")
    } else if (turn % 2 != 0) {
        $("#whosturn").text("Yellow Wins!")
    } else {
        $("#whosturn").text("Red Wins!")
    }

}
