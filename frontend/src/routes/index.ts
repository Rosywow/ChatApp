import Login from "../views/oldviews/login.svelte"
import Chat from "../views/oldviews/chat.svelte"
import Filemanage from "../views/oldviews/filemanage.svelte"
import Usermanage from "../views/oldviews/usermanage.svelte"
import Loginpage from "../views/loginpage.svelte"
import Sidebar from "../component/sidebar.svelte"
import Chatpage from "../views/chatpage.svelte"
import Filepage from "../views/filepage.svelte"
import { wrap } from "svelte-spa-router/wrap"


const modules: any = import.meta.glob('@/modules/**/*.route.ts', {
    eager: true,
});


let modulesObject = {};
for (const path in modules) {
    console.log(modules[path].default)
    modulesObject = { ...modulesObject, ...modules[path].default }
}
// console.log("Array:", modulesObject)

const router = {
    "/": Loginpage,
    "/signup": wrap({
        asyncComponent: () => import('../views/signup.svelte')
    }),
    // "/sidebar": Sidebar,
    "/peoplemanagement": wrap({
        asyncComponent: () => import('../views/peoplemanagement.svelte')
    }),
    "/chatpage": Chatpage,
    "/filepage": Filepage,
    "/choose": wrap({
        asyncComponent: () => import('../views/oldviews/choose.svelte')
    }),
    "/chat/:first/:second/:third": Chat,
    "/choosepepole": wrap({
        asyncComponent: () => import('../views/oldviews/choosepeole.svelte')
    }),
    "/filemanage": Filemanage,
    "/usermanage": Usermanage,
    ...modulesObject
}
export default router;