import{S as w,i as y,s as C,e as h,b as k,A as g,f,y as p,k as u,n as j,o as v,C as H,v as N,g as S,D as T,E as $}from"./index.5afb83b7.js";function _(c,s,t){const n=c.slice();return n[3]=s[t],n[5]=t,n}function m(c){let s,t=c[3].username+"",n,i,l,e,o;function a(){return c[2](c[3])}return{c(){s=h("button"),n=N(t),i=k(),l=h("br")},m(r,d){f(r,s,d),S(s,n),f(r,i,d),f(r,l,d),e||(o=T(s,"click",a),e=!0)},p(r,d){c=r,d&2&&t!==(t=c[3].username+"")&&$(n,t)},d(r){r&&u(s),r&&u(i),r&&u(l),e=!1,o()}}}function b(c){let s,t=c[3].uid!=c[0]&&m(c);return{c(){t&&t.c(),s=g()},m(n,i){t&&t.m(n,i),f(n,s,i)},p(n,i){n[3].uid!=n[0]?t?t.p(n,i):(t=m(n),t.c(),t.m(s.parentNode,s)):t&&(t.d(1),t=null)},d(n){t&&t.d(n),n&&u(s)}}}function q(c){let s,t,n,i=c[1],l=[];for(let e=0;e<i.length;e+=1)l[e]=b(_(c,i,e));return{c(){s=h("br"),t=k();for(let e=0;e<l.length;e+=1)l[e].c();n=g()},m(e,o){f(e,s,o),f(e,t,o);for(let a=0;a<l.length;a+=1)l[a].m(e,o);f(e,n,o)},p(e,[o]){if(o&3){i=e[1];let a;for(a=0;a<i.length;a+=1){const r=_(e,i,a);l[a]?l[a].p(r,o):(l[a]=b(r),l[a].c(),l[a].m(n.parentNode,n))}for(;a<l.length;a+=1)l[a].d(1);l.length=i.length}},i:p,o:p,d(e){e&&u(s),e&&u(t),j(l,e),e&&u(n)}}}function A(c,s,t){let n="",i=[];return v(async()=>{let e="api/getUid";await fetch(e,{method:"get",mode:"no-cors",cache:"no-cache",credentials:"same-origin",headers:new Headers({"Content-Type":"application/json"}),redirect:"follow"}).then(o=>o.json()).then(o=>{t(0,n=o.uid),console.log("sender:",n)}).catch(o=>{console.log(o)}),e="api/getuser",await fetch(e,{method:"get",mode:"no-cors",cache:"no-cache",credentials:"same-origin",headers:new Headers({"Content-Type":"application/json"}),redirect:"follow"}).then(o=>o.json()).then(o=>{o.status==200&&t(1,i=o.data)}).catch(o=>{console.log(o)})}),[n,i,e=>{H("/chat/"+n+"/"+e.uid+"/"+e.username)}]}class E extends w{constructor(s){super(),y(this,s,A,q,C,{})}}export{E as default};