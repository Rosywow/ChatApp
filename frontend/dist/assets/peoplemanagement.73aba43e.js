import{S as X,i as Y,s as Z,a as ee,e as _,c as k,b as w,d as S,f as U,g as p,m as j,t as g,h as te,j as v,k as C,l as y,n as ne,o as se,T as O,p as x,q as A,B as D,r as E,u as le,L as I,v as K}from"./index.5afb83b7.js";function F(s,e,l){const t=s.slice();return t[6]=e[l],t[7]=e,t[8]=l,t}function ae(s){let e;return{c(){e=K("\u4FEE\u6539")},m(l,t){U(l,e,t)},d(l){l&&C(e)}}}function re(s){let e,l;return e=new I({props:{$$slots:{default:[ae]},$$scope:{ctx:s}}}),{c(){k(e.$$.fragment)},m(t,i){j(e,t,i),l=!0},p(t,i){const f={};i&512&&(f.$$scope={dirty:i,ctx:t}),e.$set(f)},i(t){l||(g(e.$$.fragment,t),l=!0)},o(t){v(e.$$.fragment,t),l=!1},d(t){y(e,t)}}}function oe(s){let e;return{c(){e=K("\u5220\u9664")},m(l,t){U(l,e,t)},d(l){l&&C(e)}}}function ie(s){let e,l;return e=new I({props:{$$slots:{default:[oe]},$$scope:{ctx:s}}}),{c(){k(e.$$.fragment)},m(t,i){j(e,t,i),l=!0},p(t,i){const f={};i&512&&(f.$$scope={dirty:i,ctx:t}),e.$set(f)},i(t){l||(g(e.$$.fragment,t),l=!0)},o(t){v(e.$$.fragment,t),l=!1},d(t){y(e,t)}}}function G(s){let e,l,t,i,f,m,d,h,o,n,$,a,u,r,T,B,H,q;function Q(c){s[2](c,s[6])}let L={class:"shaped-outlined",variant:"outlined",type:"text",label:"Username"};s[6].username!==void 0&&(L.value=s[6].username),t=new O({props:L}),x.push(()=>A(t,"value",Q));function R(c){s[3](c,s[6])}let z={class:"shaped-outlined",variant:"outlined",type:"password",label:"Password"};s[6].password!==void 0&&(z.value=s[6].password),d=new O({props:z}),x.push(()=>A(d,"value",R));function V(){return s[4](s[6])}$=new D({props:{size:"large",$$slots:{default:[re]},$$scope:{ctx:s}}}),$.$on("click",V);function W(){return s[5](s[8],s[6])}return r=new D({props:{size:"large",$$slots:{default:[ie]},$$scope:{ctx:s}}}),r.$on("click",W),{c(){e=_("div"),l=_("div"),k(t.$$.fragment),f=w(),m=_("div"),k(d.$$.fragment),o=w(),n=_("div"),k($.$$.fragment),a=w(),u=_("div"),k(r.$$.fragment),T=w(),B=_("br"),H=w(),S(e,"class","item svelte-1gsf2rk")},m(c,b){U(c,e,b),p(e,l),j(t,l,null),p(e,f),p(e,m),j(d,m,null),p(e,o),p(e,n),j($,n,null),p(e,a),p(e,u),j(r,u,null),p(e,T),p(e,B),p(e,H),q=!0},p(c,b){s=c;const P={};!i&&b&1&&(i=!0,P.value=s[6].username,E(()=>i=!1)),t.$set(P);const J={};!h&&b&1&&(h=!0,J.value=s[6].password,E(()=>h=!1)),d.$set(J);const M={};b&512&&(M.$$scope={dirty:b,ctx:s}),$.$set(M);const N={};b&512&&(N.$$scope={dirty:b,ctx:s}),r.$set(N)},i(c){q||(g(t.$$.fragment,c),g(d.$$.fragment,c),g($.$$.fragment,c),g(r.$$.fragment,c),q=!0)},o(c){v(t.$$.fragment,c),v(d.$$.fragment,c),v($.$$.fragment,c),v(r.$$.fragment,c),q=!1},d(c){c&&C(e),y(t),y(d),y($),y(r)}}}function ce(s){let e,l,t,i,f,m,d,h;t=new ee({});let o=s[0],n=[];for(let a=0;a<o.length;a+=1)n[a]=G(F(s,o,a));const $=a=>v(n[a],1,1,()=>{n[a]=null});return{c(){e=_("div"),l=_("div"),k(t.$$.fragment),i=w(),f=_("section");for(let a=0;a<n.length;a+=1)n[a].c();m=w(),d=_("style"),d.textContent=`body {\r
            background-color: #fff;\r
        }`,S(l,"class","sidebar svelte-1gsf2rk"),S(f,"class","home svelte-1gsf2rk"),S(e,"class","clearfloat svelte-1gsf2rk")},m(a,u){U(a,e,u),p(e,l),j(t,l,null),p(e,i),p(e,f);for(let r=0;r<n.length;r+=1)n[r].m(f,null);U(a,m,u),p(document.head,d),h=!0},p(a,[u]){if(u&3){o=a[0];let r;for(r=0;r<o.length;r+=1){const T=F(a,o,r);n[r]?(n[r].p(T,u),g(n[r],1)):(n[r]=G(T),n[r].c(),g(n[r],1),n[r].m(f,null))}for(le(),r=o.length;r<n.length;r+=1)$(r);te()}},i(a){if(!h){g(t.$$.fragment,a);for(let u=0;u<o.length;u+=1)g(n[u]);h=!0}},o(a){v(t.$$.fragment,a),n=n.filter(Boolean);for(let u=0;u<n.length;u+=1)v(n[u]);h=!1},d(a){a&&C(e),y(t),ne(n,a),a&&C(m),C(d)}}}function ue(s){console.log("u",s),fetch("api/changeUser",{method:"put",mode:"cors",cache:"no-cache",credentials:"same-origin",headers:new Headers({"Content-Type":"application/json"}),redirect:"follow",body:JSON.stringify({uid:s.uid,username:s.username,password:s.password})}).then(l=>l.json()).then(l=>{l.status==200&&console.log("success")}).catch(l=>{console.log(l)})}function fe(s,e,l){let t=[];se(()=>{fetch("api/getuser",{method:"get",mode:"no-cors",cache:"no-cache",credentials:"same-origin",headers:new Headers({"Content-Type":"application/json"}),redirect:"follow"}).then(n=>n.json()).then(n=>{n.status==200&&l(0,t=n.data)}).catch(n=>{console.log(n)})});function i(o,n){let $=`api/deleteUser?uid=${n}`;fetch($,{method:"delete",mode:"cors",cache:"no-cache",credentials:"same-origin",headers:new Headers({"Content-Type":"application/json"}),redirect:"follow"}).then(a=>a.json()).then(a=>{a.status==200&&(t.splice(o,1),l(0,t=[...t]))}).catch(a=>{console.log(a)})}function f(o,n){s.$$.not_equal(n.username,o)&&(n.username=o,l(0,t))}function m(o,n){s.$$.not_equal(n.password,o)&&(n.password=o,l(0,t))}return[t,i,f,m,o=>{ue(o)},(o,n)=>{i(o,n.uid)}]}class pe extends X{constructor(e){super(),Y(this,e,fe,ce,Z,{})}}export{pe as default};