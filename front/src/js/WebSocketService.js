class WebSocketService {
    constructor() {
        this.websock = null;
        this.reconnectAttempts = 0;
        this.maxReconnectAttempts = 10;
        this.currentEmail = '';
        this.onOpen = this.onOpen.bind(this);
        this.onClose = this.onClose.bind(this);
        this.onError = this.onError.bind(this);
        this.onMessage = this.onMessage.bind(this);
        this.retryConnection = this.retryConnection.bind(this);
    }

    establishConnection(email) {
        this.currentEmail = email;
        if (this.reconnectAttempts >= this.maxReconnectAttempts) {
            console.log("达到最大重连次数，停止尝试。");
            return;
        }

        try {
            const wsuri = `ws://localhost:8090/web_link?email=${email}`;
            this.websock = new WebSocket(wsuri);

            this.websock.onmessage = this.onMessage;
            this.websock.onopen = this.onOpen;
            this.websock.onerror = this.onError;
            this.websock.onclose = this.onClose;
        } catch (e) {
            console.error("建立连接时发生错误: ", e);
            this.retryConnection();
        }
    }

    onOpen() {
        console.log("WebSocket 连接已打开");
        // 这里可以添加其他打开连接时的逻辑
    }

    onClose(e) {
        console.log("断开连接", e);
        this.retryConnection();
    }

    onError(error) {
        console.log("出现错误", error);
        this.retryConnection();
    }

    onMessage(e) {
        document.dispatchEvent(new CustomEvent('websocket-message', { detail: e.data }));
        console.log("收到消息: ", e.data);
        // 这里可以根据接收到的消息执行操作
    }

    retryConnection() {
        if (this.reconnectAttempts < this.maxReconnectAttempts) {
            this.reconnectAttempts++;
            setTimeout(() => this.establishConnection(this.currentEmail), 10000); // 10秒后重试
        } else {
            console.log("达到最大重连次数，停止尝试。");
        }
    }

    sendMessage(message) {
        if (this.websock) {
            this.websock.send(message);
        }
    }

    closeConnection() {
        if (this.websock) {
            this.websock.close();
        }
    }
}

export default new WebSocketService();
