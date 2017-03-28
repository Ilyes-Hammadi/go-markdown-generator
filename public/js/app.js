$(document).ready(function () {
	$("textarea").keyup(function () {
		console.log("Starting the request !!")
		$.getJSON("http://localhost:9000/markdown", {
				body: $("textarea").val()
			})
			.done(function (data) {
                console.log("SUCCESS")
                $("#markdown").html(data)
			})
            .fail(function (err) {
                console.log("ERROR")
            })
	})
})
