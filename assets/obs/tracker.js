var app = new Vue({
    el: '#app',
    data: {
        timeout: 3 * 60 * 60,
        stats: {
            kills: 0,
            wins: 0,
            kdr: 0.0
        }
    },
    
    created: function(){
        this.start();
    },
    methods:{
        start: function(){
            this.getData();
            setTimeout(function(){
               app.getData();
            }, this.timeout);
        },
        getData: function(){
            axios.get(`/obs/laughingcabbage`)
            .then(response => {
                this.stats = response.data
            })
        }
    }
})