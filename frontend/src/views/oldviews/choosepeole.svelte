<script>
    import { onMount } from "svelte";
    import { push, link } from "svelte-spa-router";
    // import { sender } from "../stores/index";
    import Chat from "./chat.svelte";
    let sender = "";
    let receiver = "";
    let userList = [];
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
                }
            })
            .catch((err) => {
                console.log(err);
            });
    });
</script>

<!-- <br />
<label>Receiver</label>
<input bind:value={receiver} /> -->
<br />
{#each userList as u, i}
    {#if u.uid != sender}
        <button
            on:click={() => {
                push("/chat" + "/" + sender + "/" + u.uid + "/" + u.username);
            }}>{u.username}</button
        >
        <br />
    {/if}
{/each}
<!-- {#if sender != "" && receiver != ""}
    <Chat senderid={sender} receiverid={receiver} />
{/if} -->
