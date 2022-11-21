<script>
    import Switch from "@smui/switch";
    import FormField from "@smui/form-field";
    import Button, { Label } from "@smui/button";
    let bright = true;
    let toggle = true;
</script>

<nav class={toggle == true ? "sidebar" : "close"}>
    <header>
        <i
            class="bx bx-chevron-right toggle"
            on:click={() => (toggle = !toggle)}
        />
    </header>

    <div class="menu-bar">
        <div class="menu">
            <ul class="menu-links">
                <li class="nav-link">
                    <a href="#/chatpage">
                        <i class="bx bx-chat icon" />
                        <span class="text nav-text">Chat</span>
                    </a>
                </li>

                <li class="nav-link">
                    <a href="#/filepage">
                        <i class="bx bxs-cloud-upload icon" />
                        <span class="text nav-text">File</span>
                    </a>
                </li>

                <li class="nav-link">
                    <a href="#/peoplemanagement">
                        <i class="bx bx-user icon" />
                        <span class="text nav-text">Management</span>
                    </a>
                </li>

                <li class="nav-link">
                    <a href="#/">
                        <i class="bx bx-bell icon" />
                        <span class="text nav-text">Notifications</span>
                    </a>
                </li>
            </ul>
        </div>

        <div class="bottom-content">
            <li class="">
                <a href="#/">
                    <i class="bx bx-log-out icon" />
                    <span class="text nav-text">Logout</span>
                </a>
            </li>

            <li class="mode">
                <div class="sun-moon">
                    {#if bright == false}
                        <i class="bx bx-sun icon sun" />
                    {:else}
                        <i class="bx bx-moon icon moon" />
                    {/if}
                </div>

                {#if bright == true}
                    <span class="mode-text text">Dark mode</span>
                {:else}
                    <span class="mode-text text">Light mode</span>
                {/if}

                <div class="toggle-switch">
                    <Switch bind:checked={bright} icons={false} />
                </div>
            </li>
        </div>
    </div>
</nav>

<svelte:head>
    <!-- <link rel="stylesheet" href="src\component\sidebarstyle.css" /> -->
    <link
        href="https://unpkg.com/boxicons@2.1.1/css/boxicons.min.css"
        rel="stylesheet"
    />
    <style>
        body {
            margin: 0px;
            padding: 0px;
            height: 100vh;
        }
    </style>
</svelte:head>

<style lang="less">
    @import "svelte-material-ui/bare.css";
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        font-family: "Poppins", sans-serif;
    }

    :root {
        /* ===== Colors ===== */
        --body-color: #e4e9f7;
        --sidebar-color: #fff;
        --primary-color: #695cfe;
        --primary-color-light: #f6f5ff;
        --toggle-color: #ddd;
        --text-color: #707070;

        /* ====== Transition ====== */
        --tran-03: all 0.2s ease;
        --tran-03: all 0.3s ease;
        --tran-04: all 0.3s ease;
        --tran-05: all 0.3s ease;
    }
    /* ===== Sidebar ===== */

    li {
        border-radius: 6px;
        list-style: none;
        height: 60px;
        margin-top: 10px;
        display: flex;
        align-items: center;
        color: var(--text-color);
        a {
            border-radius: 6px;
            display: flex;
            align-items: center;
            text-decoration: none;
            color: var(--text-color);
            width: 100%;
            height: 100%;
        }
        a:hover {
            background-color: var(--primary-color);
            .icon,
            .text {
                color: var(--sidebar-color);
            }
        }
        .icon {
            min-width: 60px;
            border-radius: 6px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 20px;
            height: 100%;
        }
    }

    .toggle {
        position: absolute;
        top: 25px;
        right: -25px;
        transform: translateY(-50%) rotate(180deg);
        height: 25px;
        width: 25px;
        background-color: var(--primary-color);
        color: var(--sidebar-color);
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 22px;
        cursor: pointer;
        transition: var(--tran-05);
    }

    .sidebar {
        position: fixed;
        padding: 10px 14px;
        height: 100%;
        width: 250px;
        border-right: 1px solid #eee;
        box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.05);
        background-color: var(--sidebar-color);
        transition: var(--tran-05);
        z-index: 100;
        .menu-bar {
            height: calc(100% - 55px);
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            .menu {
                margin-top: 40px;
            }
        }
        .text {
            color: var(--text-color);
            font-size: 17px;
            font-weight: 500;
            white-space: nowrap;
            opacity: 1;
            transition: var(--tran-03);
        }
        header {
            position: relative;
        }
    }

    .close {
        position: fixed;
        padding: 10px 14px;
        height: 100%;
        width: 88px;
        border-right: 1px solid #eee;
        box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.05);
        background-color: var(--sidebar-color);
        transition: var(--tran-05);
        z-index: 100;
        .menu-bar {
            height: calc(100% - 55px);
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            .menu {
                margin-top: 40px;
            }
            .text {
                color: var(--text-color);
                font-size: 17px;
                font-weight: 500;
                white-space: nowrap;
                transition: var(--tran-03);
                opacity: 0;
            }
        }
        .toggle {
            transform: translateY(-50%) rotate(0deg);
        }
        header {
            position: relative;
        }
        .mode {
            position: relative;
            .toggle-switch {
                position: absolute;
                background-color: var(--sidebar-color);
            }
        }
    }

    // .home {
    //     position: absolute;
    //     top: 0;
    //     top: 0;
    //     left: 250px;
    //     height: 100vh;
    //     width: calc(100% - 250px);
    //     background-color: var(--sidebar-color);
    //     transition: var(--tran-05);
    // }

    // .close ~ .home {
    //     position: absolute;
    //     top: 0;
    //     top: 0;
    //     left: 88px;
    //     height: 100vh;
    //     width: calc(100% - 250px);
    //     background-color: var(--sidebar-color);
    //     transition: var(--tran-05);
    // }
</style>
