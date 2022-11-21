<script>
    import { pop } from "svelte-spa-router";
    let message = "";
    let messageList = [];
    // export let senderid;
    // export let receiverid;
    export let params = {};
    let chater = params.third;
    console.log(
        "params.first:",
        params.first,
        "params.second",
        params.second,
        "params.third",
        params.third
    );
    const ws = new WebSocket(
        `ws://localhost:9090/api/chat?sender=${params.first}&receiver=${params.second}`
    );
    ws.onopen = function (evt) {
        console.log("Connection open ...");
    };
    ws.onmessage = function (evt) {
        console.log("Received Message: " + evt.data);
        const data = JSON.parse(evt.data);
        messageList = [
            ...messageList,
            {
                msg: data.message,
                senderid: data.sender,
            },
        ];
        console.log("messageList:", messageList);
        console.log("messageList[0].msg", messageList[0].msg);
    };
    ws.onclose = function (evt) {
        console.log("Connection closed.");
    };

    function sendMessage() {
        console.log("发送的消息是：", message);
        console.log("params.first", params.first);
        messageList = [
            ...messageList,
            {
                msg: message,
                senderid: params.first,
            },
        ];
        let body = JSON.stringify({
            sender: params.first,
            receiver: params.second,
            message: message,
        });
        ws.send(body);
        message = "";
        console.log("messageList", messageList);
    }
</script>

<main>
    <div>
        <div class="header">
            <div class="logo">{chater}</div>
            <button
                on:click={() => {
                    pop();
                }}>{"<"}</button
            >
        </div>

        <div class="messages">
            {#each messageList as mes, i}
                {#if mes.senderid != params.first}
                    <div class="message other-message">
                        <div class="text">
                            {mes.msg}
                        </div>
                    </div>
                {:else}
                    <div class="message my-message">
                        <div class="text">
                            {mes.msg}
                        </div>
                    </div>
                {/if}
            {/each}
        </div>

        <div class="form">
            <input type="text" bind:value={message} />
            <button on:click={sendMessage}>
                <i class="fa fa-paper-plane" />
            </button>
        </div>
    </div>
</main>

<!-- {#each messageList as mes, i}
    {#if mes.senderid != undefined}
        <strong>{mes.senderid + ":"}</strong>
        {mes.msg}
    {:else}
        <strong>{"我" + ":"}</strong>
        {mes.msg}
    {/if}
    <br />
{/each}
{"聊天界面"}
<input type="text" bind:value={message} />
<button on:click={sendMessage}>发送</button> -->

<svelte:head>
    <style>
        body {
            background: #f5f5f5;
        }
        main {
            /* fixed：固定在一个地方，无论浏览器上拉下拉都固定在一个地方 */
            position: fixed;
            /* 以左border为准离外容器的距离，外容器的position 是relative 若是absolute则无效 */
            /* top:50% left 50% 使得左上角在页面中间 */
            left: 50%;
            top: 50%;
            /* transform(-50%,-50%) 向上向左移动自身长宽的50% 使得main在页面居中*/
            transform: translate(-50%, -50%);
            width: 100%;
            height: 100%;
            max-width: 400px;
            max-height: 500px;
            background: #fff;
            border: 1px solid #eee;
            box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.05);
        }
        .header {
            width: 95%;
            height: 50px;
            /* header内部子元素按行排列 */
            display: flex;
            justify-content: space-between;
            /* align—items：竖直方向居中 */
            align-items: center;
            padding: 0px 10px;
            border-bottom: 1px solid #ddd;
        }
        .header .logo {
            font-size: 15px;
            font-weight: 600;
            color: #111;
        }
        .header button {
            background: rgb(255, 255, 255);
            /* margin-right: 15px; */
            padding: 5px 10px;
            border: 0px solid #ddd;
            color: #555;
            font-size: 20px;
            cursor: pointer;
        }
        .form {
            position: absolute;
            display: flex;
            bottom: 0px;
            width: 100%;
            height: 40px;
            border-top: 1px solid #ddd;
        }
        .form input {
            flex: 1;
            border: none;
            outline: none;
            font-size: 16px;
            color: #111;
            padding: 10px;
        }
        .form button {
            padding: 10px;
            font-size: 20px;
            color: #555;
            background: transparent;
            border: none;
            outline: none;
            cursor: pointer;
        }
        .messages {
            width: 95%;
            height: 389px;
            padding: 10px;
            overflow-y: auto;
        }
        .messages::-webkit-scrollbar {
            width: 4px;
            background: #fff;
        }

        .messages::-webkit-scrollbar-thumb {
            width: 4px;
            background: #888;
        }
        .messages .message {
            margin: 10px 0px;
            display: flex;
        }
        .messages .my-message {
            justify-content: flex-end;
        }
        .messages .message .text {
            padding: 10px;
            background: #f0f4f7;
            color: #555;
            font-size: 13px;
            font-weight: 600;
            max-width: fit-content;
            flex: 1;
            border-radius: 5px;
        }
        .messages .my-message .text {
            background: #555;
            color: #fff;
        }
    </style>
</svelte:head>
