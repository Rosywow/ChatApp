<script lang="ts">
    import Button, { Label, Icon } from "@smui/button";
    import Sidebar from "../component/sidebar.svelte";
    import Textfield from "@smui/textfield";
    import { onMount } from "svelte";
    let userList = [];
    onMount(() => {
        let url = `api/getuser`;
        fetch(url, {
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
    function deleteUser(i, uid) {
        let url = `api/deleteUser?uid=${uid}`;
        fetch(url, {
            method: "delete",
            mode: "cors",
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
                    userList.splice(i, 1);
                    userList = [...userList];
                }
            })
            .catch((err) => {
                console.log(err);
            });
    }
    function changeUser(u) {
        console.log("u", u);
        let url = `api/changeUser`;
        fetch(url, {
            method: "put",
            mode: "cors",
            cache: "no-cache",
            credentials: "same-origin",
            headers: new Headers({
                "Content-Type": "application/json",
            }),
            redirect: "follow",
            body: JSON.stringify({
                uid: u.uid,
                username: u.username,
                password: u.password,
            }),
        })
            .then((v) => {
                return v.json();
            })
            .then((v) => {
                if (v.status == 200) {
                    console.log("success");
                }
            })
            .catch((err) => {
                console.log(err);
            });
    }
</script>

<div class="clearfloat">
    <div class="sidebar">
        <Sidebar />
    </div>
    <section class="home">
        {#each userList as u, i}
            <!-- <label>uid:</label><input type="text" bind:value={u.uid} /> -->
            <div class="item">
                <div>
                    <Textfield
                        class="shaped-outlined"
                        variant="outlined"
                        type="text"
                        bind:value={u.username}
                        label="Username"
                    />
                </div>
                <div>
                    <Textfield
                        class="shaped-outlined"
                        variant="outlined"
                        type="password"
                        bind:value={u.password}
                        label="Password"
                    />
                </div>
                <div>
                    <Button
                        size="large"
                        on:click={() => {
                            changeUser(u);
                        }}
                    >
                        <Label>修改</Label>
                    </Button>
                </div>
                <div>
                    <Button
                        size="large"
                        on:click={() => {
                            deleteUser(i, u.uid);
                        }}
                    >
                        <Label>删除</Label>
                    </Button>
                </div>

                <br />
            </div>
        {/each}
    </section>
</div>

<svelte:head>
    <style>
        body {
            background-color: #fff;
        }
    </style>
</svelte:head>

<style>
    @import "../global.css";
    .clearfloat {
        display: flex;
    }
    .sidebar {
        top: 0;
        left: 0;
        height: 100%;
        width: 88px;
        padding: 10px 14px;
        background: var(--sidebar-color);
        transition: var(--tran-05);
        z-index: 100;
        /* float: left; */
    }
    .home {
        /* top: 60px;
        left: 120px; */
        margin-top: 50px;
        background-color: #fff;
        width: auto;
        /* float: right; */
    }
    .item {
        margin-bottom: 10px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 20px;
    }
</style>
