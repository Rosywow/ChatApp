<script>
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

{"用户管理"}
<br />
<br />
<div>
    {#each userList as u, i}
        <!-- <label>uid:</label><input type="text" bind:value={u.uid} /> -->
        {u.uid + ":"}
        <label>username:</label><input type="text" bind:value={u.username} />
        <label>password:</label><input
            type="password"
            bind:value={u.password}
        />
        <button
            on:click={() => {
                changeUser(u);
            }}>修改</button
        >
        <button
            on:click={() => {
                deleteUser(i, u.uid);
            }}>删除</button
        >
        <br />
    {/each}
</div>
