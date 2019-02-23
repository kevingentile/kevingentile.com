var app = new Vue({
    el: '#app',
    data: {
        interval: 60000,
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
            axios.get(window.location.pathname).then(response => {
                self.stats = response.data;
            });
        }
    }
});