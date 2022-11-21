<script lang="ts">
    import { onMount } from "svelte";
    import SparkMD5 from "spark-md5";
    import Sidebar from "../component/sidebar.svelte";
    import ImageList, {
        Item,
        ImageAspectContainer,
        Image,
        Supporting,
        Label,
    } from "@smui/image-list";
    import Textfield from "@smui/textfield";
    import Button from "@smui/button/src/Button.svelte";
    let valueTypeFiles: FileList | null = null;
    let fileVar;
    const GetFileMd5 = (file) => {
        return new Promise((resolve) => {
            ``;
            const fileReader = new FileReader();
            const spark = new SparkMD5();
            fileReader.readAsBinaryString(file);
            fileReader.onload = (e) => {
                spark.appendBinary(e.target.result);
                const md5Key = spark.end();
                resolve(md5Key);
            };
        });
    };

    function DownloadFile(filepath, filename) {
        let url = `api/downloadfile?filepath=${filepath}`;
        fetch(url, {
            method: "get",
        })
            .then((v) => {
                console.log("v", v);
                return v.blob();
            })
            .then((Blob) => {
                console.log("Blob", Blob);
                console.log("type", Blob.type);
                var link = document.createElement("a");
                link.href = window.URL.createObjectURL(Blob);
                link.download = filename;
                link.click();
                window.URL.revokeObjectURL(link.href);
            })
            .catch((err) => {
                console.log("err", err);
            });
    }

    async function SubmitFile() {
        const file = new FormData();
        let md5Key = (await GetFileMd5(valueTypeFiles[0])) as string;
        // console.log("md5:", md5Key);
        console.log("filename:", valueTypeFiles[0].name);
        file.append("file", valueTypeFiles[0]);
        file.append("MD5", md5Key);
        file.append("filename", valueTypeFiles[0].name);
        let url = `api/uploadfile`;
        fetch(url, {
            method: "POST",
            mode: "no-cors",
            cache: "no-cache",
            credentials: "same-origin",
            redirect: "follow",
            body: file,
        })
            .then((v) => {
                return v.json();
            })
            .then((res) => {
                console.log(res);
                console.log(res.status);
                if (res.status == 200) {
                    console.log("upload success");
                    data = [
                        ...data,
                        {
                            filepath: res.filepath,
                            filename: valueTypeFiles[0].name,
                        },
                    ];
                }
            })
            .catch((err) => {
                console.log(err);
            });
    }

    let data = [];
    onMount(async () => {
        let url = `api/getFileList`;
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
                data = v.data;
                console.log("res:", v.data);
            })
            .catch((err) => {
                console.log(err);
            });
    });
</script>

<div class="clearfloat">
    <div class="sidebar">
        <Sidebar />
    </div>

    <section class="home">
        <div class="hide-file-ui">
            <Textfield
                bind:files={valueTypeFiles}
                label="选择文件"
                type="file"
            />

            <Button on:click={SubmitFile}>
                <Label>上传文件</Label>
            </Button>
        </div>

        <div class="my-image-list-standard">
            {#each data as d, i}
                <div
                    class="fileitem"
                    on:click={() => {
                        DownloadFile(d.filepath, d.filename);
                    }}
                >
                    <Item>
                        <ImageAspectContainer>
                            <Image src="/file.png" alt="Image {i + 1}" />
                        </ImageAspectContainer>
                        <Supporting>
                            <Label>{d.filename}</Label>
                        </Supporting>
                    </Item>
                </div>
            {/each}
        </div>
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
        float: left;
    }
    .home {
        left: 100px;
        width: auto;
        float: right;
        background-color: #fff;
    }
    .hide-file-ui :global(input[type="file"]::file-selector-button) {
        display: none;
    }

    .hide-file-ui
        :global(:not(.mdc-text-field--label-floating) input[type="file"]) {
        color: transparent;
    }
    .hide-file-ui {
        display: flex;
        justify-content: flex-start;
        align-items: flex-end;
    }
    .my-image-list-standard {
        margin-top: 30px;
        margin-left: 30px;
        width: 100%;
        display: flex;
        flex-wrap: wrap;
        justify-content: flex-start;
        gap: 40px;
    }
    .fileitem {
        flex: 0 0 10%;
        width: 200px;
    }
    :global(.mdc-button:not(:disabled)) {
        font-family: "Roboto", "Helvetica", "Arial", sans-serif;
        font-weight: 500;
        font-size: 0.875rem;
        color: rgb(53, 53, 53);
        background-color: transparent;
        width: 100%;
    }
</style>
