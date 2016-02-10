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
            "SubsNoti-app"
        ]
    }
});

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
        }
        vue.status = data;
    }
})