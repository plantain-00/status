var vue = new Vue({
    el: "#vue",
    data: {
        status: [],
        repositories: [
            "plantain-00.github.io",
            "news-fetcher-client",
            "deploy-robot",
            "SubsNoti-frontends",
            "SubsNoti-doc",
            "news-fetcher",
            "SubsNoti",
            "co-md"
        ]
    }
});

function getDuration(minutes) {
    var duration;
    if (minutes % 60 !== 0) {
        duration = minutes % 60 + "min";
    } else {
        duration = "";
    }
    var hours = Math.floor(minutes / 60);

    if (hours % 24 !== 0) {
        duration = hours % 24 + "h" + (duration ? " " + duration : duration);
    }
    var days = Math.floor(hours / 24);

    if (days !== 0) {
        duration = days + "d" + (duration ? " " + duration : duration);
    }
    return duration;
}

$.ajax({
    url: "/api/status",
    cache: false,
    success: function (data) {
        for (var i = 0; i < data.length; i++) {
            var rate = data[i].total === 0 ? 1 : (data[i].total - data[i].fail) / data[i].total;
            var first = Math.round(255 - rate * 255);
            var second = Math.round(rate * 255);
            var color = (first * 256 * 256 + second * 256).toString(16);
            if (color.length === 4) {
                color = "00" + color;
            } else if (color.length === 5) {
                color = "0" + color;
            }
            data[i].color = color;
            data[i].rate = (rate * 100).toFixed(4) + "%";

            data[i].duration = getDuration(data[i].total);
        }
        vue.status = data;
    }
})