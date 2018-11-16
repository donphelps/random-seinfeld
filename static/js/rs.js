
$(document).ready(function () {	
	$("#rnd-any").click(function () { loadRandomEpisode(); });
	$("#rnd-jerry").click(function () { loadRandomEpisode("jerry"); });
	$("#rnd-george").click(function () { loadRandomEpisode("george"); });
	$("#rnd-kramer").click(function () { loadRandomEpisode("kramer"); });
	$("#rnd-elaine").click(function () { loadRandomEpisode("elaine"); });
	
	loadRandomEpisode();
});

function loadRandomEpisode(criteria) {
	var url = "/rnd/" + criteria;

	$.getJSON(url, function (data, textStatus, jqXHR) {
		var episode = data;
		$("#ep-title").text(episode.season + "." + episode.episode + ": " + episode.title);
		$("#ep-director").text("Directed by: " + episode.director);
		$("#ep-writers").text("Written by: " + episode.writers);
		$("#ep-desc").text(episode.description);
		$("#ep-hulu-btn").attr("href", episode.link);
	});
}
