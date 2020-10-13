import store from '@/store'

class Ws {
    ws = null;
    reconnectInterval = 1000;
    maxReconnectInterval = 3000;
    maxRetriesToConnect = 5;
    retriesToConnect = 0;

    constructor(options) {
        this.options = options;
    }

    connect() {
        this.retriesToConnect++;

        console.log(`connecting to ws ...`);

        let ws = new WebSocket(this.options.url);
        let reconnectInterval = this.options.reconnectInterval || this.reconnectInterval;

        ws.onopen = () => {
            // Restart reconnect interval
            reconnectInterval = this.options.reconnectInterval || this.reconnectInterval;
            store.dispatch(`ws/onOpenHandler`);
        };

        ws.onmessage = (event) => {
            store.dispatch(`ws/eventsHandler`, JSON.parse(event.data));
        };

        ws.onclose = (event) => {
            if (event) {
                console.log(`ws has gone away`, event);
                // Event.code 1000 is our normal close event
                if (event.code === 1000) {
                    return ;
                }

                if (this.retriesToConnect > this.maxRetriesToConnect) {
                    return ;
                }

                let maxReconnectInterval = this.options.maxReconnectInterval || this.maxReconnectInterval;
                setTimeout(() => {
                    if (reconnectInterval < maxReconnectInterval) {
                        // Reconnect interval can't be > x seconds
                        reconnectInterval += 2000
                    }
                    this.connect()
                }, reconnectInterval)
            }
        };

        ws.onerror = (error) => {
            console.log(`ws`, error);
            ws.close()
        };

        this.ws = ws
    }

    disconnect() {
        if (null === this.ws) {
            return
        }

        this.ws.onclose = function () {}; // disable onclose handler first
        this.ws.close();
        this.ws = null;
        this.retriesToConnect = 0;
    }

    send(data) {
        if (null === this.ws) {
            return
        }

        this.ws.send(JSON.stringify(data))
    }
}

export const ws = {
    install(Vue, options) {
        console.log(options);

        Vue.prototype.$ws = new Ws(options)
    }
};

