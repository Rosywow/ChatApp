<script>
    import SparkMD5 from "spark-md5";
    import { onMount } from "svelte";
    let fileVar;

    // function SubmitFile() {
    //     // 获取到文件对象
    //     var file = fileVar[0];
    //     let base64;
    //     // 获取FileReader实例
    //     var reader = new FileReader();
    //     // 将文件加载进入
    //     reader.readAsDataURL(file);
    //     reader.onload = function (e) {
    //         // 转换完成输出该文件base64编码
    //         fileList = [
    //             {
    //                 fileName: fileVar[0].name,
    //                 Base64String: this.result,
    //             },
    //             ...fileList,
    //         ];
    //     };
    //     console.log(fileList);
    // }
    // 生成MD5
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
    //提交文件
    async function SubmitFile() {
        const file = new FormData();
        let md5Key = await GetFileMd5(fileVar[0]);
        // console.log("md5:", md5Key);
        console.log("filename:", fileVar[0].name);
        file.append("file", fileVar[0]);
        file.append("MD5", md5Key);
        file.append("filename", fileVar[0].name);
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
                            filename: fileVar[0].name,
                        },
                    ];
                }
            })
            .catch((err) => {
                console.log(err);
            });
    }
    //文件下载
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

<div class="app">
    <div class="screen fs-screen active">
        <div class="file-input">
            <label for="file-input">
                Click here to select files for sharing
            </label>
            <input
                type="file"
                id="file-input"
                bind:files={fileVar}
                on:change={() => console.log(fileVar)}
            />
        </div>
        <input type="submit" on:click={SubmitFile} />

        <div>
            {#each data as d, i}
                <button
                    on:click={() => {
                        DownloadFile(d.filepath, d.filename);
                    }}>{d.filename}</button
                >
                <br />
            {/each}
        </div>
    </div>
</div>

<svelte:head>
    <style>
        * {
            margin: 0px;
            padding: 0px;
            box-sizing: border-box;
        }
        body {
            background-color: #fff;
            font-family: "Roboto", sans-serif;
        }
        .screen {
            display: none;
        }
        .screen.active {
            display: block;
        }
        .app {
            position: fixed;
            top: 0px;
            left: 50%;
            transform: translateX(-50%);
            width: 100%;
            height: 100%;
            max-width: 500px;
            background: #f5f5f5;
            border-right: 1px solid #ddd;
            border-left: 1px solid #ddd;
        }
        .fs-screen {
            padding: 20px;
        }
        .fs-screen .file-input {
            width: 100%;
            border: 2px dashed #555;
        }
        .fs-screen .file-input label {
            display: block;
            width: 100%;
            padding: 40px 50px;
            text-align: center;
            color: #111;
            font-size: 18px;
        }
        .fs-screen .file-input input {
            display: none;
        }
    </style>
</svelte:head>
