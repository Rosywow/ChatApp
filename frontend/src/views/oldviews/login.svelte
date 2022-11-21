<script>
    import { push } from "svelte-spa-router";
    import { sender } from "../../stores/index";
    // let count_value;

    // const unsubscribe = sender.subscribe((value) => {
    //     count_value = value;
    // });
    let username = "";
    let password = "";
    function Login() {
        let url = `api/login`;
        fetch(url, {
            method: "POST",
            mode: "cors",
            cache: "no-cache",
            credentials: "same-origin",
            headers: new Headers({
                "Content-Type": "application/json",
            }),
            redirect: "follow",
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        })
            .then((v) => {
                console.log(v);
                return v.json();
            })
            .then((res) => {
                if (res.status === 200) {
                    console.log("登录成功");

                    push("/choose");
                }
            })
            .catch((err) => {
                console.log("err:", err);
            });
    }
</script>

<div class="login-form">
    <div class="form">
        <div class="form-element">
            <h2>Sign In</h2>
        </div>
        <div class="form-element">
            <input type="text" placeholder="username" bind:value={username} />
        </div>
        <div class="form-element">
            <input
                type="password"
                placeholder="password"
                bind:value={password}
            />
        </div>
        <div class="form-element">
            <button
                on:click={() => {
                    Login();
                    localStorage.setItem("username", username);
                    push("/choose");
                }}>Sign In</button
            >
        </div>
    </div>
</div>
<div class="open-login-btn">Sign In</div>

<style>
    * {
        margin: 0px;
        padding: 0px;
        /* box-sizing */
        box-sizing: border-box;
    }
    .login-form {
        /* absolute */
        position: absolute;
        /* top left transform 实现login-form在页面居中 */
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 100%;
        height: 100%;
        /* background and ba color */
        background: #b2b3b4;
        /* flex */
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 2;
    }
    .login-form .form {
        padding: 10px;
        width: 400px;
    }
    .login-form .form .form-element {
        margin: 20px 0px;
    }
    .login-form .form .form-element h2 {
        margin: 20px 0px;
        text-align: center;
        color: #f5f5f5;
        font-size: 25px;
        font-weight: 400;
        text-transform: uppercase;
    }
    .login-form .form .form-element input {
        padding: 10px 0px;
        width: 100%;
        border: none;
        outline: none;
        border-radius: 0px;
        border-bottom: 1px solid #f5f5f5;
        color: #f5f5f5;
        background: transparent;
        font-size: 16px;
    }
    .login-form .form .form-element input::placeholder {
        color: #ddd;
        font-size: 16px;
        text-transform: uppercase;
    }
    .login-form .form .form-element button {
        width: 100%;
        height: 40px;
        margin-top: 10px;
        color: #000000;
        background: #f5f5f5;
        font-size: 18px;
        border: none;
        outline: none;
        text-transform: uppercase;
        transition: 200ms ease-in-out, color 200ms ease-in-out;
    }
    .login-form .form .form-element button:hover {
        color: #f5f5f5;
        background: #000000;
        border: 1px solid #f5f5f5;
    }
</style>
