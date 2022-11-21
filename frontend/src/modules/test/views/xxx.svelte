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

<h1>{"res:" + Response}</h1>
