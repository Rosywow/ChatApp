<script lang="ts">
    import { onMount } from "svelte";
    import Sidebar from "../component/sidebar.svelte";
    import List, {
        Item,
        Graphic,
        Meta,
        Text,
        PrimaryText,
        SecondaryText,
    } from "@smui/list";
    let message = "";
    let messageList = [];
    let chater = "";
    let wb;
    let receiver;
    let selection = "Stephen Hawking";
    let selectionIndex: number | undefined = undefined;
    let userList = [];
    let sender = "";
    onMount(async () => {
        let url = `api/getUid`;
        await fetch(url, {
            method: "get",
            mode: "no-cors",
            cache: "no-cache",
            credentials: "same-origin",
            headers: new Headers({
                "Content-Type": "application/json",
            }),
            redirect: "follow",
        })
            .then((v) => {
                return v.json();
            })
            .then((v) => {
                sender = v.uid;
                console.log("sender:", sender);
            })
            .catch((err) => {
                console.log(err);
            });

        url = `api/getuser`;
        await fetch(url, {
            method: "get",
            mode: "no-cors",
            cache: "no-cache",
            credentials: "same-origin",
            headers: new Headers({
                "Content-Type": "application/json",
            }),
            redirect: "follow",
        })
            .then((v) => {
                return v.json();
            })
            .then((v) => {
                if (v.status == 200) {
                    userList = v.data;
                    console.log("userList", userList);
                }
            })
            .catch((err) => {
                console.log(err);
            });
    });

    function OpenWbSocket() {
        const ws = new WebSocket(
            `ws://localhost:9090/api/chat?sender=${sender}&receiver=${receiver}`
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
        return ws;
    }

    function sendMessage() {
        console.log("发送的消息是：", message);
        console.log("params.first", sender);
        messageList = [
            ...messageList,
            {
                msg: message,
                senderid: sender,
            },
        ];
        let body = JSON.stringify({
            sender: sender,
            receiver: receiver,
            message: message,
        });
        wb.send(body);
        message = "";
        console.log("messageList", messageList);
    }
</script>

<div class="clearfloat">
    <div class="sidebar">
        <Sidebar />
    </div>

    <div class="content">
        <div class="chatlist">
            <List
                style="overflow:auto;"
                class="demo-list"
                twoLine
                avatarList
                singleSelection
                bind:selectedIndex={selectionIndex}
            >
                {#each userList as item}
                    {#if item.uid !== sender}
                        <Item
                            on:SMUI:action={() => (selection = item.username)}
                            selected={selection === item.username}
                            on:click={() => {
                                chater = item.username;
                                receiver = item.uid;
                                wb = OpenWbSocket();
                            }}
                        >
                            <Graphic
                                style="background-image: url(https://place-hold.it/40x40?text=U&fontsize=16);"
                            />
                            <Text>
                                <PrimaryText>{item.username}</PrimaryText>
                            </Text>
                        </Item>
                    {/if}
                {/each}
            </List>
        </div>
    </div>

    <!-- <div class="chatplace">chatlist</div> -->
    <div>
        <main>
            <div>
                <div class="header">
                    <div class="logo">{chater}</div>
                </div>

                <div class="messages">
                    {#each messageList as mes, i}
                        {#if mes.senderid != sender}
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
    </div>
</div>

<svelte:head>
    <link
        rel="stylesheet"
        type="text/css"
        href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"
    />
    <style>
        body {
            background-color: #fff;
        }
    </style>
</svelte:head>

<style>
    .clearfloat {
        display: flex;
        align-items: flex-start;
    }
    .sidebar {
        position: relative;
        top: 0;
        left: 0;
        height: 100%;
        width: 88px;
        background: var(--sidebar-color);
        transition: var(--tran-05);
        z-index: 100;
        /* float: left; */
    }
    .content {
        /* left: 100px;
        right: 0px; */
        height: 100%;
        width: 248px;
        /* float: left; */
        background-color: #fff;
    }
    .chatlist {
        width: 100%;
        overflow: auto;
    }
    /* .chatplace {
        float: left;
    } */

    /* 聊天框 */
    main {
        /* fixed：固定在一个地方，无论浏览器上拉下拉都固定在一个地方 */
        /* position: fixed; */
        /* 以左border为准离外容器的距离，外容器的position 是relative 若是absolute则无效 */
        /* top:50% left 50% 使得左上角在页面中间 */
        /* left: 50%;
        top: 50%; */
        /* transform(-50%,-50%) 向上向左移动自身长宽的50% 使得main在页面居中*/
        /* transform: translate(-50%, -50%); */
        /* width: 100%; */
        /* float: right; */
        width: 880px;
        height: 580px;
        /* max-width: 400px;
        max-height: 500px; */
        background: #fff;
        border: 1px solid #eee;
        box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.05);
        border-top: none;
        border-right: none;
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
    .form {
        position: absolute;
        display: flex;
        bottom: 0px;
        width: 880px;
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
        margin-right: 5px;
        padding: 10px;
        font-size: 20px;
        color: #555;
        background: #fff;
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

    .clearfloat::after {
        clear: both;
        content: "";
        display: block;
    }
    :global(.mdc-deprecated-list-item__text) {
        overflow: unset;
    }
</style>
