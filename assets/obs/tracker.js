var app = new Vue({
    el: '#app',
    data: {
        interval: 30000,
        stats: {
            kills: 0,
            wins: 0,
            kdr: 0.0
        }
    },
    created: function () {
        this.start();
    },
    methods: {
        start: function () {
            var self = this;
            self.getData();
            setInterval(function () {
                self.getData();
            }, this.interval);
        },
        getData: function () {
            var self = this;
            var apiRoute = window.location.pathname.split('fortnite')[1]
            axios.get(window.location.origin + '/obs' + apiRoute).then(response => {
                self.stats = response.data;
            });
        }
    }
});