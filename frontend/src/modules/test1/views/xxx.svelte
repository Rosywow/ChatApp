<script>
    import { onMount } from "svelte";

    let Response = "";
    onMount(() => {
        fetch(`api/xxx`, {
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
                setTimeout(() => {
                    Response = v.data;
                }, 1000);
            })
            .catch((err) => {
                console.log(err);
            });
    });
</script>

<h1>{"res1:" + Response}</h1>

<i class="bx bx-log-out icon" />
<i class="bx bxs-file-doc" />
<i class="bx bx-bar-chart-alt-2" />
<svelte:head>
    <link
        href="https://unpkg.com/boxicons@2.1.1/css/boxicons.min.css"
        rel="stylesheet"
    />
</svelte:head>
