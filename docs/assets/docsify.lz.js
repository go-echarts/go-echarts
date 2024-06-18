!function () {
    "use strict";

    function e(e) {
        const n = Object.create(null);
        return function (t) {
            const o = i(t) ? t : JSON.stringify(t);
            return n[o] || (n[o] = e(t))
        }
    }

    const n = e((e => e.replace(/([A-Z])/g, (e => "-" + e.toLowerCase()))));

    function i(e) {
        return "string" == typeof e || "number" == typeof e
    }

    function t() {
    }

    function o(e) {
        return "function" == typeof e
    }

    function a(e) {
        const n = e.match(/^([^:/?#]+:)?(?:\/{2,}([^/?#]*))?([^?#]+)?(\?[^#]*)?(#.*)?/);
        return "string" == typeof n[1] && n[1].length > 0 && n[1].toLowerCase() !== location.protocol || ("string" == typeof n[2] && n[2].length > 0 && n[2].replace(new RegExp(":(" + {
            "http:": 80,
            "https:": 443
        }[location.protocol] + ")?$"), "") !== location.host || !!/^\/\\/.test(e))
    }

    const c = document.body.clientWidth <= 600, r = {};

    function s(e) {
        let n = arguments.length > 1 && void 0 !== arguments[1] && arguments[1];
        if ("string" == typeof e) {
            if (void 0 !== window.Vue) return g(e);
            e = n ? g(e) : r[e] || (r[e] = g(e))
        }
        return e
    }

    const u = document, d = u.body, f = u.head;

    function g(e, n) {
        return n ? e.querySelector(n) : u.querySelector(e)
    }

    function p(e, n) {
        return Array.from(n ? e.querySelectorAll(n) : u.querySelectorAll(e))
    }

    function l(e, n) {
        return e = u.createElement(e), n && (e.innerHTML = n), e
    }

    function v(e, n) {
        return e.appendChild(n)
    }

    function h(e, n) {
        return e.insertBefore(n, e.children[0])
    }

    function _(e, n, i) {
        o(n) ? window.addEventListener(e, n) : e.addEventListener(n, i)
    }

    function m(e, n, i) {
        e && e.classList[i ? n : "toggle"](i || n)
    }

    function b(e) {
        let n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : document;
        const i = n.readyState;
        if ("complete" === i || "interactive" === i) return setTimeout(e, 0);
        n.addEventListener("DOMContentLoaded", e)
    }

    var k = Object.freeze({
        __proto__: null,
        $: u,
        appendTo: v,
        before: h,
        body: d,
        create: l,
        documentReady: b,
        find: g,
        findAll: p,
        getNode: s,
        head: f,
        off: function (e, n, i) {
            o(n) ? window.removeEventListener(e, n) : e.removeEventListener(n, i)
        },
        on: _,
        style: function (e) {
            v(f, l("style", e))
        },
        toggleClass: m
    });
    const w = decodeURIComponent, y = encodeURIComponent;

    function x(e) {
        const n = {};
        return (e = e.trim().replace(/^(\?|#|&)/, "")) ? (e.split("&").forEach((e => {
            const i = e.replace(/\+/g, " ").split("=");
            n[i[0]] = i[1] && w(i[1])
        })), n) : n
    }

    function $(e) {
        let n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : [];
        const i = [];
        for (const t in e) n.indexOf(t) > -1 || i.push(e[t] ? `${y(t)}=${y(e[t])}`.toLowerCase() : y(t));
        return i.length ? `?${i.join("&")}` : ""
    }

    const A = e((e => /(:|(\/{2}))/g.test(e))), S = e((e => e.split(/[?#]/)[0])), E = e((e => {
        if (/\/$/g.test(e)) return e;
        const n = e.match(/(\S*\/)[^/]+$/);
        return n ? n[1] : ""
    })), z = e((e => e.replace(/^\/+/, "/").replace(/([^:])\/{2,}/g, "$1/"))), T = e((e => {
        const n = e.replace(/^\//, "").split("/"), i = [];
        for (const e of n) ".." === e ? i.pop() : "." !== e && i.push(e);
        return "/" + i.join("/")
    }));

    function R(e) {
        return e.split("/").filter((e => -1 === e.indexOf("#"))).join("/")
    }

    function L() {
        for (var e = arguments.length, n = new Array(e), i = 0; i < e; i++) n[i] = arguments[i];
        return z(n.map(R).join("/"))
    }

    const F = e((e => e.replace("#", "?id=")));

    class C {
        #e = {};

        constructor(e) {
            this.config = e
        }

        #n(e, n, i) {
            const t = Object.keys(n).filter((n => (this.#e[n] || (this.#e[n] = new RegExp(`^${n}$`))).test(e) && e !== i))[0];
            return t ? this.#n(e.replace(this.#e[t], n[t]), n, e) : e
        }

        #i(e, n) {
            return new RegExp(`\\.(${n.replace(/^\./,"")}|html)$`, "g").test(e) ? e : /\/$/g.test(e) ? `${e}README${n}` : `${e}${n}`
        }

        getBasePath() {
            return this.config.basePath
        }

        getFile() {
            let e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : this.getCurrentPath(),
                n = arguments.length > 1 ? arguments[1] : void 0;
            const {config: i} = this, t = this.getBasePath(), o = "string" == typeof i.ext ? i.ext : ".md";
            return e = i.alias ? this.#n(e, i.alias) : e, e = this.#i(e, o), e = e === `/README${o}` && i.homepage || e, e = A(e) ? e : L(t, e), n && (e = e.replace(new RegExp(`^${t}`), "")), e
        }

        onchange() {
            (arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : t)()
        }

        getCurrentPath() {
        }

        normalize() {
        }

        parse() {
        }

        toURL(e, n, i) {
            const t = i && "#" === e[0], o = this.parse(F(e));
            if (o.query = {...o.query, ...n}, e = (e = o.path + $(o.query)).replace(/\.md(\?)|\.md$/, "$1"), t) {
                const n = i.indexOf("?");
                e = (n > 0 ? i.substring(0, n) : i) + e
            }
            if (this.config.relativePath && 0 !== e.indexOf("/")) {
                const n = i.substring(0, i.lastIndexOf("/") + 1);
                return z(T(n + e))
            }
            return z("/" + e)
        }
    }

    function j(e) {
        const n = location.href.indexOf("#");
        location.replace(location.href.slice(0, n >= 0 ? n : 0) + "#" + e)
    }

    class O extends C {
        mode = "hash";

        getBasePath() {
            const e = window.location.pathname || "", n = this.config.basePath,
                i = e.endsWith(".html") ? e + "#/" + n : e + "/" + n;
            return /^(\/|https?:)/g.test(n) ? n : z(i)
        }

        getCurrentPath() {
            const e = location.href, n = e.indexOf("#");
            return -1 === n ? "" : e.slice(n + 1)
        }

        onchange() {
            let e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : t, n = !1;
            _("click", (e => {
                const i = "A" === e.target.tagName ? e.target : e.target.parentNode;
                i && "A" === i.tagName && !a(i.href) && (n = !0)
            })), _("hashchange", (i => {
                const t = n ? "navigate" : "history";
                n = !1, e({event: i, source: t})
            }))
        }

        normalize() {
            let e = this.getCurrentPath();
            if (e = F(e), "/" === e.charAt(0)) return j(e);
            j("/" + e)
        }

        parse() {
            let e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : location.href, n = "";
            const i = e.indexOf("#");
            i >= 0 && (e = e.slice(i + 1));
            const t = e.indexOf("?");
            return t >= 0 && (n = e.slice(t + 1), e = e.slice(0, t)), {
                path: e,
                file: this.getFile(e, !0),
                query: x(n),
                response: {}
            }
        }

        toURL(e, n, i) {
            return "#" + super.toURL(e, n, i)
        }
    }

    class q extends C {
        mode = "history";

        getCurrentPath() {
            const e = this.getBasePath();
            let n = window.location.pathname;
            return e && 0 === n.indexOf(e) && (n = n.slice(e.length)), (n || "/") + window.location.search + window.location.hash
        }

        onchange() {
            let e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : t;
            _("click", (n => {
                const i = "A" === n.target.tagName ? n.target : n.target.parentNode;
                if (i && "A" === i.tagName && !a(i.href)) {
                    n.preventDefault();
                    const t = i.href;
                    window.history.pushState({key: t}, "", t), e({event: n, source: "navigate"})
                }
            })), _("popstate", (n => {
                e({event: n, source: "history"})
            }))
        }

        parse() {
            let e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : location.href, n = "";
            const i = e.indexOf("?");
            i >= 0 && (n = e.slice(i + 1), e = e.slice(0, i));
            const t = L(location.origin), o = e.indexOf(t);
            return o > -1 && (e = e.slice(o + t.length)), {path: e, file: this.getFile(e), query: x(n), response: {}}
        }
    }

    let P = {};
    var M = /([^{]*?)\w(?=\})/g, I = {
        YYYY: "getFullYear", YY: "getYear", MM: function (e) {
            return e.getMonth() + 1
        }, DD: "getDate", HH: "getHours", mm: "getMinutes", ss: "getSeconds", fff: "getMilliseconds"
    };

    function N() {
        return {
            async: !1,
            breaks: !1,
            extensions: null,
            gfm: !0,
            hooks: null,
            pedantic: !1,
            renderer: null,
            silent: !1,
            tokenizer: null,
            walkTokens: null
        }
    }

    let H = {
        async: !1,
        breaks: !1,
        extensions: null,
        gfm: !0,
        hooks: null,
        pedantic: !1,
        renderer: null,
        silent: !1,
        tokenizer: null,
        walkTokens: null
    };

    function B(e) {
        H = e
    }

    const D = /[&<>"']/, U = new RegExp(D.source, "g"), Z = /[<>"']|&(?!(#\d{1,7}|#[Xx][a-fA-F0-9]{1,6}|\w+);)/,
        Q = new RegExp(Z.source, "g"), V = {"&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;"},
        W = e => V[e];

    function G(e, n) {
        if (n) {
            if (D.test(e)) return e.replace(U, W)
        } else if (Z.test(e)) return e.replace(Q, W);
        return e
    }

    const Y = /&(#(?:\d+)|(?:#x[0-9A-Fa-f]+)|(?:\w+));?/gi;

    function X(e) {
        return e.replace(Y, ((e, n) => "colon" === (n = n.toLowerCase()) ? ":" : "#" === n.charAt(0) ? "x" === n.charAt(1) ? String.fromCharCode(parseInt(n.substring(2), 16)) : String.fromCharCode(+n.substring(1)) : ""))
    }

    const J = /(^|[^\[])\^/g;

    function K(e, n) {
        let i = "string" == typeof e ? e : e.source;
        n = n || "";
        const t = {
            replace: (e, n) => {
                let o = "string" == typeof n ? n : n.source;
                return o = o.replace(J, "$1"), i = i.replace(e, o), t
            }, getRegex: () => new RegExp(i, n)
        };
        return t
    }

    function ee(e) {
        try {
            e = encodeURI(e).replace(/%25/g, "%")
        } catch (e) {
            return null
        }
        return e
    }

    const ne = {exec: () => null};

    function ie(e, n) {
        const i = e.replace(/\|/g, ((e, n, i) => {
            let t = !1, o = n;
            for (; --o >= 0 && "\\" === i[o];) t = !t;
            return t ? "|" : " |"
        })).split(/ \|/);
        let t = 0;
        if (i[0].trim() || i.shift(), i.length > 0 && !i[i.length - 1].trim() && i.pop(), n) if (i.length > n) i.splice(n); else for (; i.length < n;) i.push("");
        for (; t < i.length; t++) i[t] = i[t].trim().replace(/\\\|/g, "|");
        return i
    }

    function te(e, n, i) {
        const t = e.length;
        if (0 === t) return "";
        let o = 0;
        for (; o < t;) {
            const a = e.charAt(t - o - 1);
            if (a !== n || i) {
                if (a === n || !i) break;
                o++
            } else o++
        }
        return e.slice(0, t - o)
    }

    function oe(e, n, i, t) {
        const o = n.href, a = n.title ? G(n.title) : null, c = e[1].replace(/\\([\[\]])/g, "$1");
        if ("!" !== e[0].charAt(0)) {
            t.state.inLink = !0;
            const e = {type: "link", raw: i, href: o, title: a, text: c, tokens: t.inlineTokens(c)};
            return t.state.inLink = !1, e
        }
        return {type: "image", raw: i, href: o, title: a, text: G(c)}
    }

    class ae {
        options;
        rules;
        lexer;

        constructor(e) {
            this.options = e || H
        }

        space(e) {
            const n = this.rules.block.newline.exec(e);
            if (n && n[0].length > 0) return {type: "space", raw: n[0]}
        }

        code(e) {
            const n = this.rules.block.code.exec(e);
            if (n) {
                const e = n[0].replace(/^ {1,4}/gm, "");
                return {
                    type: "code",
                    raw: n[0],
                    codeBlockStyle: "indented",
                    text: this.options.pedantic ? e : te(e, "\n")
                }
            }
        }

        fences(e) {
            const n = this.rules.block.fences.exec(e);
            if (n) {
                const e = n[0], i = function (e, n) {
                    const i = e.match(/^(\s+)(?:```)/);
                    if (null === i) return n;
                    const t = i[1];
                    return n.split("\n").map((e => {
                        const n = e.match(/^\s+/);
                        if (null === n) return e;
                        const [i] = n;
                        return i.length >= t.length ? e.slice(t.length) : e
                    })).join("\n")
                }(e, n[3] || "");
                return {
                    type: "code",
                    raw: e,
                    lang: n[2] ? n[2].trim().replace(this.rules.inline.anyPunctuation, "$1") : n[2],
                    text: i
                }
            }
        }

        heading(e) {
            const n = this.rules.block.heading.exec(e);
            if (n) {
                let e = n[2].trim();
                if (/#$/.test(e)) {
                    const n = te(e, "#");
                    this.options.pedantic ? e = n.trim() : n && !/ $/.test(n) || (e = n.trim())
                }
                return {type: "heading", raw: n[0], depth: n[1].length, text: e, tokens: this.lexer.inline(e)}
            }
        }

        hr(e) {
            const n = this.rules.block.hr.exec(e);
            if (n) return {type: "hr", raw: n[0]}
        }

        blockquote(e) {
            const n = this.rules.block.blockquote.exec(e);
            if (n) {
                let e = n[0].replace(/\n {0,3}((?:=+|-+) *)(?=\n|$)/g, "\n    $1");
                e = te(e.replace(/^ *>[ \t]?/gm, ""), "\n");
                const i = this.lexer.state.top;
                this.lexer.state.top = !0;
                const t = this.lexer.blockTokens(e);
                return this.lexer.state.top = i, {type: "blockquote", raw: n[0], tokens: t, text: e}
            }
        }

        list(e) {
            let n = this.rules.block.list.exec(e);
            if (n) {
                let i = n[1].trim();
                const t = i.length > 1,
                    o = {type: "list", raw: "", ordered: t, start: t ? +i.slice(0, -1) : "", loose: !1, items: []};
                i = t ? `\\d{1,9}\\${i.slice(-1)}` : `\\${i}`, this.options.pedantic && (i = t ? i : "[*+-]");
                const a = new RegExp(`^( {0,3}${i})((?:[\t ][^\\n]*)?(?:\\n|$))`);
                let c = "", r = "", s = !1;
                for (; e;) {
                    let i = !1;
                    if (!(n = a.exec(e))) break;
                    if (this.rules.block.hr.test(e)) break;
                    c = n[0], e = e.substring(c.length);
                    let t = n[2].split("\n", 1)[0].replace(/^\t+/, (e => " ".repeat(3 * e.length))),
                        u = e.split("\n", 1)[0], d = 0;
                    this.options.pedantic ? (d = 2, r = t.trimStart()) : (d = n[2].search(/[^ ]/), d = d > 4 ? 1 : d, r = t.slice(d), d += n[1].length);
                    let f = !1;
                    if (!t && /^ *$/.test(u) && (c += u + "\n", e = e.substring(u.length + 1), i = !0), !i) {
                        const n = new RegExp(`^ {0,${Math.min(3,d-1)}}(?:[*+-]|\\d{1,9}[.)])((?:[ \t][^\\n]*)?(?:\\n|$))`),
                            i = new RegExp(`^ {0,${Math.min(3,d-1)}}((?:- *){3,}|(?:_ *){3,}|(?:\\* *){3,})(?:\\n+|$)`),
                            o = new RegExp(`^ {0,${Math.min(3,d-1)}}(?:\`\`\`|~~~)`),
                            a = new RegExp(`^ {0,${Math.min(3,d-1)}}#`);
                        for (; e;) {
                            const s = e.split("\n", 1)[0];
                            if (u = s, this.options.pedantic && (u = u.replace(/^ {1,4}(?=( {4})*[^ ])/g, "  ")), o.test(u)) break;
                            if (a.test(u)) break;
                            if (n.test(u)) break;
                            if (i.test(e)) break;
                            if (u.search(/[^ ]/) >= d || !u.trim()) r += "\n" + u.slice(d); else {
                                if (f) break;
                                if (t.search(/[^ ]/) >= 4) break;
                                if (o.test(t)) break;
                                if (a.test(t)) break;
                                if (i.test(t)) break;
                                r += "\n" + u
                            }
                            f || u.trim() || (f = !0), c += s + "\n", e = e.substring(s.length + 1), t = u.slice(d)
                        }
                    }
                    o.loose || (s ? o.loose = !0 : /\n *\n *$/.test(c) && (s = !0));
                    let g, p = null;
                    this.options.gfm && (p = /^\[[ xX]\] /.exec(r), p && (g = "[ ] " !== p[0], r = r.replace(/^\[[ xX]\] +/, ""))), o.items.push({
                        type: "list_item",
                        raw: c,
                        task: !!p,
                        checked: g,
                        loose: !1,
                        text: r,
                        tokens: []
                    }), o.raw += c
                }
                o.items[o.items.length - 1].raw = c.trimEnd(), o.items[o.items.length - 1].text = r.trimEnd(), o.raw = o.raw.trimEnd();
                for (let e = 0; e < o.items.length; e++) if (this.lexer.state.top = !1, o.items[e].tokens = this.lexer.blockTokens(o.items[e].text, []), !o.loose) {
                    const n = o.items[e].tokens.filter((e => "space" === e.type)),
                        i = n.length > 0 && n.some((e => /\n.*\n/.test(e.raw)));
                    o.loose = i
                }
                if (o.loose) for (let e = 0; e < o.items.length; e++) o.items[e].loose = !0;
                return o
            }
        }

        html(e) {
            const n = this.rules.block.html.exec(e);
            if (n) {
                return {
                    type: "html",
                    block: !0,
                    raw: n[0],
                    pre: "pre" === n[1] || "script" === n[1] || "style" === n[1],
                    text: n[0]
                }
            }
        }

        def(e) {
            const n = this.rules.block.def.exec(e);
            if (n) {
                const e = n[1].toLowerCase().replace(/\s+/g, " "),
                    i = n[2] ? n[2].replace(/^<(.*)>$/, "$1").replace(this.rules.inline.anyPunctuation, "$1") : "",
                    t = n[3] ? n[3].substring(1, n[3].length - 1).replace(this.rules.inline.anyPunctuation, "$1") : n[3];
                return {type: "def", tag: e, raw: n[0], href: i, title: t}
            }
        }

        table(e) {
            const n = this.rules.block.table.exec(e);
            if (!n) return;
            if (!/[:|]/.test(n[2])) return;
            const i = ie(n[1]), t = n[2].replace(/^\||\| *$/g, "").split("|"),
                o = n[3] && n[3].trim() ? n[3].replace(/\n[ \t]*$/, "").split("\n") : [],
                a = {type: "table", raw: n[0], header: [], align: [], rows: []};
            if (i.length === t.length) {
                for (const e of t) /^ *-+: *$/.test(e) ? a.align.push("right") : /^ *:-+: *$/.test(e) ? a.align.push("center") : /^ *:-+ *$/.test(e) ? a.align.push("left") : a.align.push(null);
                for (const e of i) a.header.push({text: e, tokens: this.lexer.inline(e)});
                for (const e of o) a.rows.push(ie(e, a.header.length).map((e => ({
                    text: e,
                    tokens: this.lexer.inline(e)
                }))));
                return a
            }
        }

        lheading(e) {
            const n = this.rules.block.lheading.exec(e);
            if (n) return {
                type: "heading",
                raw: n[0],
                depth: "=" === n[2].charAt(0) ? 1 : 2,
                text: n[1],
                tokens: this.lexer.inline(n[1])
            }
        }

        paragraph(e) {
            const n = this.rules.block.paragraph.exec(e);
            if (n) {
                const e = "\n" === n[1].charAt(n[1].length - 1) ? n[1].slice(0, -1) : n[1];
                return {type: "paragraph", raw: n[0], text: e, tokens: this.lexer.inline(e)}
            }
        }

        text(e) {
            const n = this.rules.block.text.exec(e);
            if (n) return {type: "text", raw: n[0], text: n[0], tokens: this.lexer.inline(n[0])}
        }

        escape(e) {
            const n = this.rules.inline.escape.exec(e);
            if (n) return {type: "escape", raw: n[0], text: G(n[1])}
        }

        tag(e) {
            const n = this.rules.inline.tag.exec(e);
            if (n) return !this.lexer.state.inLink && /^<a /i.test(n[0]) ? this.lexer.state.inLink = !0 : this.lexer.state.inLink && /^<\/a>/i.test(n[0]) && (this.lexer.state.inLink = !1), !this.lexer.state.inRawBlock && /^<(pre|code|kbd|script)(\s|>)/i.test(n[0]) ? this.lexer.state.inRawBlock = !0 : this.lexer.state.inRawBlock && /^<\/(pre|code|kbd|script)(\s|>)/i.test(n[0]) && (this.lexer.state.inRawBlock = !1), {
                type: "html",
                raw: n[0],
                inLink: this.lexer.state.inLink,
                inRawBlock: this.lexer.state.inRawBlock,
                block: !1,
                text: n[0]
            }
        }

        link(e) {
            const n = this.rules.inline.link.exec(e);
            if (n) {
                const e = n[2].trim();
                if (!this.options.pedantic && /^</.test(e)) {
                    if (!/>$/.test(e)) return;
                    const n = te(e.slice(0, -1), "\\");
                    if ((e.length - n.length) % 2 == 0) return
                } else {
                    const e = function (e, n) {
                        if (-1 === e.indexOf(n[1])) return -1;
                        let i = 0;
                        for (let t = 0; t < e.length; t++) if ("\\" === e[t]) t++; else if (e[t] === n[0]) i++; else if (e[t] === n[1] && (i--, i < 0)) return t;
                        return -1
                    }(n[2], "()");
                    if (e > -1) {
                        const i = (0 === n[0].indexOf("!") ? 5 : 4) + n[1].length + e;
                        n[2] = n[2].substring(0, e), n[0] = n[0].substring(0, i).trim(), n[3] = ""
                    }
                }
                let i = n[2], t = "";
                if (this.options.pedantic) {
                    const e = /^([^'"]*[^\s])\s+(['"])(.*)\2/.exec(i);
                    e && (i = e[1], t = e[3])
                } else t = n[3] ? n[3].slice(1, -1) : "";
                return i = i.trim(), /^</.test(i) && (i = this.options.pedantic && !/>$/.test(e) ? i.slice(1) : i.slice(1, -1)), oe(n, {
                    href: i ? i.replace(this.rules.inline.anyPunctuation, "$1") : i,
                    title: t ? t.replace(this.rules.inline.anyPunctuation, "$1") : t
                }, n[0], this.lexer)
            }
        }

        reflink(e, n) {
            let i;
            if ((i = this.rules.inline.reflink.exec(e)) || (i = this.rules.inline.nolink.exec(e))) {
                const e = n[(i[2] || i[1]).replace(/\s+/g, " ").toLowerCase()];
                if (!e) {
                    const e = i[0].charAt(0);
                    return {type: "text", raw: e, text: e}
                }
                return oe(i, e, i[0], this.lexer)
            }
        }

        emStrong(e, n) {
            let i = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : "",
                t = this.rules.inline.emStrongLDelim.exec(e);
            if (!t) return;
            if (t[3] && i.match(/[\p{L}\p{N}]/u)) return;
            if (!(t[1] || t[2] || "") || !i || this.rules.inline.punctuation.exec(i)) {
                const i = [...t[0]].length - 1;
                let o, a, c = i, r = 0;
                const s = "*" === t[0][0] ? this.rules.inline.emStrongRDelimAst : this.rules.inline.emStrongRDelimUnd;
                for (s.lastIndex = 0, n = n.slice(-1 * e.length + i); null != (t = s.exec(n));) {
                    if (o = t[1] || t[2] || t[3] || t[4] || t[5] || t[6], !o) continue;
                    if (a = [...o].length, t[3] || t[4]) {
                        c += a;
                        continue
                    }
                    if ((t[5] || t[6]) && i % 3 && !((i + a) % 3)) {
                        r += a;
                        continue
                    }
                    if (c -= a, c > 0) continue;
                    a = Math.min(a, a + c + r);
                    const n = [...t[0]][0].length, s = e.slice(0, i + t.index + n + a);
                    if (Math.min(i, a) % 2) {
                        const e = s.slice(1, -1);
                        return {type: "em", raw: s, text: e, tokens: this.lexer.inlineTokens(e)}
                    }
                    const u = s.slice(2, -2);
                    return {type: "strong", raw: s, text: u, tokens: this.lexer.inlineTokens(u)}
                }
            }
        }

        codespan(e) {
            const n = this.rules.inline.code.exec(e);
            if (n) {
                let e = n[2].replace(/\n/g, " ");
                const i = /[^ ]/.test(e), t = /^ /.test(e) && / $/.test(e);
                return i && t && (e = e.substring(1, e.length - 1)), e = G(e, !0), {
                    type: "codespan",
                    raw: n[0],
                    text: e
                }
            }
        }

        br(e) {
            const n = this.rules.inline.br.exec(e);
            if (n) return {type: "br", raw: n[0]}
        }

        del(e) {
            const n = this.rules.inline.del.exec(e);
            if (n) return {type: "del", raw: n[0], text: n[2], tokens: this.lexer.inlineTokens(n[2])}
        }

        autolink(e) {
            const n = this.rules.inline.autolink.exec(e);
            if (n) {
                let e, i;
                return "@" === n[2] ? (e = G(n[1]), i = "mailto:" + e) : (e = G(n[1]), i = e), {
                    type: "link",
                    raw: n[0],
                    text: e,
                    href: i,
                    tokens: [{type: "text", raw: e, text: e}]
                }
            }
        }

        url(e) {
            let n;
            if (n = this.rules.inline.url.exec(e)) {
                let e, i;
                if ("@" === n[2]) e = G(n[0]), i = "mailto:" + e; else {
                    let t;
                    do {
                        t = n[0], n[0] = this.rules.inline._backpedal.exec(n[0])?.[0] ?? ""
                    } while (t !== n[0]);
                    e = G(n[0]), i = "www." === n[1] ? "http://" + n[0] : n[0]
                }
                return {type: "link", raw: n[0], text: e, href: i, tokens: [{type: "text", raw: e, text: e}]}
            }
        }

        inlineText(e) {
            const n = this.rules.inline.text.exec(e);
            if (n) {
                let e;
                return e = this.lexer.state.inRawBlock ? n[0] : G(n[0]), {type: "text", raw: n[0], text: e}
            }
        }
    }

    const ce = /^ {0,3}((?:-[\t ]*){3,}|(?:_[ \t]*){3,}|(?:\*[ \t]*){3,})(?:\n+|$)/, re = /(?:[*+-]|\d{1,9}[.)])/,
        se = K(/^(?!bull |blockCode|fences|blockquote|heading|html)((?:.|\n(?!\s*?\n|bull |blockCode|fences|blockquote|heading|html))+?)\n {0,3}(=+|-+) *(?:\n+|$)/).replace(/bull/g, re).replace(/blockCode/g, / {4}/).replace(/fences/g, / {0,3}(?:`{3,}|~{3,})/).replace(/blockquote/g, / {0,3}>/).replace(/heading/g, / {0,3}#{1,6}/).replace(/html/g, / {0,3}<[^\n>]+>\n/).getRegex(),
        ue = /^([^\n]+(?:\n(?!hr|heading|lheading|blockquote|fences|list|html|table| +\n)[^\n]+)*)/,
        de = /(?!\s*\])(?:\\.|[^\[\]\\])+/,
        fe = K(/^ {0,3}\[(label)\]: *(?:\n *)?([^<\s][^\s]*|<.*?>)(?:(?: +(?:\n *)?| *\n *)(title))? *(?:\n+|$)/).replace("label", de).replace("title", /(?:"(?:\\"?|[^"\\])*"|'[^'\n]*(?:\n[^'\n]+)*\n?'|\([^()]*\))/).getRegex(),
        ge = K(/^( {0,3}bull)([ \t][^\n]+?)?(?:\n|$)/).replace(/bull/g, re).getRegex(),
        pe = "address|article|aside|base|basefont|blockquote|body|caption|center|col|colgroup|dd|details|dialog|dir|div|dl|dt|fieldset|figcaption|figure|footer|form|frame|frameset|h[1-6]|head|header|hr|html|iframe|legend|li|link|main|menu|menuitem|meta|nav|noframes|ol|optgroup|option|p|param|search|section|summary|table|tbody|td|tfoot|th|thead|title|tr|track|ul",
        le = /<!--(?:-?>|[\s\S]*?(?:-->|$))/,
        ve = K("^ {0,3}(?:<(script|pre|style|textarea)[\\s>][\\s\\S]*?(?:</\\1>[^\\n]*\\n+|$)|comment[^\\n]*(\\n+|$)|<\\?[\\s\\S]*?(?:\\?>\\n*|$)|<![A-Z][\\s\\S]*?(?:>\\n*|$)|<!\\[CDATA\\[[\\s\\S]*?(?:\\]\\]>\\n*|$)|</?(tag)(?: +|\\n|/?>)[\\s\\S]*?(?:(?:\\n *)+\\n|$)|<(?!script|pre|style|textarea)([a-z][\\w-]*)(?:attribute)*? */?>(?=[ \\t]*(?:\\n|$))[\\s\\S]*?(?:(?:\\n *)+\\n|$)|</(?!script|pre|style|textarea)[a-z][\\w-]*\\s*>(?=[ \\t]*(?:\\n|$))[\\s\\S]*?(?:(?:\\n *)+\\n|$))", "i").replace("comment", le).replace("tag", pe).replace("attribute", / +[a-zA-Z:_][\w.:-]*(?: *= *"[^"\n]*"| *= *'[^'\n]*'| *= *[^\s"'=<>`]+)?/).getRegex(),
        he = K(ue).replace("hr", ce).replace("heading", " {0,3}#{1,6}(?:\\s|$)").replace("|lheading", "").replace("|table", "").replace("blockquote", " {0,3}>").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", pe).getRegex(),
        _e = {
            blockquote: K(/^( {0,3}> ?(paragraph|[^\n]*)(?:\n|$))+/).replace("paragraph", he).getRegex(),
            code: /^( {4}[^\n]+(?:\n(?: *(?:\n|$))*)?)+/,
            def: fe,
            fences: /^ {0,3}(`{3,}(?=[^`\n]*(?:\n|$))|~{3,})([^\n]*)(?:\n|$)(?:|([\s\S]*?)(?:\n|$))(?: {0,3}\1[~`]* *(?=\n|$)|$)/,
            heading: /^ {0,3}(#{1,6})(?=\s|$)(.*)(?:\n+|$)/,
            hr: ce,
            html: ve,
            lheading: se,
            list: ge,
            newline: /^(?: *(?:\n|$))+/,
            paragraph: he,
            table: ne,
            text: /^[^\n]+/
        },
        me = K("^ *([^\\n ].*)\\n {0,3}((?:\\| *)?:?-+:? *(?:\\| *:?-+:? *)*(?:\\| *)?)(?:\\n((?:(?! *\\n|hr|heading|blockquote|code|fences|list|html).*(?:\\n|$))*)\\n*|$)").replace("hr", ce).replace("heading", " {0,3}#{1,6}(?:\\s|$)").replace("blockquote", " {0,3}>").replace("code", " {4}[^\\n]").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", pe).getRegex(),
        be = {
            ..._e,
            table: me,
            paragraph: K(ue).replace("hr", ce).replace("heading", " {0,3}#{1,6}(?:\\s|$)").replace("|lheading", "").replace("table", me).replace("blockquote", " {0,3}>").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", pe).getRegex()
        }, ke = {
            ..._e,
            html: K("^ *(?:comment *(?:\\n|\\s*$)|<(tag)[\\s\\S]+?</\\1> *(?:\\n{2,}|\\s*$)|<tag(?:\"[^\"]*\"|'[^']*'|\\s[^'\"/>\\s]*)*?/?> *(?:\\n{2,}|\\s*$))").replace("comment", le).replace(/tag/g, "(?!(?:a|em|strong|small|s|cite|q|dfn|abbr|data|time|code|var|samp|kbd|sub|sup|i|b|u|mark|ruby|rt|rp|bdi|bdo|span|br|wbr|ins|del|img)\\b)\\w+(?!:|[^\\w\\s@]*@)\\b").getRegex(),
            def: /^ *\[([^\]]+)\]: *<?([^\s>]+)>?(?: +(["(][^\n]+[")]))? *(?:\n+|$)/,
            heading: /^(#{1,6})(.*)(?:\n+|$)/,
            fences: ne,
            lheading: /^(.+?)\n {0,3}(=+|-+) *(?:\n+|$)/,
            paragraph: K(ue).replace("hr", ce).replace("heading", " *#{1,6} *[^\n]").replace("lheading", se).replace("|table", "").replace("blockquote", " {0,3}>").replace("|fences", "").replace("|list", "").replace("|html", "").replace("|tag", "").getRegex()
        }, we = /^\\([!"#$%&'()*+,\-./:;<=>?@\[\]\\^_`{|}~])/, ye = /^( {2,}|\\)\n(?!\s*$)/, xe = "\\p{P}\\p{S}",
        $e = K(/^((?![*_])[\spunctuation])/, "u").replace(/punctuation/g, xe).getRegex(),
        Ae = K(/^(?:\*+(?:((?!\*)[punct])|[^\s*]))|^_+(?:((?!_)[punct])|([^\s_]))/, "u").replace(/punct/g, xe).getRegex(),
        Se = K("^[^_*]*?__[^_*]*?\\*[^_*]*?(?=__)|[^*]+(?=[^*])|(?!\\*)[punct](\\*+)(?=[\\s]|$)|[^punct\\s](\\*+)(?!\\*)(?=[punct\\s]|$)|(?!\\*)[punct\\s](\\*+)(?=[^punct\\s])|[\\s](\\*+)(?!\\*)(?=[punct])|(?!\\*)[punct](\\*+)(?!\\*)(?=[punct])|[^punct\\s](\\*+)(?=[^punct\\s])", "gu").replace(/punct/g, xe).getRegex(),
        Ee = K("^[^_*]*?\\*\\*[^_*]*?_[^_*]*?(?=\\*\\*)|[^_]+(?=[^_])|(?!_)[punct](_+)(?=[\\s]|$)|[^punct\\s](_+)(?!_)(?=[punct\\s]|$)|(?!_)[punct\\s](_+)(?=[^punct\\s])|[\\s](_+)(?!_)(?=[punct])|(?!_)[punct](_+)(?!_)(?=[punct])", "gu").replace(/punct/g, xe).getRegex(),
        ze = K(/\\([punct])/, "gu").replace(/punct/g, xe).getRegex(),
        Te = K(/^<(scheme:[^\s\x00-\x1f<>]*|email)>/).replace("scheme", /[a-zA-Z][a-zA-Z0-9+.-]{1,31}/).replace("email", /[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+(@)[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+(?![-_])/).getRegex(),
        Re = K(le).replace("(?:--\x3e|$)", "--\x3e").getRegex(),
        Le = K("^comment|^</[a-zA-Z][\\w:-]*\\s*>|^<[a-zA-Z][\\w-]*(?:attribute)*?\\s*/?>|^<\\?[\\s\\S]*?\\?>|^<![a-zA-Z]+\\s[\\s\\S]*?>|^<!\\[CDATA\\[[\\s\\S]*?\\]\\]>").replace("comment", Re).replace("attribute", /\s+[a-zA-Z:_][\w.:-]*(?:\s*=\s*"[^"]*"|\s*=\s*'[^']*'|\s*=\s*[^\s"'=<>`]+)?/).getRegex(),
        Fe = /(?:\[(?:\\.|[^\[\]\\])*\]|\\.|`[^`]*`|[^\[\]\\`])*?/,
        Ce = K(/^!?\[(label)\]\(\s*(href)(?:\s+(title))?\s*\)/).replace("label", Fe).replace("href", /<(?:\\.|[^\n<>\\])+>|[^\s\x00-\x1f]*/).replace("title", /"(?:\\"?|[^"\\])*"|'(?:\\'?|[^'\\])*'|\((?:\\\)?|[^)\\])*\)/).getRegex(),
        je = K(/^!?\[(label)\]\[(ref)\]/).replace("label", Fe).replace("ref", de).getRegex(),
        Oe = K(/^!?\[(ref)\](?:\[\])?/).replace("ref", de).getRegex(), qe = {
            _backpedal: ne,
            anyPunctuation: ze,
            autolink: Te,
            blockSkip: /\[[^[\]]*?\]\([^\(\)]*?\)|`[^`]*?`|<[^<>]*?>/g,
            br: ye,
            code: /^(`+)([^`]|[^`][\s\S]*?[^`])\1(?!`)/,
            del: ne,
            emStrongLDelim: Ae,
            emStrongRDelimAst: Se,
            emStrongRDelimUnd: Ee,
            escape: we,
            link: Ce,
            nolink: Oe,
            punctuation: $e,
            reflink: je,
            reflinkSearch: K("reflink|nolink(?!\\()", "g").replace("reflink", je).replace("nolink", Oe).getRegex(),
            tag: Le,
            text: /^(`+|[^`])(?:(?= {2,}\n)|[\s\S]*?(?:(?=[\\<!\[`*_]|\b_|$)|[^ ](?= {2,}\n)))/,
            url: ne
        }, Pe = {
            ...qe,
            link: K(/^!?\[(label)\]\((.*?)\)/).replace("label", Fe).getRegex(),
            reflink: K(/^!?\[(label)\]\s*\[([^\]]*)\]/).replace("label", Fe).getRegex()
        }, Me = {
            ...qe,
            escape: K(we).replace("])", "~|])").getRegex(),
            url: K(/^((?:ftp|https?):\/\/|www\.)(?:[a-zA-Z0-9\-]+\.?)+[^\s<]*|^email/, "i").replace("email", /[A-Za-z0-9._+-]+(@)[a-zA-Z0-9-_]+(?:\.[a-zA-Z0-9-_]*[a-zA-Z0-9])+(?![-_])/).getRegex(),
            _backpedal: /(?:[^?!.,:;*_'"~()&]+|\([^)]*\)|&(?![a-zA-Z0-9]+;$)|[?!.,:;*_'"~)]+(?!$))+/,
            del: /^(~~?)(?=[^\s~])([\s\S]*?[^\s~])\1(?=[^~]|$)/,
            text: /^([`~]+|[^`~])(?:(?= {2,}\n)|(?=[a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-]+@)|[\s\S]*?(?:(?=[\\<!\[`*~_]|\b_|https?:\/\/|ftp:\/\/|www\.|$)|[^ ](?= {2,}\n)|[^a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-](?=[a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-]+@)))/
        }, Ie = {
            ...Me,
            br: K(ye).replace("{2,}", "*").getRegex(),
            text: K(Me.text).replace("\\b_", "\\b_| {2,}\\n").replace(/\{2,\}/g, "*").getRegex()
        }, Ne = {normal: _e, gfm: be, pedantic: ke}, He = {normal: qe, gfm: Me, breaks: Ie, pedantic: Pe};

    class Be {
        tokens;
        options;
        state;
        tokenizer;
        inlineQueue;

        constructor(e) {
            this.tokens = [], this.tokens.links = Object.create(null), this.options = e || H, this.options.tokenizer = this.options.tokenizer || new ae, this.tokenizer = this.options.tokenizer, this.tokenizer.options = this.options, this.tokenizer.lexer = this, this.inlineQueue = [], this.state = {
                inLink: !1,
                inRawBlock: !1,
                top: !0
            };
            const n = {block: Ne.normal, inline: He.normal};
            this.options.pedantic ? (n.block = Ne.pedantic, n.inline = He.pedantic) : this.options.gfm && (n.block = Ne.gfm, this.options.breaks ? n.inline = He.breaks : n.inline = He.gfm), this.tokenizer.rules = n
        }

        static get rules() {
            return {block: Ne, inline: He}
        }

        static lex(e, n) {
            return new Be(n).lex(e)
        }

        static lexInline(e, n) {
            return new Be(n).inlineTokens(e)
        }

        lex(e) {
            e = e.replace(/\r\n|\r/g, "\n"), this.blockTokens(e, this.tokens);
            for (let e = 0; e < this.inlineQueue.length; e++) {
                const n = this.inlineQueue[e];
                this.inlineTokens(n.src, n.tokens)
            }
            return this.inlineQueue = [], this.tokens
        }

        blockTokens(e) {
            let n, i, t, o, a = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : [];
            for (e = this.options.pedantic ? e.replace(/\t/g, "    ").replace(/^ +$/gm, "") : e.replace(/^( *)(\t+)/gm, ((e, n, i) => n + "    ".repeat(i.length))); e;) if (!(this.options.extensions && this.options.extensions.block && this.options.extensions.block.some((i => !!(n = i.call({lexer: this}, e, a)) && (e = e.substring(n.raw.length), a.push(n), !0))))) if (n = this.tokenizer.space(e)) e = e.substring(n.raw.length), 1 === n.raw.length && a.length > 0 ? a[a.length - 1].raw += "\n" : a.push(n); else if (n = this.tokenizer.code(e)) e = e.substring(n.raw.length), i = a[a.length - 1], !i || "paragraph" !== i.type && "text" !== i.type ? a.push(n) : (i.raw += "\n" + n.raw, i.text += "\n" + n.text, this.inlineQueue[this.inlineQueue.length - 1].src = i.text); else if (n = this.tokenizer.fences(e)) e = e.substring(n.raw.length), a.push(n); else if (n = this.tokenizer.heading(e)) e = e.substring(n.raw.length), a.push(n); else if (n = this.tokenizer.hr(e)) e = e.substring(n.raw.length), a.push(n); else if (n = this.tokenizer.blockquote(e)) e = e.substring(n.raw.length), a.push(n); else if (n = this.tokenizer.list(e)) e = e.substring(n.raw.length), a.push(n); else if (n = this.tokenizer.html(e)) e = e.substring(n.raw.length), a.push(n); else if (n = this.tokenizer.def(e)) e = e.substring(n.raw.length), i = a[a.length - 1], !i || "paragraph" !== i.type && "text" !== i.type ? this.tokens.links[n.tag] || (this.tokens.links[n.tag] = {
                href: n.href,
                title: n.title
            }) : (i.raw += "\n" + n.raw, i.text += "\n" + n.raw, this.inlineQueue[this.inlineQueue.length - 1].src = i.text); else if (n = this.tokenizer.table(e)) e = e.substring(n.raw.length), a.push(n); else if (n = this.tokenizer.lheading(e)) e = e.substring(n.raw.length), a.push(n); else {
                if (t = e, this.options.extensions && this.options.extensions.startBlock) {
                    let n = 1 / 0;
                    const i = e.slice(1);
                    let o;
                    this.options.extensions.startBlock.forEach((e => {
                        o = e.call({lexer: this}, i), "number" == typeof o && o >= 0 && (n = Math.min(n, o))
                    })), n < 1 / 0 && n >= 0 && (t = e.substring(0, n + 1))
                }
                if (this.state.top && (n = this.tokenizer.paragraph(t))) i = a[a.length - 1], o && "paragraph" === i.type ? (i.raw += "\n" + n.raw, i.text += "\n" + n.text, this.inlineQueue.pop(), this.inlineQueue[this.inlineQueue.length - 1].src = i.text) : a.push(n), o = t.length !== e.length, e = e.substring(n.raw.length); else if (n = this.tokenizer.text(e)) e = e.substring(n.raw.length), i = a[a.length - 1], i && "text" === i.type ? (i.raw += "\n" + n.raw, i.text += "\n" + n.text, this.inlineQueue.pop(), this.inlineQueue[this.inlineQueue.length - 1].src = i.text) : a.push(n); else if (e) {
                    const n = "Infinite loop on byte: " + e.charCodeAt(0);
                    if (this.options.silent) {
                        console.error(n);
                        break
                    }
                    throw new Error(n)
                }
            }
            return this.state.top = !0, a
        }

        inline(e) {
            let n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : [];
            return this.inlineQueue.push({src: e, tokens: n}), n
        }

        inlineTokens(e) {
            let n, i, t, o, a, c, r = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : [], s = e;
            if (this.tokens.links) {
                const e = Object.keys(this.tokens.links);
                if (e.length > 0) for (; null != (o = this.tokenizer.rules.inline.reflinkSearch.exec(s));) e.includes(o[0].slice(o[0].lastIndexOf("[") + 1, -1)) && (s = s.slice(0, o.index) + "[" + "a".repeat(o[0].length - 2) + "]" + s.slice(this.tokenizer.rules.inline.reflinkSearch.lastIndex))
            }
            for (; null != (o = this.tokenizer.rules.inline.blockSkip.exec(s));) s = s.slice(0, o.index) + "[" + "a".repeat(o[0].length - 2) + "]" + s.slice(this.tokenizer.rules.inline.blockSkip.lastIndex);
            for (; null != (o = this.tokenizer.rules.inline.anyPunctuation.exec(s));) s = s.slice(0, o.index) + "++" + s.slice(this.tokenizer.rules.inline.anyPunctuation.lastIndex);
            for (; e;) if (a || (c = ""), a = !1, !(this.options.extensions && this.options.extensions.inline && this.options.extensions.inline.some((i => !!(n = i.call({lexer: this}, e, r)) && (e = e.substring(n.raw.length), r.push(n), !0))))) if (n = this.tokenizer.escape(e)) e = e.substring(n.raw.length), r.push(n); else if (n = this.tokenizer.tag(e)) e = e.substring(n.raw.length), i = r[r.length - 1], i && "text" === n.type && "text" === i.type ? (i.raw += n.raw, i.text += n.text) : r.push(n); else if (n = this.tokenizer.link(e)) e = e.substring(n.raw.length), r.push(n); else if (n = this.tokenizer.reflink(e, this.tokens.links)) e = e.substring(n.raw.length), i = r[r.length - 1], i && "text" === n.type && "text" === i.type ? (i.raw += n.raw, i.text += n.text) : r.push(n); else if (n = this.tokenizer.emStrong(e, s, c)) e = e.substring(n.raw.length), r.push(n); else if (n = this.tokenizer.codespan(e)) e = e.substring(n.raw.length), r.push(n); else if (n = this.tokenizer.br(e)) e = e.substring(n.raw.length), r.push(n); else if (n = this.tokenizer.del(e)) e = e.substring(n.raw.length), r.push(n); else if (n = this.tokenizer.autolink(e)) e = e.substring(n.raw.length), r.push(n); else if (this.state.inLink || !(n = this.tokenizer.url(e))) {
                if (t = e, this.options.extensions && this.options.extensions.startInline) {
                    let n = 1 / 0;
                    const i = e.slice(1);
                    let o;
                    this.options.extensions.startInline.forEach((e => {
                        o = e.call({lexer: this}, i), "number" == typeof o && o >= 0 && (n = Math.min(n, o))
                    })), n < 1 / 0 && n >= 0 && (t = e.substring(0, n + 1))
                }
                if (n = this.tokenizer.inlineText(t)) e = e.substring(n.raw.length), "_" !== n.raw.slice(-1) && (c = n.raw.slice(-1)), a = !0, i = r[r.length - 1], i && "text" === i.type ? (i.raw += n.raw, i.text += n.text) : r.push(n); else if (e) {
                    const n = "Infinite loop on byte: " + e.charCodeAt(0);
                    if (this.options.silent) {
                        console.error(n);
                        break
                    }
                    throw new Error(n)
                }
            } else e = e.substring(n.raw.length), r.push(n);
            return r
        }
    }

    class De {
        options;

        constructor(e) {
            this.options = e || H
        }

        code(e, n, i) {
            const t = (n || "").match(/^\S*/)?.[0];
            return e = e.replace(/\n$/, "") + "\n", t ? '<pre><code class="language-' + G(t) + '">' + (i ? e : G(e, !0)) + "</code></pre>\n" : "<pre><code>" + (i ? e : G(e, !0)) + "</code></pre>\n"
        }

        blockquote(e) {
            return `<blockquote>\n${e}</blockquote>\n`
        }

        html(e, n) {
            return e
        }

        heading(e, n, i) {
            return `<h${n}>${e}</h${n}>\n`
        }

        hr() {
            return "<hr>\n"
        }

        list(e, n, i) {
            const t = n ? "ol" : "ul";
            return "<" + t + (n && 1 !== i ? ' start="' + i + '"' : "") + ">\n" + e + "</" + t + ">\n"
        }

        listitem(e, n, i) {
            return `<li>${e}</li>\n`
        }

        checkbox(e) {
            return "<input " + (e ? 'checked="" ' : "") + 'disabled="" type="checkbox">'
        }

        paragraph(e) {
            return `<p>${e}</p>\n`
        }

        table(e, n) {
            return n && (n = `<tbody>${n}</tbody>`), "<table>\n<thead>\n" + e + "</thead>\n" + n + "</table>\n"
        }

        tablerow(e) {
            return `<tr>\n${e}</tr>\n`
        }

        tablecell(e, n) {
            const i = n.header ? "th" : "td";
            return (n.align ? `<${i} align="${n.align}">` : `<${i}>`) + e + `</${i}>\n`
        }

        strong(e) {
            return `<strong>${e}</strong>`
        }

        em(e) {
            return `<em>${e}</em>`
        }

        codespan(e) {
            return `<code>${e}</code>`
        }

        br() {
            return "<br>"
        }

        del(e) {
            return `<del>${e}</del>`
        }

        link(e, n, i) {
            const t = ee(e);
            if (null === t) return i;
            let o = '<a href="' + (e = t) + '"';
            return n && (o += ' title="' + n + '"'), o += ">" + i + "</a>", o
        }

        image(e, n, i) {
            const t = ee(e);
            if (null === t) return i;
            let o = `<img src="${e = t}" alt="${i}"`;
            return n && (o += ` title="${n}"`), o += ">", o
        }

        text(e) {
            return e
        }
    }

    class Ue {
        strong(e) {
            return e
        }

        em(e) {
            return e
        }

        codespan(e) {
            return e
        }

        del(e) {
            return e
        }

        html(e) {
            return e
        }

        text(e) {
            return e
        }

        link(e, n, i) {
            return "" + i
        }

        image(e, n, i) {
            return "" + i
        }

        br() {
            return ""
        }
    }

    class Ze {
        options;
        renderer;
        textRenderer;

        constructor(e) {
            this.options = e || H, this.options.renderer = this.options.renderer || new De, this.renderer = this.options.renderer, this.renderer.options = this.options, this.textRenderer = new Ue
        }

        static parse(e, n) {
            return new Ze(n).parse(e)
        }

        static parseInline(e, n) {
            return new Ze(n).parseInline(e)
        }

        parse(e) {
            let n = !(arguments.length > 1 && void 0 !== arguments[1]) || arguments[1], i = "";
            for (let t = 0; t < e.length; t++) {
                const o = e[t];
                if (this.options.extensions && this.options.extensions.renderers && this.options.extensions.renderers[o.type]) {
                    const e = o, n = this.options.extensions.renderers[e.type].call({parser: this}, e);
                    if (!1 !== n || !["space", "hr", "heading", "code", "table", "blockquote", "list", "html", "paragraph", "text"].includes(e.type)) {
                        i += n || "";
                        continue
                    }
                }
                switch (o.type) {
                    case"space":
                        continue;
                    case"hr":
                        i += this.renderer.hr();
                        continue;
                    case"heading": {
                        const e = o;
                        i += this.renderer.heading(this.parseInline(e.tokens), e.depth, X(this.parseInline(e.tokens, this.textRenderer)));
                        continue
                    }
                    case"code": {
                        const e = o;
                        i += this.renderer.code(e.text, e.lang, !!e.escaped);
                        continue
                    }
                    case"table": {
                        const e = o;
                        let n = "", t = "";
                        for (let n = 0; n < e.header.length; n++) t += this.renderer.tablecell(this.parseInline(e.header[n].tokens), {
                            header: !0,
                            align: e.align[n]
                        });
                        n += this.renderer.tablerow(t);
                        let a = "";
                        for (let n = 0; n < e.rows.length; n++) {
                            const i = e.rows[n];
                            t = "";
                            for (let n = 0; n < i.length; n++) t += this.renderer.tablecell(this.parseInline(i[n].tokens), {
                                header: !1,
                                align: e.align[n]
                            });
                            a += this.renderer.tablerow(t)
                        }
                        i += this.renderer.table(n, a);
                        continue
                    }
                    case"blockquote": {
                        const e = o, n = this.parse(e.tokens);
                        i += this.renderer.blockquote(n);
                        continue
                    }
                    case"list": {
                        const e = o, n = e.ordered, t = e.start, a = e.loose;
                        let c = "";
                        for (let n = 0; n < e.items.length; n++) {
                            const i = e.items[n], t = i.checked, o = i.task;
                            let r = "";
                            if (i.task) {
                                const e = this.renderer.checkbox(!!t);
                                a ? i.tokens.length > 0 && "paragraph" === i.tokens[0].type ? (i.tokens[0].text = e + " " + i.tokens[0].text, i.tokens[0].tokens && i.tokens[0].tokens.length > 0 && "text" === i.tokens[0].tokens[0].type && (i.tokens[0].tokens[0].text = e + " " + i.tokens[0].tokens[0].text)) : i.tokens.unshift({
                                    type: "text",
                                    text: e + " "
                                }) : r += e + " "
                            }
                            r += this.parse(i.tokens, a), c += this.renderer.listitem(r, o, !!t)
                        }
                        i += this.renderer.list(c, n, t);
                        continue
                    }
                    case"html": {
                        const e = o;
                        i += this.renderer.html(e.text, e.block);
                        continue
                    }
                    case"paragraph": {
                        const e = o;
                        i += this.renderer.paragraph(this.parseInline(e.tokens));
                        continue
                    }
                    case"text": {
                        let a = o, c = a.tokens ? this.parseInline(a.tokens) : a.text;
                        for (; t + 1 < e.length && "text" === e[t + 1].type;) a = e[++t], c += "\n" + (a.tokens ? this.parseInline(a.tokens) : a.text);
                        i += n ? this.renderer.paragraph(c) : c;
                        continue
                    }
                    default: {
                        const e = 'Token with "' + o.type + '" type was not found.';
                        if (this.options.silent) return console.error(e), "";
                        throw new Error(e)
                    }
                }
            }
            return i
        }

        parseInline(e, n) {
            n = n || this.renderer;
            let i = "";
            for (let t = 0; t < e.length; t++) {
                const o = e[t];
                if (this.options.extensions && this.options.extensions.renderers && this.options.extensions.renderers[o.type]) {
                    const e = this.options.extensions.renderers[o.type].call({parser: this}, o);
                    if (!1 !== e || !["escape", "html", "link", "image", "strong", "em", "codespan", "br", "del", "text"].includes(o.type)) {
                        i += e || "";
                        continue
                    }
                }
                switch (o.type) {
                    case"escape": {
                        const e = o;
                        i += n.text(e.text);
                        break
                    }
                    case"html": {
                        const e = o;
                        i += n.html(e.text);
                        break
                    }
                    case"link": {
                        const e = o;
                        i += n.link(e.href, e.title, this.parseInline(e.tokens, n));
                        break
                    }
                    case"image": {
                        const e = o;
                        i += n.image(e.href, e.title, e.text);
                        break
                    }
                    case"strong": {
                        const e = o;
                        i += n.strong(this.parseInline(e.tokens, n));
                        break
                    }
                    case"em": {
                        const e = o;
                        i += n.em(this.parseInline(e.tokens, n));
                        break
                    }
                    case"codespan": {
                        const e = o;
                        i += n.codespan(e.text);
                        break
                    }
                    case"br":
                        i += n.br();
                        break;
                    case"del": {
                        const e = o;
                        i += n.del(this.parseInline(e.tokens, n));
                        break
                    }
                    case"text": {
                        const e = o;
                        i += n.text(e.text);
                        break
                    }
                    default: {
                        const e = 'Token with "' + o.type + '" type was not found.';
                        if (this.options.silent) return console.error(e), "";
                        throw new Error(e)
                    }
                }
            }
            return i
        }
    }

    class Qe {
        options;

        constructor(e) {
            this.options = e || H
        }

        static passThroughHooks = new Set(["preprocess", "postprocess", "processAllTokens"]);

        preprocess(e) {
            return e
        }

        postprocess(e) {
            return e
        }

        processAllTokens(e) {
            return e
        }
    }

    const Ve = new class {
        defaults = {
            async: !1,
            breaks: !1,
            extensions: null,
            gfm: !0,
            hooks: null,
            pedantic: !1,
            renderer: null,
            silent: !1,
            tokenizer: null,
            walkTokens: null
        };
        options = this.setOptions;
        parse = this.#t(Be.lex, Ze.parse);
        parseInline = this.#t(Be.lexInline, Ze.parseInline);
        Parser = Ze;
        Renderer = De;
        TextRenderer = Ue;
        Lexer = Be;
        Tokenizer = ae;
        Hooks = Qe;

        constructor() {
            this.use(...arguments)
        }

        walkTokens(e, n) {
            let i = [];
            for (const t of e) switch (i = i.concat(n.call(this, t)), t.type) {
                case"table": {
                    const e = t;
                    for (const t of e.header) i = i.concat(this.walkTokens(t.tokens, n));
                    for (const t of e.rows) for (const e of t) i = i.concat(this.walkTokens(e.tokens, n));
                    break
                }
                case"list": {
                    const e = t;
                    i = i.concat(this.walkTokens(e.items, n));
                    break
                }
                default: {
                    const e = t;
                    this.defaults.extensions?.childTokens?.[e.type] ? this.defaults.extensions.childTokens[e.type].forEach((t => {
                        const o = e[t].flat(1 / 0);
                        i = i.concat(this.walkTokens(o, n))
                    })) : e.tokens && (i = i.concat(this.walkTokens(e.tokens, n)))
                }
            }
            return i
        }

        use() {
            const e = this.defaults.extensions || {renderers: {}, childTokens: {}};
            for (var n = arguments.length, i = new Array(n), t = 0; t < n; t++) i[t] = arguments[t];
            return i.forEach((n => {
                const i = {...n};
                if (i.async = this.defaults.async || i.async || !1, n.extensions && (n.extensions.forEach((n => {
                    if (!n.name) throw new Error("extension name required");
                    if ("renderer" in n) {
                        const i = e.renderers[n.name];
                        e.renderers[n.name] = i ? function () {
                            for (var e = arguments.length, t = new Array(e), o = 0; o < e; o++) t[o] = arguments[o];
                            let a = n.renderer.apply(this, t);
                            return !1 === a && (a = i.apply(this, t)), a
                        } : n.renderer
                    }
                    if ("tokenizer" in n) {
                        if (!n.level || "block" !== n.level && "inline" !== n.level) throw new Error("extension level must be 'block' or 'inline'");
                        const i = e[n.level];
                        i ? i.unshift(n.tokenizer) : e[n.level] = [n.tokenizer], n.start && ("block" === n.level ? e.startBlock ? e.startBlock.push(n.start) : e.startBlock = [n.start] : "inline" === n.level && (e.startInline ? e.startInline.push(n.start) : e.startInline = [n.start]))
                    }
                    "childTokens" in n && n.childTokens && (e.childTokens[n.name] = n.childTokens)
                })), i.extensions = e), n.renderer) {
                    const e = this.defaults.renderer || new De(this.defaults);
                    for (const i in n.renderer) {
                        if (!(i in e)) throw new Error(`renderer '${i}' does not exist`);
                        if ("options" === i) continue;
                        const t = i, o = n.renderer[t], a = e[t];
                        e[t] = function () {
                            for (var n = arguments.length, i = new Array(n), t = 0; t < n; t++) i[t] = arguments[t];
                            let c = o.apply(e, i);
                            return !1 === c && (c = a.apply(e, i)), c || ""
                        }
                    }
                    i.renderer = e
                }
                if (n.tokenizer) {
                    const e = this.defaults.tokenizer || new ae(this.defaults);
                    for (const i in n.tokenizer) {
                        if (!(i in e)) throw new Error(`tokenizer '${i}' does not exist`);
                        if (["options", "rules", "lexer"].includes(i)) continue;
                        const t = i, o = n.tokenizer[t], a = e[t];
                        e[t] = function () {
                            for (var n = arguments.length, i = new Array(n), t = 0; t < n; t++) i[t] = arguments[t];
                            let c = o.apply(e, i);
                            return !1 === c && (c = a.apply(e, i)), c
                        }
                    }
                    i.tokenizer = e
                }
                if (n.hooks) {
                    const e = this.defaults.hooks || new Qe;
                    for (const i in n.hooks) {
                        if (!(i in e)) throw new Error(`hook '${i}' does not exist`);
                        if ("options" === i) continue;
                        const t = i, o = n.hooks[t], a = e[t];
                        Qe.passThroughHooks.has(i) ? e[t] = n => {
                            if (this.defaults.async) return Promise.resolve(o.call(e, n)).then((n => a.call(e, n)));
                            const i = o.call(e, n);
                            return a.call(e, i)
                        } : e[t] = function () {
                            for (var n = arguments.length, i = new Array(n), t = 0; t < n; t++) i[t] = arguments[t];
                            let c = o.apply(e, i);
                            return !1 === c && (c = a.apply(e, i)), c
                        }
                    }
                    i.hooks = e
                }
                if (n.walkTokens) {
                    const e = this.defaults.walkTokens, t = n.walkTokens;
                    i.walkTokens = function (n) {
                        let i = [];
                        return i.push(t.call(this, n)), e && (i = i.concat(e.call(this, n))), i
                    }
                }
                this.defaults = {...this.defaults, ...i}
            })), this
        }

        setOptions(e) {
            return this.defaults = {...this.defaults, ...e}, this
        }

        lexer(e, n) {
            return Be.lex(e, n ?? this.defaults)
        }

        parser(e, n) {
            return Ze.parse(e, n ?? this.defaults)
        }

        #t(e, n) {
            return (i, t) => {
                const o = {...t}, a = {...this.defaults, ...o};
                !0 === this.defaults.async && !1 === o.async && (a.silent || console.warn("marked(): The async option was set to true by an extension. The async: false option sent to parse will be ignored."), a.async = !0);
                const c = this.#o(!!a.silent, !!a.async);
                if (null == i) return c(new Error("marked(): input parameter is undefined or null"));
                if ("string" != typeof i) return c(new Error("marked(): input parameter is of type " + Object.prototype.toString.call(i) + ", string expected"));
                if (a.hooks && (a.hooks.options = a), a.async) return Promise.resolve(a.hooks ? a.hooks.preprocess(i) : i).then((n => e(n, a))).then((e => a.hooks ? a.hooks.processAllTokens(e) : e)).then((e => a.walkTokens ? Promise.all(this.walkTokens(e, a.walkTokens)).then((() => e)) : e)).then((e => n(e, a))).then((e => a.hooks ? a.hooks.postprocess(e) : e)).catch(c);
                try {
                    a.hooks && (i = a.hooks.preprocess(i));
                    let t = e(i, a);
                    a.hooks && (t = a.hooks.processAllTokens(t)), a.walkTokens && this.walkTokens(t, a.walkTokens);
                    let o = n(t, a);
                    return a.hooks && (o = a.hooks.postprocess(o)), o
                } catch (e) {
                    return c(e)
                }
            }
        }

        #o(e, n) {
            return i => {
                if (i.message += "\nPlease report this to https://github.com/markedjs/marked.", e) {
                    const e = "<p>An error occurred:</p><pre>" + G(i.message + "", !0) + "</pre>";
                    return n ? Promise.resolve(e) : e
                }
                if (n) return Promise.reject(i);
                throw i
            }
        }
    };

    function We(e, n) {
        return Ve.parse(e, n)
    }

    function Ge(e) {
        let n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : '<ul class="app-sub-sidebar">{inner}</ul>';
        if (!e || !e.length) return "";
        let i = "";
        return e.forEach((e => {
            const t = e.title.replace(/(<([^>]+)>)/g, "");
            i += `<li><a class="section-link" href="${e.slug}" title="${t}">${e.title}</a></li>`, e.children && (i += Ge(e.children, n))
        })), n.replace("{inner}", i)
    }

    function Ye(e, n) {
        return `<p class="${e}">${n.slice(5).trim()}</p>`
    }

    function Xe(e, n) {
        const i = [], t = {};
        return e.forEach((e => {
            const o = e.level || 1, a = o - 1;
            o > n || (t[a] ? t[a].children = [...t[a].children || [], e] : i.push(e), t[o] = e)
        })), i
    }

    We.options = We.setOptions = function (e) {
        return Ve.setOptions(e), We.defaults = Ve.defaults, B(We.defaults), We
    }, We.getDefaults = N, We.defaults = H, We.use = function () {
        return Ve.use(...arguments), We.defaults = Ve.defaults, B(We.defaults), We
    }, We.walkTokens = function (e, n) {
        return Ve.walkTokens(e, n)
    }, We.parseInline = Ve.parseInline, We.Parser = Ze, We.parser = Ze.parse, We.Renderer = De, We.TextRenderer = Ue, We.Lexer = Be, We.lexer = Be.lex, We.Tokenizer = ae, We.Hooks = Qe, We.parse = We, We.options, We.setOptions, We.use, We.walkTokens, We.parseInline, Ze.parse, Be.lex;
    let Je = {};
    const Ke = /[\u2000-\u206F\u2E00-\u2E7F\\'!"#$%&()*+,./:;<=>?@[\]^`{|}~]/g;

    function en(e) {
        return e.toLowerCase()
    }

    function nn(e) {
        if ("string" != typeof e) return "";
        let n = e.trim().replace(/[A-Z]+/g, en).replace(/<[^>]+>/g, "").replace(Ke, "").replace(/\s/g, "-").replace(/-+/g, "-").replace(/^(\d)/, "_$1"),
            i = Je[n];
        return i = Object.keys(Je).includes(n) ? i + 1 : 0, Je[n] = i, i && (n = n + "-" + i), n
    }

    nn.clear = function () {
        Je = {}
    };
    var tn = {
        baseURL: "https://github.githubassets.com/images/icons/emoji/", data: {
            100: "unicode/1f4af.png?v8",
            1234: "unicode/1f522.png?v8",
            "+1": "unicode/1f44d.png?v8",
            "-1": "unicode/1f44e.png?v8",
            "1st_place_medal": "unicode/1f947.png?v8",
            "2nd_place_medal": "unicode/1f948.png?v8",
            "3rd_place_medal": "unicode/1f949.png?v8",
            "8ball": "unicode/1f3b1.png?v8",
            a: "unicode/1f170.png?v8",
            ab: "unicode/1f18e.png?v8",
            abacus: "unicode/1f9ee.png?v8",
            abc: "unicode/1f524.png?v8",
            abcd: "unicode/1f521.png?v8",
            accept: "unicode/1f251.png?v8",
            accessibility: "accessibility.png?v8",
            accordion: "unicode/1fa97.png?v8",
            adhesive_bandage: "unicode/1fa79.png?v8",
            adult: "unicode/1f9d1.png?v8",
            aerial_tramway: "unicode/1f6a1.png?v8",
            afghanistan: "unicode/1f1e6-1f1eb.png?v8",
            airplane: "unicode/2708.png?v8",
            aland_islands: "unicode/1f1e6-1f1fd.png?v8",
            alarm_clock: "unicode/23f0.png?v8",
            albania: "unicode/1f1e6-1f1f1.png?v8",
            alembic: "unicode/2697.png?v8",
            algeria: "unicode/1f1e9-1f1ff.png?v8",
            alien: "unicode/1f47d.png?v8",
            ambulance: "unicode/1f691.png?v8",
            american_samoa: "unicode/1f1e6-1f1f8.png?v8",
            amphora: "unicode/1f3fa.png?v8",
            anatomical_heart: "unicode/1fac0.png?v8",
            anchor: "unicode/2693.png?v8",
            andorra: "unicode/1f1e6-1f1e9.png?v8",
            angel: "unicode/1f47c.png?v8",
            anger: "unicode/1f4a2.png?v8",
            angola: "unicode/1f1e6-1f1f4.png?v8",
            angry: "unicode/1f620.png?v8",
            anguilla: "unicode/1f1e6-1f1ee.png?v8",
            anguished: "unicode/1f627.png?v8",
            ant: "unicode/1f41c.png?v8",
            antarctica: "unicode/1f1e6-1f1f6.png?v8",
            antigua_barbuda: "unicode/1f1e6-1f1ec.png?v8",
            apple: "unicode/1f34e.png?v8",
            aquarius: "unicode/2652.png?v8",
            argentina: "unicode/1f1e6-1f1f7.png?v8",
            aries: "unicode/2648.png?v8",
            armenia: "unicode/1f1e6-1f1f2.png?v8",
            arrow_backward: "unicode/25c0.png?v8",
            arrow_double_down: "unicode/23ec.png?v8",
            arrow_double_up: "unicode/23eb.png?v8",
            arrow_down: "unicode/2b07.png?v8",
            arrow_down_small: "unicode/1f53d.png?v8",
            arrow_forward: "unicode/25b6.png?v8",
            arrow_heading_down: "unicode/2935.png?v8",
            arrow_heading_up: "unicode/2934.png?v8",
            arrow_left: "unicode/2b05.png?v8",
            arrow_lower_left: "unicode/2199.png?v8",
            arrow_lower_right: "unicode/2198.png?v8",
            arrow_right: "unicode/27a1.png?v8",
            arrow_right_hook: "unicode/21aa.png?v8",
            arrow_up: "unicode/2b06.png?v8",
            arrow_up_down: "unicode/2195.png?v8",
            arrow_up_small: "unicode/1f53c.png?v8",
            arrow_upper_left: "unicode/2196.png?v8",
            arrow_upper_right: "unicode/2197.png?v8",
            arrows_clockwise: "unicode/1f503.png?v8",
            arrows_counterclockwise: "unicode/1f504.png?v8",
            art: "unicode/1f3a8.png?v8",
            articulated_lorry: "unicode/1f69b.png?v8",
            artificial_satellite: "unicode/1f6f0.png?v8",
            artist: "unicode/1f9d1-1f3a8.png?v8",
            aruba: "unicode/1f1e6-1f1fc.png?v8",
            ascension_island: "unicode/1f1e6-1f1e8.png?v8",
            asterisk: "unicode/002a-20e3.png?v8",
            astonished: "unicode/1f632.png?v8",
            astronaut: "unicode/1f9d1-1f680.png?v8",
            athletic_shoe: "unicode/1f45f.png?v8",
            atm: "unicode/1f3e7.png?v8",
            atom: "atom.png?v8",
            atom_symbol: "unicode/269b.png?v8",
            australia: "unicode/1f1e6-1f1fa.png?v8",
            austria: "unicode/1f1e6-1f1f9.png?v8",
            auto_rickshaw: "unicode/1f6fa.png?v8",
            avocado: "unicode/1f951.png?v8",
            axe: "unicode/1fa93.png?v8",
            azerbaijan: "unicode/1f1e6-1f1ff.png?v8",
            b: "unicode/1f171.png?v8",
            baby: "unicode/1f476.png?v8",
            baby_bottle: "unicode/1f37c.png?v8",
            baby_chick: "unicode/1f424.png?v8",
            baby_symbol: "unicode/1f6bc.png?v8",
            back: "unicode/1f519.png?v8",
            bacon: "unicode/1f953.png?v8",
            badger: "unicode/1f9a1.png?v8",
            badminton: "unicode/1f3f8.png?v8",
            bagel: "unicode/1f96f.png?v8",
            baggage_claim: "unicode/1f6c4.png?v8",
            baguette_bread: "unicode/1f956.png?v8",
            bahamas: "unicode/1f1e7-1f1f8.png?v8",
            bahrain: "unicode/1f1e7-1f1ed.png?v8",
            balance_scale: "unicode/2696.png?v8",
            bald_man: "unicode/1f468-1f9b2.png?v8",
            bald_woman: "unicode/1f469-1f9b2.png?v8",
            ballet_shoes: "unicode/1fa70.png?v8",
            balloon: "unicode/1f388.png?v8",
            ballot_box: "unicode/1f5f3.png?v8",
            ballot_box_with_check: "unicode/2611.png?v8",
            bamboo: "unicode/1f38d.png?v8",
            banana: "unicode/1f34c.png?v8",
            bangbang: "unicode/203c.png?v8",
            bangladesh: "unicode/1f1e7-1f1e9.png?v8",
            banjo: "unicode/1fa95.png?v8",
            bank: "unicode/1f3e6.png?v8",
            bar_chart: "unicode/1f4ca.png?v8",
            barbados: "unicode/1f1e7-1f1e7.png?v8",
            barber: "unicode/1f488.png?v8",
            baseball: "unicode/26be.png?v8",
            basecamp: "basecamp.png?v8",
            basecampy: "basecampy.png?v8",
            basket: "unicode/1f9fa.png?v8",
            basketball: "unicode/1f3c0.png?v8",
            basketball_man: "unicode/26f9-2642.png?v8",
            basketball_woman: "unicode/26f9-2640.png?v8",
            bat: "unicode/1f987.png?v8",
            bath: "unicode/1f6c0.png?v8",
            bathtub: "unicode/1f6c1.png?v8",
            battery: "unicode/1f50b.png?v8",
            beach_umbrella: "unicode/1f3d6.png?v8",
            beans: "unicode/1fad8.png?v8",
            bear: "unicode/1f43b.png?v8",
            bearded_person: "unicode/1f9d4.png?v8",
            beaver: "unicode/1f9ab.png?v8",
            bed: "unicode/1f6cf.png?v8",
            bee: "unicode/1f41d.png?v8",
            beer: "unicode/1f37a.png?v8",
            beers: "unicode/1f37b.png?v8",
            beetle: "unicode/1fab2.png?v8",
            beginner: "unicode/1f530.png?v8",
            belarus: "unicode/1f1e7-1f1fe.png?v8",
            belgium: "unicode/1f1e7-1f1ea.png?v8",
            belize: "unicode/1f1e7-1f1ff.png?v8",
            bell: "unicode/1f514.png?v8",
            bell_pepper: "unicode/1fad1.png?v8",
            bellhop_bell: "unicode/1f6ce.png?v8",
            benin: "unicode/1f1e7-1f1ef.png?v8",
            bento: "unicode/1f371.png?v8",
            bermuda: "unicode/1f1e7-1f1f2.png?v8",
            beverage_box: "unicode/1f9c3.png?v8",
            bhutan: "unicode/1f1e7-1f1f9.png?v8",
            bicyclist: "unicode/1f6b4.png?v8",
            bike: "unicode/1f6b2.png?v8",
            biking_man: "unicode/1f6b4-2642.png?v8",
            biking_woman: "unicode/1f6b4-2640.png?v8",
            bikini: "unicode/1f459.png?v8",
            billed_cap: "unicode/1f9e2.png?v8",
            biohazard: "unicode/2623.png?v8",
            bird: "unicode/1f426.png?v8",
            birthday: "unicode/1f382.png?v8",
            bison: "unicode/1f9ac.png?v8",
            biting_lip: "unicode/1fae6.png?v8",
            black_bird: "unicode/1f426-2b1b.png?v8",
            black_cat: "unicode/1f408-2b1b.png?v8",
            black_circle: "unicode/26ab.png?v8",
            black_flag: "unicode/1f3f4.png?v8",
            black_heart: "unicode/1f5a4.png?v8",
            black_joker: "unicode/1f0cf.png?v8",
            black_large_square: "unicode/2b1b.png?v8",
            black_medium_small_square: "unicode/25fe.png?v8",
            black_medium_square: "unicode/25fc.png?v8",
            black_nib: "unicode/2712.png?v8",
            black_small_square: "unicode/25aa.png?v8",
            black_square_button: "unicode/1f532.png?v8",
            blond_haired_man: "unicode/1f471-2642.png?v8",
            blond_haired_person: "unicode/1f471.png?v8",
            blond_haired_woman: "unicode/1f471-2640.png?v8",
            blonde_woman: "unicode/1f471-2640.png?v8",
            blossom: "unicode/1f33c.png?v8",
            blowfish: "unicode/1f421.png?v8",
            blue_book: "unicode/1f4d8.png?v8",
            blue_car: "unicode/1f699.png?v8",
            blue_heart: "unicode/1f499.png?v8",
            blue_square: "unicode/1f7e6.png?v8",
            blueberries: "unicode/1fad0.png?v8",
            blush: "unicode/1f60a.png?v8",
            boar: "unicode/1f417.png?v8",
            boat: "unicode/26f5.png?v8",
            bolivia: "unicode/1f1e7-1f1f4.png?v8",
            bomb: "unicode/1f4a3.png?v8",
            bone: "unicode/1f9b4.png?v8",
            book: "unicode/1f4d6.png?v8",
            bookmark: "unicode/1f516.png?v8",
            bookmark_tabs: "unicode/1f4d1.png?v8",
            books: "unicode/1f4da.png?v8",
            boom: "unicode/1f4a5.png?v8",
            boomerang: "unicode/1fa83.png?v8",
            boot: "unicode/1f462.png?v8",
            bosnia_herzegovina: "unicode/1f1e7-1f1e6.png?v8",
            botswana: "unicode/1f1e7-1f1fc.png?v8",
            bouncing_ball_man: "unicode/26f9-2642.png?v8",
            bouncing_ball_person: "unicode/26f9.png?v8",
            bouncing_ball_woman: "unicode/26f9-2640.png?v8",
            bouquet: "unicode/1f490.png?v8",
            bouvet_island: "unicode/1f1e7-1f1fb.png?v8",
            bow: "unicode/1f647.png?v8",
            bow_and_arrow: "unicode/1f3f9.png?v8",
            bowing_man: "unicode/1f647-2642.png?v8",
            bowing_woman: "unicode/1f647-2640.png?v8",
            bowl_with_spoon: "unicode/1f963.png?v8",
            bowling: "unicode/1f3b3.png?v8",
            bowtie: "bowtie.png?v8",
            boxing_glove: "unicode/1f94a.png?v8",
            boy: "unicode/1f466.png?v8",
            brain: "unicode/1f9e0.png?v8",
            brazil: "unicode/1f1e7-1f1f7.png?v8",
            bread: "unicode/1f35e.png?v8",
            breast_feeding: "unicode/1f931.png?v8",
            bricks: "unicode/1f9f1.png?v8",
            bride_with_veil: "unicode/1f470-2640.png?v8",
            bridge_at_night: "unicode/1f309.png?v8",
            briefcase: "unicode/1f4bc.png?v8",
            british_indian_ocean_territory: "unicode/1f1ee-1f1f4.png?v8",
            british_virgin_islands: "unicode/1f1fb-1f1ec.png?v8",
            broccoli: "unicode/1f966.png?v8",
            broken_heart: "unicode/1f494.png?v8",
            broom: "unicode/1f9f9.png?v8",
            brown_circle: "unicode/1f7e4.png?v8",
            brown_heart: "unicode/1f90e.png?v8",
            brown_square: "unicode/1f7eb.png?v8",
            brunei: "unicode/1f1e7-1f1f3.png?v8",
            bubble_tea: "unicode/1f9cb.png?v8",
            bubbles: "unicode/1fae7.png?v8",
            bucket: "unicode/1faa3.png?v8",
            bug: "unicode/1f41b.png?v8",
            building_construction: "unicode/1f3d7.png?v8",
            bulb: "unicode/1f4a1.png?v8",
            bulgaria: "unicode/1f1e7-1f1ec.png?v8",
            bullettrain_front: "unicode/1f685.png?v8",
            bullettrain_side: "unicode/1f684.png?v8",
            burkina_faso: "unicode/1f1e7-1f1eb.png?v8",
            burrito: "unicode/1f32f.png?v8",
            burundi: "unicode/1f1e7-1f1ee.png?v8",
            bus: "unicode/1f68c.png?v8",
            business_suit_levitating: "unicode/1f574.png?v8",
            busstop: "unicode/1f68f.png?v8",
            bust_in_silhouette: "unicode/1f464.png?v8",
            busts_in_silhouette: "unicode/1f465.png?v8",
            butter: "unicode/1f9c8.png?v8",
            butterfly: "unicode/1f98b.png?v8",
            cactus: "unicode/1f335.png?v8",
            cake: "unicode/1f370.png?v8",
            calendar: "unicode/1f4c6.png?v8",
            call_me_hand: "unicode/1f919.png?v8",
            calling: "unicode/1f4f2.png?v8",
            cambodia: "unicode/1f1f0-1f1ed.png?v8",
            camel: "unicode/1f42b.png?v8",
            camera: "unicode/1f4f7.png?v8",
            camera_flash: "unicode/1f4f8.png?v8",
            cameroon: "unicode/1f1e8-1f1f2.png?v8",
            camping: "unicode/1f3d5.png?v8",
            canada: "unicode/1f1e8-1f1e6.png?v8",
            canary_islands: "unicode/1f1ee-1f1e8.png?v8",
            cancer: "unicode/264b.png?v8",
            candle: "unicode/1f56f.png?v8",
            candy: "unicode/1f36c.png?v8",
            canned_food: "unicode/1f96b.png?v8",
            canoe: "unicode/1f6f6.png?v8",
            cape_verde: "unicode/1f1e8-1f1fb.png?v8",
            capital_abcd: "unicode/1f520.png?v8",
            capricorn: "unicode/2651.png?v8",
            car: "unicode/1f697.png?v8",
            card_file_box: "unicode/1f5c3.png?v8",
            card_index: "unicode/1f4c7.png?v8",
            card_index_dividers: "unicode/1f5c2.png?v8",
            caribbean_netherlands: "unicode/1f1e7-1f1f6.png?v8",
            carousel_horse: "unicode/1f3a0.png?v8",
            carpentry_saw: "unicode/1fa9a.png?v8",
            carrot: "unicode/1f955.png?v8",
            cartwheeling: "unicode/1f938.png?v8",
            cat: "unicode/1f431.png?v8",
            cat2: "unicode/1f408.png?v8",
            cayman_islands: "unicode/1f1f0-1f1fe.png?v8",
            cd: "unicode/1f4bf.png?v8",
            central_african_republic: "unicode/1f1e8-1f1eb.png?v8",
            ceuta_melilla: "unicode/1f1ea-1f1e6.png?v8",
            chad: "unicode/1f1f9-1f1e9.png?v8",
            chains: "unicode/26d3.png?v8",
            chair: "unicode/1fa91.png?v8",
            champagne: "unicode/1f37e.png?v8",
            chart: "unicode/1f4b9.png?v8",
            chart_with_downwards_trend: "unicode/1f4c9.png?v8",
            chart_with_upwards_trend: "unicode/1f4c8.png?v8",
            checkered_flag: "unicode/1f3c1.png?v8",
            cheese: "unicode/1f9c0.png?v8",
            cherries: "unicode/1f352.png?v8",
            cherry_blossom: "unicode/1f338.png?v8",
            chess_pawn: "unicode/265f.png?v8",
            chestnut: "unicode/1f330.png?v8",
            chicken: "unicode/1f414.png?v8",
            child: "unicode/1f9d2.png?v8",
            children_crossing: "unicode/1f6b8.png?v8",
            chile: "unicode/1f1e8-1f1f1.png?v8",
            chipmunk: "unicode/1f43f.png?v8",
            chocolate_bar: "unicode/1f36b.png?v8",
            chopsticks: "unicode/1f962.png?v8",
            christmas_island: "unicode/1f1e8-1f1fd.png?v8",
            christmas_tree: "unicode/1f384.png?v8",
            church: "unicode/26ea.png?v8",
            cinema: "unicode/1f3a6.png?v8",
            circus_tent: "unicode/1f3aa.png?v8",
            city_sunrise: "unicode/1f307.png?v8",
            city_sunset: "unicode/1f306.png?v8",
            cityscape: "unicode/1f3d9.png?v8",
            cl: "unicode/1f191.png?v8",
            clamp: "unicode/1f5dc.png?v8",
            clap: "unicode/1f44f.png?v8",
            clapper: "unicode/1f3ac.png?v8",
            classical_building: "unicode/1f3db.png?v8",
            climbing: "unicode/1f9d7.png?v8",
            climbing_man: "unicode/1f9d7-2642.png?v8",
            climbing_woman: "unicode/1f9d7-2640.png?v8",
            clinking_glasses: "unicode/1f942.png?v8",
            clipboard: "unicode/1f4cb.png?v8",
            clipperton_island: "unicode/1f1e8-1f1f5.png?v8",
            clock1: "unicode/1f550.png?v8",
            clock10: "unicode/1f559.png?v8",
            clock1030: "unicode/1f565.png?v8",
            clock11: "unicode/1f55a.png?v8",
            clock1130: "unicode/1f566.png?v8",
            clock12: "unicode/1f55b.png?v8",
            clock1230: "unicode/1f567.png?v8",
            clock130: "unicode/1f55c.png?v8",
            clock2: "unicode/1f551.png?v8",
            clock230: "unicode/1f55d.png?v8",
            clock3: "unicode/1f552.png?v8",
            clock330: "unicode/1f55e.png?v8",
            clock4: "unicode/1f553.png?v8",
            clock430: "unicode/1f55f.png?v8",
            clock5: "unicode/1f554.png?v8",
            clock530: "unicode/1f560.png?v8",
            clock6: "unicode/1f555.png?v8",
            clock630: "unicode/1f561.png?v8",
            clock7: "unicode/1f556.png?v8",
            clock730: "unicode/1f562.png?v8",
            clock8: "unicode/1f557.png?v8",
            clock830: "unicode/1f563.png?v8",
            clock9: "unicode/1f558.png?v8",
            clock930: "unicode/1f564.png?v8",
            closed_book: "unicode/1f4d5.png?v8",
            closed_lock_with_key: "unicode/1f510.png?v8",
            closed_umbrella: "unicode/1f302.png?v8",
            cloud: "unicode/2601.png?v8",
            cloud_with_lightning: "unicode/1f329.png?v8",
            cloud_with_lightning_and_rain: "unicode/26c8.png?v8",
            cloud_with_rain: "unicode/1f327.png?v8",
            cloud_with_snow: "unicode/1f328.png?v8",
            clown_face: "unicode/1f921.png?v8",
            clubs: "unicode/2663.png?v8",
            cn: "unicode/1f1e8-1f1f3.png?v8",
            coat: "unicode/1f9e5.png?v8",
            cockroach: "unicode/1fab3.png?v8",
            cocktail: "unicode/1f378.png?v8",
            coconut: "unicode/1f965.png?v8",
            cocos_islands: "unicode/1f1e8-1f1e8.png?v8",
            coffee: "unicode/2615.png?v8",
            coffin: "unicode/26b0.png?v8",
            coin: "unicode/1fa99.png?v8",
            cold_face: "unicode/1f976.png?v8",
            cold_sweat: "unicode/1f630.png?v8",
            collision: "unicode/1f4a5.png?v8",
            colombia: "unicode/1f1e8-1f1f4.png?v8",
            comet: "unicode/2604.png?v8",
            comoros: "unicode/1f1f0-1f1f2.png?v8",
            compass: "unicode/1f9ed.png?v8",
            computer: "unicode/1f4bb.png?v8",
            computer_mouse: "unicode/1f5b1.png?v8",
            confetti_ball: "unicode/1f38a.png?v8",
            confounded: "unicode/1f616.png?v8",
            confused: "unicode/1f615.png?v8",
            congo_brazzaville: "unicode/1f1e8-1f1ec.png?v8",
            congo_kinshasa: "unicode/1f1e8-1f1e9.png?v8",
            congratulations: "unicode/3297.png?v8",
            construction: "unicode/1f6a7.png?v8",
            construction_worker: "unicode/1f477.png?v8",
            construction_worker_man: "unicode/1f477-2642.png?v8",
            construction_worker_woman: "unicode/1f477-2640.png?v8",
            control_knobs: "unicode/1f39b.png?v8",
            convenience_store: "unicode/1f3ea.png?v8",
            cook: "unicode/1f9d1-1f373.png?v8",
            cook_islands: "unicode/1f1e8-1f1f0.png?v8",
            cookie: "unicode/1f36a.png?v8",
            cool: "unicode/1f192.png?v8",
            cop: "unicode/1f46e.png?v8",
            copyright: "unicode/00a9.png?v8",
            coral: "unicode/1fab8.png?v8",
            corn: "unicode/1f33d.png?v8",
            costa_rica: "unicode/1f1e8-1f1f7.png?v8",
            cote_divoire: "unicode/1f1e8-1f1ee.png?v8",
            couch_and_lamp: "unicode/1f6cb.png?v8",
            couple: "unicode/1f46b.png?v8",
            couple_with_heart: "unicode/1f491.png?v8",
            couple_with_heart_man_man: "unicode/1f468-2764-1f468.png?v8",
            couple_with_heart_woman_man: "unicode/1f469-2764-1f468.png?v8",
            couple_with_heart_woman_woman: "unicode/1f469-2764-1f469.png?v8",
            couplekiss: "unicode/1f48f.png?v8",
            couplekiss_man_man: "unicode/1f468-2764-1f48b-1f468.png?v8",
            couplekiss_man_woman: "unicode/1f469-2764-1f48b-1f468.png?v8",
            couplekiss_woman_woman: "unicode/1f469-2764-1f48b-1f469.png?v8",
            cow: "unicode/1f42e.png?v8",
            cow2: "unicode/1f404.png?v8",
            cowboy_hat_face: "unicode/1f920.png?v8",
            crab: "unicode/1f980.png?v8",
            crayon: "unicode/1f58d.png?v8",
            credit_card: "unicode/1f4b3.png?v8",
            crescent_moon: "unicode/1f319.png?v8",
            cricket: "unicode/1f997.png?v8",
            cricket_game: "unicode/1f3cf.png?v8",
            croatia: "unicode/1f1ed-1f1f7.png?v8",
            crocodile: "unicode/1f40a.png?v8",
            croissant: "unicode/1f950.png?v8",
            crossed_fingers: "unicode/1f91e.png?v8",
            crossed_flags: "unicode/1f38c.png?v8",
            crossed_swords: "unicode/2694.png?v8",
            crown: "unicode/1f451.png?v8",
            crutch: "unicode/1fa7c.png?v8",
            cry: "unicode/1f622.png?v8",
            crying_cat_face: "unicode/1f63f.png?v8",
            crystal_ball: "unicode/1f52e.png?v8",
            cuba: "unicode/1f1e8-1f1fa.png?v8",
            cucumber: "unicode/1f952.png?v8",
            cup_with_straw: "unicode/1f964.png?v8",
            cupcake: "unicode/1f9c1.png?v8",
            cupid: "unicode/1f498.png?v8",
            curacao: "unicode/1f1e8-1f1fc.png?v8",
            curling_stone: "unicode/1f94c.png?v8",
            curly_haired_man: "unicode/1f468-1f9b1.png?v8",
            curly_haired_woman: "unicode/1f469-1f9b1.png?v8",
            curly_loop: "unicode/27b0.png?v8",
            currency_exchange: "unicode/1f4b1.png?v8",
            curry: "unicode/1f35b.png?v8",
            cursing_face: "unicode/1f92c.png?v8",
            custard: "unicode/1f36e.png?v8",
            customs: "unicode/1f6c3.png?v8",
            cut_of_meat: "unicode/1f969.png?v8",
            cyclone: "unicode/1f300.png?v8",
            cyprus: "unicode/1f1e8-1f1fe.png?v8",
            czech_republic: "unicode/1f1e8-1f1ff.png?v8",
            dagger: "unicode/1f5e1.png?v8",
            dancer: "unicode/1f483.png?v8",
            dancers: "unicode/1f46f.png?v8",
            dancing_men: "unicode/1f46f-2642.png?v8",
            dancing_women: "unicode/1f46f-2640.png?v8",
            dango: "unicode/1f361.png?v8",
            dark_sunglasses: "unicode/1f576.png?v8",
            dart: "unicode/1f3af.png?v8",
            dash: "unicode/1f4a8.png?v8",
            date: "unicode/1f4c5.png?v8",
            de: "unicode/1f1e9-1f1ea.png?v8",
            deaf_man: "unicode/1f9cf-2642.png?v8",
            deaf_person: "unicode/1f9cf.png?v8",
            deaf_woman: "unicode/1f9cf-2640.png?v8",
            deciduous_tree: "unicode/1f333.png?v8",
            deer: "unicode/1f98c.png?v8",
            denmark: "unicode/1f1e9-1f1f0.png?v8",
            department_store: "unicode/1f3ec.png?v8",
            dependabot: "dependabot.png?v8",
            derelict_house: "unicode/1f3da.png?v8",
            desert: "unicode/1f3dc.png?v8",
            desert_island: "unicode/1f3dd.png?v8",
            desktop_computer: "unicode/1f5a5.png?v8",
            detective: "unicode/1f575.png?v8",
            diamond_shape_with_a_dot_inside: "unicode/1f4a0.png?v8",
            diamonds: "unicode/2666.png?v8",
            diego_garcia: "unicode/1f1e9-1f1ec.png?v8",
            disappointed: "unicode/1f61e.png?v8",
            disappointed_relieved: "unicode/1f625.png?v8",
            disguised_face: "unicode/1f978.png?v8",
            diving_mask: "unicode/1f93f.png?v8",
            diya_lamp: "unicode/1fa94.png?v8",
            dizzy: "unicode/1f4ab.png?v8",
            dizzy_face: "unicode/1f635.png?v8",
            djibouti: "unicode/1f1e9-1f1ef.png?v8",
            dna: "unicode/1f9ec.png?v8",
            do_not_litter: "unicode/1f6af.png?v8",
            dodo: "unicode/1f9a4.png?v8",
            dog: "unicode/1f436.png?v8",
            dog2: "unicode/1f415.png?v8",
            dollar: "unicode/1f4b5.png?v8",
            dolls: "unicode/1f38e.png?v8",
            dolphin: "unicode/1f42c.png?v8",
            dominica: "unicode/1f1e9-1f1f2.png?v8",
            dominican_republic: "unicode/1f1e9-1f1f4.png?v8",
            donkey: "unicode/1facf.png?v8",
            door: "unicode/1f6aa.png?v8",
            dotted_line_face: "unicode/1fae5.png?v8",
            doughnut: "unicode/1f369.png?v8",
            dove: "unicode/1f54a.png?v8",
            dragon: "unicode/1f409.png?v8",
            dragon_face: "unicode/1f432.png?v8",
            dress: "unicode/1f457.png?v8",
            dromedary_camel: "unicode/1f42a.png?v8",
            drooling_face: "unicode/1f924.png?v8",
            drop_of_blood: "unicode/1fa78.png?v8",
            droplet: "unicode/1f4a7.png?v8",
            drum: "unicode/1f941.png?v8",
            duck: "unicode/1f986.png?v8",
            dumpling: "unicode/1f95f.png?v8",
            dvd: "unicode/1f4c0.png?v8",
            "e-mail": "unicode/1f4e7.png?v8",
            eagle: "unicode/1f985.png?v8",
            ear: "unicode/1f442.png?v8",
            ear_of_rice: "unicode/1f33e.png?v8",
            ear_with_hearing_aid: "unicode/1f9bb.png?v8",
            earth_africa: "unicode/1f30d.png?v8",
            earth_americas: "unicode/1f30e.png?v8",
            earth_asia: "unicode/1f30f.png?v8",
            ecuador: "unicode/1f1ea-1f1e8.png?v8",
            egg: "unicode/1f95a.png?v8",
            eggplant: "unicode/1f346.png?v8",
            egypt: "unicode/1f1ea-1f1ec.png?v8",
            eight: "unicode/0038-20e3.png?v8",
            eight_pointed_black_star: "unicode/2734.png?v8",
            eight_spoked_asterisk: "unicode/2733.png?v8",
            eject_button: "unicode/23cf.png?v8",
            el_salvador: "unicode/1f1f8-1f1fb.png?v8",
            electric_plug: "unicode/1f50c.png?v8",
            electron: "electron.png?v8",
            elephant: "unicode/1f418.png?v8",
            elevator: "unicode/1f6d7.png?v8",
            elf: "unicode/1f9dd.png?v8",
            elf_man: "unicode/1f9dd-2642.png?v8",
            elf_woman: "unicode/1f9dd-2640.png?v8",
            email: "unicode/1f4e7.png?v8",
            empty_nest: "unicode/1fab9.png?v8",
            end: "unicode/1f51a.png?v8",
            england: "unicode/1f3f4-e0067-e0062-e0065-e006e-e0067-e007f.png?v8",
            envelope: "unicode/2709.png?v8",
            envelope_with_arrow: "unicode/1f4e9.png?v8",
            equatorial_guinea: "unicode/1f1ec-1f1f6.png?v8",
            eritrea: "unicode/1f1ea-1f1f7.png?v8",
            es: "unicode/1f1ea-1f1f8.png?v8",
            estonia: "unicode/1f1ea-1f1ea.png?v8",
            ethiopia: "unicode/1f1ea-1f1f9.png?v8",
            eu: "unicode/1f1ea-1f1fa.png?v8",
            euro: "unicode/1f4b6.png?v8",
            european_castle: "unicode/1f3f0.png?v8",
            european_post_office: "unicode/1f3e4.png?v8",
            european_union: "unicode/1f1ea-1f1fa.png?v8",
            evergreen_tree: "unicode/1f332.png?v8",
            exclamation: "unicode/2757.png?v8",
            exploding_head: "unicode/1f92f.png?v8",
            expressionless: "unicode/1f611.png?v8",
            eye: "unicode/1f441.png?v8",
            eye_speech_bubble: "unicode/1f441-1f5e8.png?v8",
            eyeglasses: "unicode/1f453.png?v8",
            eyes: "unicode/1f440.png?v8",
            face_exhaling: "unicode/1f62e-1f4a8.png?v8",
            face_holding_back_tears: "unicode/1f979.png?v8",
            face_in_clouds: "unicode/1f636-1f32b.png?v8",
            face_with_diagonal_mouth: "unicode/1fae4.png?v8",
            face_with_head_bandage: "unicode/1f915.png?v8",
            face_with_open_eyes_and_hand_over_mouth: "unicode/1fae2.png?v8",
            face_with_peeking_eye: "unicode/1fae3.png?v8",
            face_with_spiral_eyes: "unicode/1f635-1f4ab.png?v8",
            face_with_thermometer: "unicode/1f912.png?v8",
            facepalm: "unicode/1f926.png?v8",
            facepunch: "unicode/1f44a.png?v8",
            factory: "unicode/1f3ed.png?v8",
            factory_worker: "unicode/1f9d1-1f3ed.png?v8",
            fairy: "unicode/1f9da.png?v8",
            fairy_man: "unicode/1f9da-2642.png?v8",
            fairy_woman: "unicode/1f9da-2640.png?v8",
            falafel: "unicode/1f9c6.png?v8",
            falkland_islands: "unicode/1f1eb-1f1f0.png?v8",
            fallen_leaf: "unicode/1f342.png?v8",
            family: "unicode/1f46a.png?v8",
            family_man_boy: "unicode/1f468-1f466.png?v8",
            family_man_boy_boy: "unicode/1f468-1f466-1f466.png?v8",
            family_man_girl: "unicode/1f468-1f467.png?v8",
            family_man_girl_boy: "unicode/1f468-1f467-1f466.png?v8",
            family_man_girl_girl: "unicode/1f468-1f467-1f467.png?v8",
            family_man_man_boy: "unicode/1f468-1f468-1f466.png?v8",
            family_man_man_boy_boy: "unicode/1f468-1f468-1f466-1f466.png?v8",
            family_man_man_girl: "unicode/1f468-1f468-1f467.png?v8",
            family_man_man_girl_boy: "unicode/1f468-1f468-1f467-1f466.png?v8",
            family_man_man_girl_girl: "unicode/1f468-1f468-1f467-1f467.png?v8",
            family_man_woman_boy: "unicode/1f468-1f469-1f466.png?v8",
            family_man_woman_boy_boy: "unicode/1f468-1f469-1f466-1f466.png?v8",
            family_man_woman_girl: "unicode/1f468-1f469-1f467.png?v8",
            family_man_woman_girl_boy: "unicode/1f468-1f469-1f467-1f466.png?v8",
            family_man_woman_girl_girl: "unicode/1f468-1f469-1f467-1f467.png?v8",
            family_woman_boy: "unicode/1f469-1f466.png?v8",
            family_woman_boy_boy: "unicode/1f469-1f466-1f466.png?v8",
            family_woman_girl: "unicode/1f469-1f467.png?v8",
            family_woman_girl_boy: "unicode/1f469-1f467-1f466.png?v8",
            family_woman_girl_girl: "unicode/1f469-1f467-1f467.png?v8",
            family_woman_woman_boy: "unicode/1f469-1f469-1f466.png?v8",
            family_woman_woman_boy_boy: "unicode/1f469-1f469-1f466-1f466.png?v8",
            family_woman_woman_girl: "unicode/1f469-1f469-1f467.png?v8",
            family_woman_woman_girl_boy: "unicode/1f469-1f469-1f467-1f466.png?v8",
            family_woman_woman_girl_girl: "unicode/1f469-1f469-1f467-1f467.png?v8",
            farmer: "unicode/1f9d1-1f33e.png?v8",
            faroe_islands: "unicode/1f1eb-1f1f4.png?v8",
            fast_forward: "unicode/23e9.png?v8",
            fax: "unicode/1f4e0.png?v8",
            fearful: "unicode/1f628.png?v8",
            feather: "unicode/1fab6.png?v8",
            feelsgood: "feelsgood.png?v8",
            feet: "unicode/1f43e.png?v8",
            female_detective: "unicode/1f575-2640.png?v8",
            female_sign: "unicode/2640.png?v8",
            ferris_wheel: "unicode/1f3a1.png?v8",
            ferry: "unicode/26f4.png?v8",
            field_hockey: "unicode/1f3d1.png?v8",
            fiji: "unicode/1f1eb-1f1ef.png?v8",
            file_cabinet: "unicode/1f5c4.png?v8",
            file_folder: "unicode/1f4c1.png?v8",
            film_projector: "unicode/1f4fd.png?v8",
            film_strip: "unicode/1f39e.png?v8",
            finland: "unicode/1f1eb-1f1ee.png?v8",
            finnadie: "finnadie.png?v8",
            fire: "unicode/1f525.png?v8",
            fire_engine: "unicode/1f692.png?v8",
            fire_extinguisher: "unicode/1f9ef.png?v8",
            firecracker: "unicode/1f9e8.png?v8",
            firefighter: "unicode/1f9d1-1f692.png?v8",
            fireworks: "unicode/1f386.png?v8",
            first_quarter_moon: "unicode/1f313.png?v8",
            first_quarter_moon_with_face: "unicode/1f31b.png?v8",
            fish: "unicode/1f41f.png?v8",
            fish_cake: "unicode/1f365.png?v8",
            fishing_pole_and_fish: "unicode/1f3a3.png?v8",
            fishsticks: "fishsticks.png?v8",
            fist: "unicode/270a.png?v8",
            fist_left: "unicode/1f91b.png?v8",
            fist_oncoming: "unicode/1f44a.png?v8",
            fist_raised: "unicode/270a.png?v8",
            fist_right: "unicode/1f91c.png?v8",
            five: "unicode/0035-20e3.png?v8",
            flags: "unicode/1f38f.png?v8",
            flamingo: "unicode/1f9a9.png?v8",
            flashlight: "unicode/1f526.png?v8",
            flat_shoe: "unicode/1f97f.png?v8",
            flatbread: "unicode/1fad3.png?v8",
            fleur_de_lis: "unicode/269c.png?v8",
            flight_arrival: "unicode/1f6ec.png?v8",
            flight_departure: "unicode/1f6eb.png?v8",
            flipper: "unicode/1f42c.png?v8",
            floppy_disk: "unicode/1f4be.png?v8",
            flower_playing_cards: "unicode/1f3b4.png?v8",
            flushed: "unicode/1f633.png?v8",
            flute: "unicode/1fa88.png?v8",
            fly: "unicode/1fab0.png?v8",
            flying_disc: "unicode/1f94f.png?v8",
            flying_saucer: "unicode/1f6f8.png?v8",
            fog: "unicode/1f32b.png?v8",
            foggy: "unicode/1f301.png?v8",
            folding_hand_fan: "unicode/1faad.png?v8",
            fondue: "unicode/1fad5.png?v8",
            foot: "unicode/1f9b6.png?v8",
            football: "unicode/1f3c8.png?v8",
            footprints: "unicode/1f463.png?v8",
            fork_and_knife: "unicode/1f374.png?v8",
            fortune_cookie: "unicode/1f960.png?v8",
            fountain: "unicode/26f2.png?v8",
            fountain_pen: "unicode/1f58b.png?v8",
            four: "unicode/0034-20e3.png?v8",
            four_leaf_clover: "unicode/1f340.png?v8",
            fox_face: "unicode/1f98a.png?v8",
            fr: "unicode/1f1eb-1f1f7.png?v8",
            framed_picture: "unicode/1f5bc.png?v8",
            free: "unicode/1f193.png?v8",
            french_guiana: "unicode/1f1ec-1f1eb.png?v8",
            french_polynesia: "unicode/1f1f5-1f1eb.png?v8",
            french_southern_territories: "unicode/1f1f9-1f1eb.png?v8",
            fried_egg: "unicode/1f373.png?v8",
            fried_shrimp: "unicode/1f364.png?v8",
            fries: "unicode/1f35f.png?v8",
            frog: "unicode/1f438.png?v8",
            frowning: "unicode/1f626.png?v8",
            frowning_face: "unicode/2639.png?v8",
            frowning_man: "unicode/1f64d-2642.png?v8",
            frowning_person: "unicode/1f64d.png?v8",
            frowning_woman: "unicode/1f64d-2640.png?v8",
            fu: "unicode/1f595.png?v8",
            fuelpump: "unicode/26fd.png?v8",
            full_moon: "unicode/1f315.png?v8",
            full_moon_with_face: "unicode/1f31d.png?v8",
            funeral_urn: "unicode/26b1.png?v8",
            gabon: "unicode/1f1ec-1f1e6.png?v8",
            gambia: "unicode/1f1ec-1f1f2.png?v8",
            game_die: "unicode/1f3b2.png?v8",
            garlic: "unicode/1f9c4.png?v8",
            gb: "unicode/1f1ec-1f1e7.png?v8",
            gear: "unicode/2699.png?v8",
            gem: "unicode/1f48e.png?v8",
            gemini: "unicode/264a.png?v8",
            genie: "unicode/1f9de.png?v8",
            genie_man: "unicode/1f9de-2642.png?v8",
            genie_woman: "unicode/1f9de-2640.png?v8",
            georgia: "unicode/1f1ec-1f1ea.png?v8",
            ghana: "unicode/1f1ec-1f1ed.png?v8",
            ghost: "unicode/1f47b.png?v8",
            gibraltar: "unicode/1f1ec-1f1ee.png?v8",
            gift: "unicode/1f381.png?v8",
            gift_heart: "unicode/1f49d.png?v8",
            ginger_root: "unicode/1fada.png?v8",
            giraffe: "unicode/1f992.png?v8",
            girl: "unicode/1f467.png?v8",
            globe_with_meridians: "unicode/1f310.png?v8",
            gloves: "unicode/1f9e4.png?v8",
            goal_net: "unicode/1f945.png?v8",
            goat: "unicode/1f410.png?v8",
            goberserk: "goberserk.png?v8",
            godmode: "godmode.png?v8",
            goggles: "unicode/1f97d.png?v8",
            golf: "unicode/26f3.png?v8",
            golfing: "unicode/1f3cc.png?v8",
            golfing_man: "unicode/1f3cc-2642.png?v8",
            golfing_woman: "unicode/1f3cc-2640.png?v8",
            goose: "unicode/1fabf.png?v8",
            gorilla: "unicode/1f98d.png?v8",
            grapes: "unicode/1f347.png?v8",
            greece: "unicode/1f1ec-1f1f7.png?v8",
            green_apple: "unicode/1f34f.png?v8",
            green_book: "unicode/1f4d7.png?v8",
            green_circle: "unicode/1f7e2.png?v8",
            green_heart: "unicode/1f49a.png?v8",
            green_salad: "unicode/1f957.png?v8",
            green_square: "unicode/1f7e9.png?v8",
            greenland: "unicode/1f1ec-1f1f1.png?v8",
            grenada: "unicode/1f1ec-1f1e9.png?v8",
            grey_exclamation: "unicode/2755.png?v8",
            grey_heart: "unicode/1fa76.png?v8",
            grey_question: "unicode/2754.png?v8",
            grimacing: "unicode/1f62c.png?v8",
            grin: "unicode/1f601.png?v8",
            grinning: "unicode/1f600.png?v8",
            guadeloupe: "unicode/1f1ec-1f1f5.png?v8",
            guam: "unicode/1f1ec-1f1fa.png?v8",
            guard: "unicode/1f482.png?v8",
            guardsman: "unicode/1f482-2642.png?v8",
            guardswoman: "unicode/1f482-2640.png?v8",
            guatemala: "unicode/1f1ec-1f1f9.png?v8",
            guernsey: "unicode/1f1ec-1f1ec.png?v8",
            guide_dog: "unicode/1f9ae.png?v8",
            guinea: "unicode/1f1ec-1f1f3.png?v8",
            guinea_bissau: "unicode/1f1ec-1f1fc.png?v8",
            guitar: "unicode/1f3b8.png?v8",
            gun: "unicode/1f52b.png?v8",
            guyana: "unicode/1f1ec-1f1fe.png?v8",
            hair_pick: "unicode/1faae.png?v8",
            haircut: "unicode/1f487.png?v8",
            haircut_man: "unicode/1f487-2642.png?v8",
            haircut_woman: "unicode/1f487-2640.png?v8",
            haiti: "unicode/1f1ed-1f1f9.png?v8",
            hamburger: "unicode/1f354.png?v8",
            hammer: "unicode/1f528.png?v8",
            hammer_and_pick: "unicode/2692.png?v8",
            hammer_and_wrench: "unicode/1f6e0.png?v8",
            hamsa: "unicode/1faac.png?v8",
            hamster: "unicode/1f439.png?v8",
            hand: "unicode/270b.png?v8",
            hand_over_mouth: "unicode/1f92d.png?v8",
            hand_with_index_finger_and_thumb_crossed: "unicode/1faf0.png?v8",
            handbag: "unicode/1f45c.png?v8",
            handball_person: "unicode/1f93e.png?v8",
            handshake: "unicode/1f91d.png?v8",
            hankey: "unicode/1f4a9.png?v8",
            hash: "unicode/0023-20e3.png?v8",
            hatched_chick: "unicode/1f425.png?v8",
            hatching_chick: "unicode/1f423.png?v8",
            headphones: "unicode/1f3a7.png?v8",
            headstone: "unicode/1faa6.png?v8",
            health_worker: "unicode/1f9d1-2695.png?v8",
            hear_no_evil: "unicode/1f649.png?v8",
            heard_mcdonald_islands: "unicode/1f1ed-1f1f2.png?v8",
            heart: "unicode/2764.png?v8",
            heart_decoration: "unicode/1f49f.png?v8",
            heart_eyes: "unicode/1f60d.png?v8",
            heart_eyes_cat: "unicode/1f63b.png?v8",
            heart_hands: "unicode/1faf6.png?v8",
            heart_on_fire: "unicode/2764-1f525.png?v8",
            heartbeat: "unicode/1f493.png?v8",
            heartpulse: "unicode/1f497.png?v8",
            hearts: "unicode/2665.png?v8",
            heavy_check_mark: "unicode/2714.png?v8",
            heavy_division_sign: "unicode/2797.png?v8",
            heavy_dollar_sign: "unicode/1f4b2.png?v8",
            heavy_equals_sign: "unicode/1f7f0.png?v8",
            heavy_exclamation_mark: "unicode/2757.png?v8",
            heavy_heart_exclamation: "unicode/2763.png?v8",
            heavy_minus_sign: "unicode/2796.png?v8",
            heavy_multiplication_x: "unicode/2716.png?v8",
            heavy_plus_sign: "unicode/2795.png?v8",
            hedgehog: "unicode/1f994.png?v8",
            helicopter: "unicode/1f681.png?v8",
            herb: "unicode/1f33f.png?v8",
            hibiscus: "unicode/1f33a.png?v8",
            high_brightness: "unicode/1f506.png?v8",
            high_heel: "unicode/1f460.png?v8",
            hiking_boot: "unicode/1f97e.png?v8",
            hindu_temple: "unicode/1f6d5.png?v8",
            hippopotamus: "unicode/1f99b.png?v8",
            hocho: "unicode/1f52a.png?v8",
            hole: "unicode/1f573.png?v8",
            honduras: "unicode/1f1ed-1f1f3.png?v8",
            honey_pot: "unicode/1f36f.png?v8",
            honeybee: "unicode/1f41d.png?v8",
            hong_kong: "unicode/1f1ed-1f1f0.png?v8",
            hook: "unicode/1fa9d.png?v8",
            horse: "unicode/1f434.png?v8",
            horse_racing: "unicode/1f3c7.png?v8",
            hospital: "unicode/1f3e5.png?v8",
            hot_face: "unicode/1f975.png?v8",
            hot_pepper: "unicode/1f336.png?v8",
            hotdog: "unicode/1f32d.png?v8",
            hotel: "unicode/1f3e8.png?v8",
            hotsprings: "unicode/2668.png?v8",
            hourglass: "unicode/231b.png?v8",
            hourglass_flowing_sand: "unicode/23f3.png?v8",
            house: "unicode/1f3e0.png?v8",
            house_with_garden: "unicode/1f3e1.png?v8",
            houses: "unicode/1f3d8.png?v8",
            hugs: "unicode/1f917.png?v8",
            hungary: "unicode/1f1ed-1f1fa.png?v8",
            hurtrealbad: "hurtrealbad.png?v8",
            hushed: "unicode/1f62f.png?v8",
            hut: "unicode/1f6d6.png?v8",
            hyacinth: "unicode/1fabb.png?v8",
            ice_cream: "unicode/1f368.png?v8",
            ice_cube: "unicode/1f9ca.png?v8",
            ice_hockey: "unicode/1f3d2.png?v8",
            ice_skate: "unicode/26f8.png?v8",
            icecream: "unicode/1f366.png?v8",
            iceland: "unicode/1f1ee-1f1f8.png?v8",
            id: "unicode/1f194.png?v8",
            identification_card: "unicode/1faaa.png?v8",
            ideograph_advantage: "unicode/1f250.png?v8",
            imp: "unicode/1f47f.png?v8",
            inbox_tray: "unicode/1f4e5.png?v8",
            incoming_envelope: "unicode/1f4e8.png?v8",
            index_pointing_at_the_viewer: "unicode/1faf5.png?v8",
            india: "unicode/1f1ee-1f1f3.png?v8",
            indonesia: "unicode/1f1ee-1f1e9.png?v8",
            infinity: "unicode/267e.png?v8",
            information_desk_person: "unicode/1f481.png?v8",
            information_source: "unicode/2139.png?v8",
            innocent: "unicode/1f607.png?v8",
            interrobang: "unicode/2049.png?v8",
            iphone: "unicode/1f4f1.png?v8",
            iran: "unicode/1f1ee-1f1f7.png?v8",
            iraq: "unicode/1f1ee-1f1f6.png?v8",
            ireland: "unicode/1f1ee-1f1ea.png?v8",
            isle_of_man: "unicode/1f1ee-1f1f2.png?v8",
            israel: "unicode/1f1ee-1f1f1.png?v8",
            it: "unicode/1f1ee-1f1f9.png?v8",
            izakaya_lantern: "unicode/1f3ee.png?v8",
            jack_o_lantern: "unicode/1f383.png?v8",
            jamaica: "unicode/1f1ef-1f1f2.png?v8",
            japan: "unicode/1f5fe.png?v8",
            japanese_castle: "unicode/1f3ef.png?v8",
            japanese_goblin: "unicode/1f47a.png?v8",
            japanese_ogre: "unicode/1f479.png?v8",
            jar: "unicode/1fad9.png?v8",
            jeans: "unicode/1f456.png?v8",
            jellyfish: "unicode/1fabc.png?v8",
            jersey: "unicode/1f1ef-1f1ea.png?v8",
            jigsaw: "unicode/1f9e9.png?v8",
            jordan: "unicode/1f1ef-1f1f4.png?v8",
            joy: "unicode/1f602.png?v8",
            joy_cat: "unicode/1f639.png?v8",
            joystick: "unicode/1f579.png?v8",
            jp: "unicode/1f1ef-1f1f5.png?v8",
            judge: "unicode/1f9d1-2696.png?v8",
            juggling_person: "unicode/1f939.png?v8",
            kaaba: "unicode/1f54b.png?v8",
            kangaroo: "unicode/1f998.png?v8",
            kazakhstan: "unicode/1f1f0-1f1ff.png?v8",
            kenya: "unicode/1f1f0-1f1ea.png?v8",
            key: "unicode/1f511.png?v8",
            keyboard: "unicode/2328.png?v8",
            keycap_ten: "unicode/1f51f.png?v8",
            khanda: "unicode/1faaf.png?v8",
            kick_scooter: "unicode/1f6f4.png?v8",
            kimono: "unicode/1f458.png?v8",
            kiribati: "unicode/1f1f0-1f1ee.png?v8",
            kiss: "unicode/1f48b.png?v8",
            kissing: "unicode/1f617.png?v8",
            kissing_cat: "unicode/1f63d.png?v8",
            kissing_closed_eyes: "unicode/1f61a.png?v8",
            kissing_heart: "unicode/1f618.png?v8",
            kissing_smiling_eyes: "unicode/1f619.png?v8",
            kite: "unicode/1fa81.png?v8",
            kiwi_fruit: "unicode/1f95d.png?v8",
            kneeling_man: "unicode/1f9ce-2642.png?v8",
            kneeling_person: "unicode/1f9ce.png?v8",
            kneeling_woman: "unicode/1f9ce-2640.png?v8",
            knife: "unicode/1f52a.png?v8",
            knot: "unicode/1faa2.png?v8",
            koala: "unicode/1f428.png?v8",
            koko: "unicode/1f201.png?v8",
            kosovo: "unicode/1f1fd-1f1f0.png?v8",
            kr: "unicode/1f1f0-1f1f7.png?v8",
            kuwait: "unicode/1f1f0-1f1fc.png?v8",
            kyrgyzstan: "unicode/1f1f0-1f1ec.png?v8",
            lab_coat: "unicode/1f97c.png?v8",
            label: "unicode/1f3f7.png?v8",
            lacrosse: "unicode/1f94d.png?v8",
            ladder: "unicode/1fa9c.png?v8",
            lady_beetle: "unicode/1f41e.png?v8",
            lantern: "unicode/1f3ee.png?v8",
            laos: "unicode/1f1f1-1f1e6.png?v8",
            large_blue_circle: "unicode/1f535.png?v8",
            large_blue_diamond: "unicode/1f537.png?v8",
            large_orange_diamond: "unicode/1f536.png?v8",
            last_quarter_moon: "unicode/1f317.png?v8",
            last_quarter_moon_with_face: "unicode/1f31c.png?v8",
            latin_cross: "unicode/271d.png?v8",
            latvia: "unicode/1f1f1-1f1fb.png?v8",
            laughing: "unicode/1f606.png?v8",
            leafy_green: "unicode/1f96c.png?v8",
            leaves: "unicode/1f343.png?v8",
            lebanon: "unicode/1f1f1-1f1e7.png?v8",
            ledger: "unicode/1f4d2.png?v8",
            left_luggage: "unicode/1f6c5.png?v8",
            left_right_arrow: "unicode/2194.png?v8",
            left_speech_bubble: "unicode/1f5e8.png?v8",
            leftwards_arrow_with_hook: "unicode/21a9.png?v8",
            leftwards_hand: "unicode/1faf2.png?v8",
            leftwards_pushing_hand: "unicode/1faf7.png?v8",
            leg: "unicode/1f9b5.png?v8",
            lemon: "unicode/1f34b.png?v8",
            leo: "unicode/264c.png?v8",
            leopard: "unicode/1f406.png?v8",
            lesotho: "unicode/1f1f1-1f1f8.png?v8",
            level_slider: "unicode/1f39a.png?v8",
            liberia: "unicode/1f1f1-1f1f7.png?v8",
            libra: "unicode/264e.png?v8",
            libya: "unicode/1f1f1-1f1fe.png?v8",
            liechtenstein: "unicode/1f1f1-1f1ee.png?v8",
            light_blue_heart: "unicode/1fa75.png?v8",
            light_rail: "unicode/1f688.png?v8",
            link: "unicode/1f517.png?v8",
            lion: "unicode/1f981.png?v8",
            lips: "unicode/1f444.png?v8",
            lipstick: "unicode/1f484.png?v8",
            lithuania: "unicode/1f1f1-1f1f9.png?v8",
            lizard: "unicode/1f98e.png?v8",
            llama: "unicode/1f999.png?v8",
            lobster: "unicode/1f99e.png?v8",
            lock: "unicode/1f512.png?v8",
            lock_with_ink_pen: "unicode/1f50f.png?v8",
            lollipop: "unicode/1f36d.png?v8",
            long_drum: "unicode/1fa98.png?v8",
            loop: "unicode/27bf.png?v8",
            lotion_bottle: "unicode/1f9f4.png?v8",
            lotus: "unicode/1fab7.png?v8",
            lotus_position: "unicode/1f9d8.png?v8",
            lotus_position_man: "unicode/1f9d8-2642.png?v8",
            lotus_position_woman: "unicode/1f9d8-2640.png?v8",
            loud_sound: "unicode/1f50a.png?v8",
            loudspeaker: "unicode/1f4e2.png?v8",
            love_hotel: "unicode/1f3e9.png?v8",
            love_letter: "unicode/1f48c.png?v8",
            love_you_gesture: "unicode/1f91f.png?v8",
            low_battery: "unicode/1faab.png?v8",
            low_brightness: "unicode/1f505.png?v8",
            luggage: "unicode/1f9f3.png?v8",
            lungs: "unicode/1fac1.png?v8",
            luxembourg: "unicode/1f1f1-1f1fa.png?v8",
            lying_face: "unicode/1f925.png?v8",
            m: "unicode/24c2.png?v8",
            macau: "unicode/1f1f2-1f1f4.png?v8",
            macedonia: "unicode/1f1f2-1f1f0.png?v8",
            madagascar: "unicode/1f1f2-1f1ec.png?v8",
            mag: "unicode/1f50d.png?v8",
            mag_right: "unicode/1f50e.png?v8",
            mage: "unicode/1f9d9.png?v8",
            mage_man: "unicode/1f9d9-2642.png?v8",
            mage_woman: "unicode/1f9d9-2640.png?v8",
            magic_wand: "unicode/1fa84.png?v8",
            magnet: "unicode/1f9f2.png?v8",
            mahjong: "unicode/1f004.png?v8",
            mailbox: "unicode/1f4eb.png?v8",
            mailbox_closed: "unicode/1f4ea.png?v8",
            mailbox_with_mail: "unicode/1f4ec.png?v8",
            mailbox_with_no_mail: "unicode/1f4ed.png?v8",
            malawi: "unicode/1f1f2-1f1fc.png?v8",
            malaysia: "unicode/1f1f2-1f1fe.png?v8",
            maldives: "unicode/1f1f2-1f1fb.png?v8",
            male_detective: "unicode/1f575-2642.png?v8",
            male_sign: "unicode/2642.png?v8",
            mali: "unicode/1f1f2-1f1f1.png?v8",
            malta: "unicode/1f1f2-1f1f9.png?v8",
            mammoth: "unicode/1f9a3.png?v8",
            man: "unicode/1f468.png?v8",
            man_artist: "unicode/1f468-1f3a8.png?v8",
            man_astronaut: "unicode/1f468-1f680.png?v8",
            man_beard: "unicode/1f9d4-2642.png?v8",
            man_cartwheeling: "unicode/1f938-2642.png?v8",
            man_cook: "unicode/1f468-1f373.png?v8",
            man_dancing: "unicode/1f57a.png?v8",
            man_facepalming: "unicode/1f926-2642.png?v8",
            man_factory_worker: "unicode/1f468-1f3ed.png?v8",
            man_farmer: "unicode/1f468-1f33e.png?v8",
            man_feeding_baby: "unicode/1f468-1f37c.png?v8",
            man_firefighter: "unicode/1f468-1f692.png?v8",
            man_health_worker: "unicode/1f468-2695.png?v8",
            man_in_manual_wheelchair: "unicode/1f468-1f9bd.png?v8",
            man_in_motorized_wheelchair: "unicode/1f468-1f9bc.png?v8",
            man_in_tuxedo: "unicode/1f935-2642.png?v8",
            man_judge: "unicode/1f468-2696.png?v8",
            man_juggling: "unicode/1f939-2642.png?v8",
            man_mechanic: "unicode/1f468-1f527.png?v8",
            man_office_worker: "unicode/1f468-1f4bc.png?v8",
            man_pilot: "unicode/1f468-2708.png?v8",
            man_playing_handball: "unicode/1f93e-2642.png?v8",
            man_playing_water_polo: "unicode/1f93d-2642.png?v8",
            man_scientist: "unicode/1f468-1f52c.png?v8",
            man_shrugging: "unicode/1f937-2642.png?v8",
            man_singer: "unicode/1f468-1f3a4.png?v8",
            man_student: "unicode/1f468-1f393.png?v8",
            man_teacher: "unicode/1f468-1f3eb.png?v8",
            man_technologist: "unicode/1f468-1f4bb.png?v8",
            man_with_gua_pi_mao: "unicode/1f472.png?v8",
            man_with_probing_cane: "unicode/1f468-1f9af.png?v8",
            man_with_turban: "unicode/1f473-2642.png?v8",
            man_with_veil: "unicode/1f470-2642.png?v8",
            mandarin: "unicode/1f34a.png?v8",
            mango: "unicode/1f96d.png?v8",
            mans_shoe: "unicode/1f45e.png?v8",
            mantelpiece_clock: "unicode/1f570.png?v8",
            manual_wheelchair: "unicode/1f9bd.png?v8",
            maple_leaf: "unicode/1f341.png?v8",
            maracas: "unicode/1fa87.png?v8",
            marshall_islands: "unicode/1f1f2-1f1ed.png?v8",
            martial_arts_uniform: "unicode/1f94b.png?v8",
            martinique: "unicode/1f1f2-1f1f6.png?v8",
            mask: "unicode/1f637.png?v8",
            massage: "unicode/1f486.png?v8",
            massage_man: "unicode/1f486-2642.png?v8",
            massage_woman: "unicode/1f486-2640.png?v8",
            mate: "unicode/1f9c9.png?v8",
            mauritania: "unicode/1f1f2-1f1f7.png?v8",
            mauritius: "unicode/1f1f2-1f1fa.png?v8",
            mayotte: "unicode/1f1fe-1f1f9.png?v8",
            meat_on_bone: "unicode/1f356.png?v8",
            mechanic: "unicode/1f9d1-1f527.png?v8",
            mechanical_arm: "unicode/1f9be.png?v8",
            mechanical_leg: "unicode/1f9bf.png?v8",
            medal_military: "unicode/1f396.png?v8",
            medal_sports: "unicode/1f3c5.png?v8",
            medical_symbol: "unicode/2695.png?v8",
            mega: "unicode/1f4e3.png?v8",
            melon: "unicode/1f348.png?v8",
            melting_face: "unicode/1fae0.png?v8",
            memo: "unicode/1f4dd.png?v8",
            men_wrestling: "unicode/1f93c-2642.png?v8",
            mending_heart: "unicode/2764-1fa79.png?v8",
            menorah: "unicode/1f54e.png?v8",
            mens: "unicode/1f6b9.png?v8",
            mermaid: "unicode/1f9dc-2640.png?v8",
            merman: "unicode/1f9dc-2642.png?v8",
            merperson: "unicode/1f9dc.png?v8",
            metal: "unicode/1f918.png?v8",
            metro: "unicode/1f687.png?v8",
            mexico: "unicode/1f1f2-1f1fd.png?v8",
            microbe: "unicode/1f9a0.png?v8",
            micronesia: "unicode/1f1eb-1f1f2.png?v8",
            microphone: "unicode/1f3a4.png?v8",
            microscope: "unicode/1f52c.png?v8",
            middle_finger: "unicode/1f595.png?v8",
            military_helmet: "unicode/1fa96.png?v8",
            milk_glass: "unicode/1f95b.png?v8",
            milky_way: "unicode/1f30c.png?v8",
            minibus: "unicode/1f690.png?v8",
            minidisc: "unicode/1f4bd.png?v8",
            mirror: "unicode/1fa9e.png?v8",
            mirror_ball: "unicode/1faa9.png?v8",
            mobile_phone_off: "unicode/1f4f4.png?v8",
            moldova: "unicode/1f1f2-1f1e9.png?v8",
            monaco: "unicode/1f1f2-1f1e8.png?v8",
            money_mouth_face: "unicode/1f911.png?v8",
            money_with_wings: "unicode/1f4b8.png?v8",
            moneybag: "unicode/1f4b0.png?v8",
            mongolia: "unicode/1f1f2-1f1f3.png?v8",
            monkey: "unicode/1f412.png?v8",
            monkey_face: "unicode/1f435.png?v8",
            monocle_face: "unicode/1f9d0.png?v8",
            monorail: "unicode/1f69d.png?v8",
            montenegro: "unicode/1f1f2-1f1ea.png?v8",
            montserrat: "unicode/1f1f2-1f1f8.png?v8",
            moon: "unicode/1f314.png?v8",
            moon_cake: "unicode/1f96e.png?v8",
            moose: "unicode/1face.png?v8",
            morocco: "unicode/1f1f2-1f1e6.png?v8",
            mortar_board: "unicode/1f393.png?v8",
            mosque: "unicode/1f54c.png?v8",
            mosquito: "unicode/1f99f.png?v8",
            motor_boat: "unicode/1f6e5.png?v8",
            motor_scooter: "unicode/1f6f5.png?v8",
            motorcycle: "unicode/1f3cd.png?v8",
            motorized_wheelchair: "unicode/1f9bc.png?v8",
            motorway: "unicode/1f6e3.png?v8",
            mount_fuji: "unicode/1f5fb.png?v8",
            mountain: "unicode/26f0.png?v8",
            mountain_bicyclist: "unicode/1f6b5.png?v8",
            mountain_biking_man: "unicode/1f6b5-2642.png?v8",
            mountain_biking_woman: "unicode/1f6b5-2640.png?v8",
            mountain_cableway: "unicode/1f6a0.png?v8",
            mountain_railway: "unicode/1f69e.png?v8",
            mountain_snow: "unicode/1f3d4.png?v8",
            mouse: "unicode/1f42d.png?v8",
            mouse2: "unicode/1f401.png?v8",
            mouse_trap: "unicode/1faa4.png?v8",
            movie_camera: "unicode/1f3a5.png?v8",
            moyai: "unicode/1f5ff.png?v8",
            mozambique: "unicode/1f1f2-1f1ff.png?v8",
            mrs_claus: "unicode/1f936.png?v8",
            muscle: "unicode/1f4aa.png?v8",
            mushroom: "unicode/1f344.png?v8",
            musical_keyboard: "unicode/1f3b9.png?v8",
            musical_note: "unicode/1f3b5.png?v8",
            musical_score: "unicode/1f3bc.png?v8",
            mute: "unicode/1f507.png?v8",
            mx_claus: "unicode/1f9d1-1f384.png?v8",
            myanmar: "unicode/1f1f2-1f1f2.png?v8",
            nail_care: "unicode/1f485.png?v8",
            name_badge: "unicode/1f4db.png?v8",
            namibia: "unicode/1f1f3-1f1e6.png?v8",
            national_park: "unicode/1f3de.png?v8",
            nauru: "unicode/1f1f3-1f1f7.png?v8",
            nauseated_face: "unicode/1f922.png?v8",
            nazar_amulet: "unicode/1f9ff.png?v8",
            neckbeard: "neckbeard.png?v8",
            necktie: "unicode/1f454.png?v8",
            negative_squared_cross_mark: "unicode/274e.png?v8",
            nepal: "unicode/1f1f3-1f1f5.png?v8",
            nerd_face: "unicode/1f913.png?v8",
            nest_with_eggs: "unicode/1faba.png?v8",
            nesting_dolls: "unicode/1fa86.png?v8",
            netherlands: "unicode/1f1f3-1f1f1.png?v8",
            neutral_face: "unicode/1f610.png?v8",
            new: "unicode/1f195.png?v8",
            new_caledonia: "unicode/1f1f3-1f1e8.png?v8",
            new_moon: "unicode/1f311.png?v8",
            new_moon_with_face: "unicode/1f31a.png?v8",
            new_zealand: "unicode/1f1f3-1f1ff.png?v8",
            newspaper: "unicode/1f4f0.png?v8",
            newspaper_roll: "unicode/1f5de.png?v8",
            next_track_button: "unicode/23ed.png?v8",
            ng: "unicode/1f196.png?v8",
            ng_man: "unicode/1f645-2642.png?v8",
            ng_woman: "unicode/1f645-2640.png?v8",
            nicaragua: "unicode/1f1f3-1f1ee.png?v8",
            niger: "unicode/1f1f3-1f1ea.png?v8",
            nigeria: "unicode/1f1f3-1f1ec.png?v8",
            night_with_stars: "unicode/1f303.png?v8",
            nine: "unicode/0039-20e3.png?v8",
            ninja: "unicode/1f977.png?v8",
            niue: "unicode/1f1f3-1f1fa.png?v8",
            no_bell: "unicode/1f515.png?v8",
            no_bicycles: "unicode/1f6b3.png?v8",
            no_entry: "unicode/26d4.png?v8",
            no_entry_sign: "unicode/1f6ab.png?v8",
            no_good: "unicode/1f645.png?v8",
            no_good_man: "unicode/1f645-2642.png?v8",
            no_good_woman: "unicode/1f645-2640.png?v8",
            no_mobile_phones: "unicode/1f4f5.png?v8",
            no_mouth: "unicode/1f636.png?v8",
            no_pedestrians: "unicode/1f6b7.png?v8",
            no_smoking: "unicode/1f6ad.png?v8",
            "non-potable_water": "unicode/1f6b1.png?v8",
            norfolk_island: "unicode/1f1f3-1f1eb.png?v8",
            north_korea: "unicode/1f1f0-1f1f5.png?v8",
            northern_mariana_islands: "unicode/1f1f2-1f1f5.png?v8",
            norway: "unicode/1f1f3-1f1f4.png?v8",
            nose: "unicode/1f443.png?v8",
            notebook: "unicode/1f4d3.png?v8",
            notebook_with_decorative_cover: "unicode/1f4d4.png?v8",
            notes: "unicode/1f3b6.png?v8",
            nut_and_bolt: "unicode/1f529.png?v8",
            o: "unicode/2b55.png?v8",
            o2: "unicode/1f17e.png?v8",
            ocean: "unicode/1f30a.png?v8",
            octocat: "octocat.png?v8",
            octopus: "unicode/1f419.png?v8",
            oden: "unicode/1f362.png?v8",
            office: "unicode/1f3e2.png?v8",
            office_worker: "unicode/1f9d1-1f4bc.png?v8",
            oil_drum: "unicode/1f6e2.png?v8",
            ok: "unicode/1f197.png?v8",
            ok_hand: "unicode/1f44c.png?v8",
            ok_man: "unicode/1f646-2642.png?v8",
            ok_person: "unicode/1f646.png?v8",
            ok_woman: "unicode/1f646-2640.png?v8",
            old_key: "unicode/1f5dd.png?v8",
            older_adult: "unicode/1f9d3.png?v8",
            older_man: "unicode/1f474.png?v8",
            older_woman: "unicode/1f475.png?v8",
            olive: "unicode/1fad2.png?v8",
            om: "unicode/1f549.png?v8",
            oman: "unicode/1f1f4-1f1f2.png?v8",
            on: "unicode/1f51b.png?v8",
            oncoming_automobile: "unicode/1f698.png?v8",
            oncoming_bus: "unicode/1f68d.png?v8",
            oncoming_police_car: "unicode/1f694.png?v8",
            oncoming_taxi: "unicode/1f696.png?v8",
            one: "unicode/0031-20e3.png?v8",
            one_piece_swimsuit: "unicode/1fa71.png?v8",
            onion: "unicode/1f9c5.png?v8",
            open_book: "unicode/1f4d6.png?v8",
            open_file_folder: "unicode/1f4c2.png?v8",
            open_hands: "unicode/1f450.png?v8",
            open_mouth: "unicode/1f62e.png?v8",
            open_umbrella: "unicode/2602.png?v8",
            ophiuchus: "unicode/26ce.png?v8",
            orange: "unicode/1f34a.png?v8",
            orange_book: "unicode/1f4d9.png?v8",
            orange_circle: "unicode/1f7e0.png?v8",
            orange_heart: "unicode/1f9e1.png?v8",
            orange_square: "unicode/1f7e7.png?v8",
            orangutan: "unicode/1f9a7.png?v8",
            orthodox_cross: "unicode/2626.png?v8",
            otter: "unicode/1f9a6.png?v8",
            outbox_tray: "unicode/1f4e4.png?v8",
            owl: "unicode/1f989.png?v8",
            ox: "unicode/1f402.png?v8",
            oyster: "unicode/1f9aa.png?v8",
            package: "unicode/1f4e6.png?v8",
            page_facing_up: "unicode/1f4c4.png?v8",
            page_with_curl: "unicode/1f4c3.png?v8",
            pager: "unicode/1f4df.png?v8",
            paintbrush: "unicode/1f58c.png?v8",
            pakistan: "unicode/1f1f5-1f1f0.png?v8",
            palau: "unicode/1f1f5-1f1fc.png?v8",
            palestinian_territories: "unicode/1f1f5-1f1f8.png?v8",
            palm_down_hand: "unicode/1faf3.png?v8",
            palm_tree: "unicode/1f334.png?v8",
            palm_up_hand: "unicode/1faf4.png?v8",
            palms_up_together: "unicode/1f932.png?v8",
            panama: "unicode/1f1f5-1f1e6.png?v8",
            pancakes: "unicode/1f95e.png?v8",
            panda_face: "unicode/1f43c.png?v8",
            paperclip: "unicode/1f4ce.png?v8",
            paperclips: "unicode/1f587.png?v8",
            papua_new_guinea: "unicode/1f1f5-1f1ec.png?v8",
            parachute: "unicode/1fa82.png?v8",
            paraguay: "unicode/1f1f5-1f1fe.png?v8",
            parasol_on_ground: "unicode/26f1.png?v8",
            parking: "unicode/1f17f.png?v8",
            parrot: "unicode/1f99c.png?v8",
            part_alternation_mark: "unicode/303d.png?v8",
            partly_sunny: "unicode/26c5.png?v8",
            partying_face: "unicode/1f973.png?v8",
            passenger_ship: "unicode/1f6f3.png?v8",
            passport_control: "unicode/1f6c2.png?v8",
            pause_button: "unicode/23f8.png?v8",
            paw_prints: "unicode/1f43e.png?v8",
            pea_pod: "unicode/1fadb.png?v8",
            peace_symbol: "unicode/262e.png?v8",
            peach: "unicode/1f351.png?v8",
            peacock: "unicode/1f99a.png?v8",
            peanuts: "unicode/1f95c.png?v8",
            pear: "unicode/1f350.png?v8",
            pen: "unicode/1f58a.png?v8",
            pencil: "unicode/1f4dd.png?v8",
            pencil2: "unicode/270f.png?v8",
            penguin: "unicode/1f427.png?v8",
            pensive: "unicode/1f614.png?v8",
            people_holding_hands: "unicode/1f9d1-1f91d-1f9d1.png?v8",
            people_hugging: "unicode/1fac2.png?v8",
            performing_arts: "unicode/1f3ad.png?v8",
            persevere: "unicode/1f623.png?v8",
            person_bald: "unicode/1f9d1-1f9b2.png?v8",
            person_curly_hair: "unicode/1f9d1-1f9b1.png?v8",
            person_feeding_baby: "unicode/1f9d1-1f37c.png?v8",
            person_fencing: "unicode/1f93a.png?v8",
            person_in_manual_wheelchair: "unicode/1f9d1-1f9bd.png?v8",
            person_in_motorized_wheelchair: "unicode/1f9d1-1f9bc.png?v8",
            person_in_tuxedo: "unicode/1f935.png?v8",
            person_red_hair: "unicode/1f9d1-1f9b0.png?v8",
            person_white_hair: "unicode/1f9d1-1f9b3.png?v8",
            person_with_crown: "unicode/1fac5.png?v8",
            person_with_probing_cane: "unicode/1f9d1-1f9af.png?v8",
            person_with_turban: "unicode/1f473.png?v8",
            person_with_veil: "unicode/1f470.png?v8",
            peru: "unicode/1f1f5-1f1ea.png?v8",
            petri_dish: "unicode/1f9eb.png?v8",
            philippines: "unicode/1f1f5-1f1ed.png?v8",
            phone: "unicode/260e.png?v8",
            pick: "unicode/26cf.png?v8",
            pickup_truck: "unicode/1f6fb.png?v8",
            pie: "unicode/1f967.png?v8",
            pig: "unicode/1f437.png?v8",
            pig2: "unicode/1f416.png?v8",
            pig_nose: "unicode/1f43d.png?v8",
            pill: "unicode/1f48a.png?v8",
            pilot: "unicode/1f9d1-2708.png?v8",
            pinata: "unicode/1fa85.png?v8",
            pinched_fingers: "unicode/1f90c.png?v8",
            pinching_hand: "unicode/1f90f.png?v8",
            pineapple: "unicode/1f34d.png?v8",
            ping_pong: "unicode/1f3d3.png?v8",
            pink_heart: "unicode/1fa77.png?v8",
            pirate_flag: "unicode/1f3f4-2620.png?v8",
            pisces: "unicode/2653.png?v8",
            pitcairn_islands: "unicode/1f1f5-1f1f3.png?v8",
            pizza: "unicode/1f355.png?v8",
            placard: "unicode/1faa7.png?v8",
            place_of_worship: "unicode/1f6d0.png?v8",
            plate_with_cutlery: "unicode/1f37d.png?v8",
            play_or_pause_button: "unicode/23ef.png?v8",
            playground_slide: "unicode/1f6dd.png?v8",
            pleading_face: "unicode/1f97a.png?v8",
            plunger: "unicode/1faa0.png?v8",
            point_down: "unicode/1f447.png?v8",
            point_left: "unicode/1f448.png?v8",
            point_right: "unicode/1f449.png?v8",
            point_up: "unicode/261d.png?v8",
            point_up_2: "unicode/1f446.png?v8",
            poland: "unicode/1f1f5-1f1f1.png?v8",
            polar_bear: "unicode/1f43b-2744.png?v8",
            police_car: "unicode/1f693.png?v8",
            police_officer: "unicode/1f46e.png?v8",
            policeman: "unicode/1f46e-2642.png?v8",
            policewoman: "unicode/1f46e-2640.png?v8",
            poodle: "unicode/1f429.png?v8",
            poop: "unicode/1f4a9.png?v8",
            popcorn: "unicode/1f37f.png?v8",
            portugal: "unicode/1f1f5-1f1f9.png?v8",
            post_office: "unicode/1f3e3.png?v8",
            postal_horn: "unicode/1f4ef.png?v8",
            postbox: "unicode/1f4ee.png?v8",
            potable_water: "unicode/1f6b0.png?v8",
            potato: "unicode/1f954.png?v8",
            potted_plant: "unicode/1fab4.png?v8",
            pouch: "unicode/1f45d.png?v8",
            poultry_leg: "unicode/1f357.png?v8",
            pound: "unicode/1f4b7.png?v8",
            pouring_liquid: "unicode/1fad7.png?v8",
            pout: "unicode/1f621.png?v8",
            pouting_cat: "unicode/1f63e.png?v8",
            pouting_face: "unicode/1f64e.png?v8",
            pouting_man: "unicode/1f64e-2642.png?v8",
            pouting_woman: "unicode/1f64e-2640.png?v8",
            pray: "unicode/1f64f.png?v8",
            prayer_beads: "unicode/1f4ff.png?v8",
            pregnant_man: "unicode/1fac3.png?v8",
            pregnant_person: "unicode/1fac4.png?v8",
            pregnant_woman: "unicode/1f930.png?v8",
            pretzel: "unicode/1f968.png?v8",
            previous_track_button: "unicode/23ee.png?v8",
            prince: "unicode/1f934.png?v8",
            princess: "unicode/1f478.png?v8",
            printer: "unicode/1f5a8.png?v8",
            probing_cane: "unicode/1f9af.png?v8",
            puerto_rico: "unicode/1f1f5-1f1f7.png?v8",
            punch: "unicode/1f44a.png?v8",
            purple_circle: "unicode/1f7e3.png?v8",
            purple_heart: "unicode/1f49c.png?v8",
            purple_square: "unicode/1f7ea.png?v8",
            purse: "unicode/1f45b.png?v8",
            pushpin: "unicode/1f4cc.png?v8",
            put_litter_in_its_place: "unicode/1f6ae.png?v8",
            qatar: "unicode/1f1f6-1f1e6.png?v8",
            question: "unicode/2753.png?v8",
            rabbit: "unicode/1f430.png?v8",
            rabbit2: "unicode/1f407.png?v8",
            raccoon: "unicode/1f99d.png?v8",
            racehorse: "unicode/1f40e.png?v8",
            racing_car: "unicode/1f3ce.png?v8",
            radio: "unicode/1f4fb.png?v8",
            radio_button: "unicode/1f518.png?v8",
            radioactive: "unicode/2622.png?v8",
            rage: "unicode/1f621.png?v8",
            rage1: "rage1.png?v8",
            rage2: "rage2.png?v8",
            rage3: "rage3.png?v8",
            rage4: "rage4.png?v8",
            railway_car: "unicode/1f683.png?v8",
            railway_track: "unicode/1f6e4.png?v8",
            rainbow: "unicode/1f308.png?v8",
            rainbow_flag: "unicode/1f3f3-1f308.png?v8",
            raised_back_of_hand: "unicode/1f91a.png?v8",
            raised_eyebrow: "unicode/1f928.png?v8",
            raised_hand: "unicode/270b.png?v8",
            raised_hand_with_fingers_splayed: "unicode/1f590.png?v8",
            raised_hands: "unicode/1f64c.png?v8",
            raising_hand: "unicode/1f64b.png?v8",
            raising_hand_man: "unicode/1f64b-2642.png?v8",
            raising_hand_woman: "unicode/1f64b-2640.png?v8",
            ram: "unicode/1f40f.png?v8",
            ramen: "unicode/1f35c.png?v8",
            rat: "unicode/1f400.png?v8",
            razor: "unicode/1fa92.png?v8",
            receipt: "unicode/1f9fe.png?v8",
            record_button: "unicode/23fa.png?v8",
            recycle: "unicode/267b.png?v8",
            red_car: "unicode/1f697.png?v8",
            red_circle: "unicode/1f534.png?v8",
            red_envelope: "unicode/1f9e7.png?v8",
            red_haired_man: "unicode/1f468-1f9b0.png?v8",
            red_haired_woman: "unicode/1f469-1f9b0.png?v8",
            red_square: "unicode/1f7e5.png?v8",
            registered: "unicode/00ae.png?v8",
            relaxed: "unicode/263a.png?v8",
            relieved: "unicode/1f60c.png?v8",
            reminder_ribbon: "unicode/1f397.png?v8",
            repeat: "unicode/1f501.png?v8",
            repeat_one: "unicode/1f502.png?v8",
            rescue_worker_helmet: "unicode/26d1.png?v8",
            restroom: "unicode/1f6bb.png?v8",
            reunion: "unicode/1f1f7-1f1ea.png?v8",
            revolving_hearts: "unicode/1f49e.png?v8",
            rewind: "unicode/23ea.png?v8",
            rhinoceros: "unicode/1f98f.png?v8",
            ribbon: "unicode/1f380.png?v8",
            rice: "unicode/1f35a.png?v8",
            rice_ball: "unicode/1f359.png?v8",
            rice_cracker: "unicode/1f358.png?v8",
            rice_scene: "unicode/1f391.png?v8",
            right_anger_bubble: "unicode/1f5ef.png?v8",
            rightwards_hand: "unicode/1faf1.png?v8",
            rightwards_pushing_hand: "unicode/1faf8.png?v8",
            ring: "unicode/1f48d.png?v8",
            ring_buoy: "unicode/1f6df.png?v8",
            ringed_planet: "unicode/1fa90.png?v8",
            robot: "unicode/1f916.png?v8",
            rock: "unicode/1faa8.png?v8",
            rocket: "unicode/1f680.png?v8",
            rofl: "unicode/1f923.png?v8",
            roll_eyes: "unicode/1f644.png?v8",
            roll_of_paper: "unicode/1f9fb.png?v8",
            roller_coaster: "unicode/1f3a2.png?v8",
            roller_skate: "unicode/1f6fc.png?v8",
            romania: "unicode/1f1f7-1f1f4.png?v8",
            rooster: "unicode/1f413.png?v8",
            rose: "unicode/1f339.png?v8",
            rosette: "unicode/1f3f5.png?v8",
            rotating_light: "unicode/1f6a8.png?v8",
            round_pushpin: "unicode/1f4cd.png?v8",
            rowboat: "unicode/1f6a3.png?v8",
            rowing_man: "unicode/1f6a3-2642.png?v8",
            rowing_woman: "unicode/1f6a3-2640.png?v8",
            ru: "unicode/1f1f7-1f1fa.png?v8",
            rugby_football: "unicode/1f3c9.png?v8",
            runner: "unicode/1f3c3.png?v8",
            running: "unicode/1f3c3.png?v8",
            running_man: "unicode/1f3c3-2642.png?v8",
            running_shirt_with_sash: "unicode/1f3bd.png?v8",
            running_woman: "unicode/1f3c3-2640.png?v8",
            rwanda: "unicode/1f1f7-1f1fc.png?v8",
            sa: "unicode/1f202.png?v8",
            safety_pin: "unicode/1f9f7.png?v8",
            safety_vest: "unicode/1f9ba.png?v8",
            sagittarius: "unicode/2650.png?v8",
            sailboat: "unicode/26f5.png?v8",
            sake: "unicode/1f376.png?v8",
            salt: "unicode/1f9c2.png?v8",
            saluting_face: "unicode/1fae1.png?v8",
            samoa: "unicode/1f1fc-1f1f8.png?v8",
            san_marino: "unicode/1f1f8-1f1f2.png?v8",
            sandal: "unicode/1f461.png?v8",
            sandwich: "unicode/1f96a.png?v8",
            santa: "unicode/1f385.png?v8",
            sao_tome_principe: "unicode/1f1f8-1f1f9.png?v8",
            sari: "unicode/1f97b.png?v8",
            sassy_man: "unicode/1f481-2642.png?v8",
            sassy_woman: "unicode/1f481-2640.png?v8",
            satellite: "unicode/1f4e1.png?v8",
            satisfied: "unicode/1f606.png?v8",
            saudi_arabia: "unicode/1f1f8-1f1e6.png?v8",
            sauna_man: "unicode/1f9d6-2642.png?v8",
            sauna_person: "unicode/1f9d6.png?v8",
            sauna_woman: "unicode/1f9d6-2640.png?v8",
            sauropod: "unicode/1f995.png?v8",
            saxophone: "unicode/1f3b7.png?v8",
            scarf: "unicode/1f9e3.png?v8",
            school: "unicode/1f3eb.png?v8",
            school_satchel: "unicode/1f392.png?v8",
            scientist: "unicode/1f9d1-1f52c.png?v8",
            scissors: "unicode/2702.png?v8",
            scorpion: "unicode/1f982.png?v8",
            scorpius: "unicode/264f.png?v8",
            scotland: "unicode/1f3f4-e0067-e0062-e0073-e0063-e0074-e007f.png?v8",
            scream: "unicode/1f631.png?v8",
            scream_cat: "unicode/1f640.png?v8",
            screwdriver: "unicode/1fa9b.png?v8",
            scroll: "unicode/1f4dc.png?v8",
            seal: "unicode/1f9ad.png?v8",
            seat: "unicode/1f4ba.png?v8",
            secret: "unicode/3299.png?v8",
            see_no_evil: "unicode/1f648.png?v8",
            seedling: "unicode/1f331.png?v8",
            selfie: "unicode/1f933.png?v8",
            senegal: "unicode/1f1f8-1f1f3.png?v8",
            serbia: "unicode/1f1f7-1f1f8.png?v8",
            service_dog: "unicode/1f415-1f9ba.png?v8",
            seven: "unicode/0037-20e3.png?v8",
            sewing_needle: "unicode/1faa1.png?v8",
            seychelles: "unicode/1f1f8-1f1e8.png?v8",
            shaking_face: "unicode/1fae8.png?v8",
            shallow_pan_of_food: "unicode/1f958.png?v8",
            shamrock: "unicode/2618.png?v8",
            shark: "unicode/1f988.png?v8",
            shaved_ice: "unicode/1f367.png?v8",
            sheep: "unicode/1f411.png?v8",
            shell: "unicode/1f41a.png?v8",
            shield: "unicode/1f6e1.png?v8",
            shinto_shrine: "unicode/26e9.png?v8",
            ship: "unicode/1f6a2.png?v8",
            shipit: "shipit.png?v8",
            shirt: "unicode/1f455.png?v8",
            shit: "unicode/1f4a9.png?v8",
            shoe: "unicode/1f45e.png?v8",
            shopping: "unicode/1f6cd.png?v8",
            shopping_cart: "unicode/1f6d2.png?v8",
            shorts: "unicode/1fa73.png?v8",
            shower: "unicode/1f6bf.png?v8",
            shrimp: "unicode/1f990.png?v8",
            shrug: "unicode/1f937.png?v8",
            shushing_face: "unicode/1f92b.png?v8",
            sierra_leone: "unicode/1f1f8-1f1f1.png?v8",
            signal_strength: "unicode/1f4f6.png?v8",
            singapore: "unicode/1f1f8-1f1ec.png?v8",
            singer: "unicode/1f9d1-1f3a4.png?v8",
            sint_maarten: "unicode/1f1f8-1f1fd.png?v8",
            six: "unicode/0036-20e3.png?v8",
            six_pointed_star: "unicode/1f52f.png?v8",
            skateboard: "unicode/1f6f9.png?v8",
            ski: "unicode/1f3bf.png?v8",
            skier: "unicode/26f7.png?v8",
            skull: "unicode/1f480.png?v8",
            skull_and_crossbones: "unicode/2620.png?v8",
            skunk: "unicode/1f9a8.png?v8",
            sled: "unicode/1f6f7.png?v8",
            sleeping: "unicode/1f634.png?v8",
            sleeping_bed: "unicode/1f6cc.png?v8",
            sleepy: "unicode/1f62a.png?v8",
            slightly_frowning_face: "unicode/1f641.png?v8",
            slightly_smiling_face: "unicode/1f642.png?v8",
            slot_machine: "unicode/1f3b0.png?v8",
            sloth: "unicode/1f9a5.png?v8",
            slovakia: "unicode/1f1f8-1f1f0.png?v8",
            slovenia: "unicode/1f1f8-1f1ee.png?v8",
            small_airplane: "unicode/1f6e9.png?v8",
            small_blue_diamond: "unicode/1f539.png?v8",
            small_orange_diamond: "unicode/1f538.png?v8",
            small_red_triangle: "unicode/1f53a.png?v8",
            small_red_triangle_down: "unicode/1f53b.png?v8",
            smile: "unicode/1f604.png?v8",
            smile_cat: "unicode/1f638.png?v8",
            smiley: "unicode/1f603.png?v8",
            smiley_cat: "unicode/1f63a.png?v8",
            smiling_face_with_tear: "unicode/1f972.png?v8",
            smiling_face_with_three_hearts: "unicode/1f970.png?v8",
            smiling_imp: "unicode/1f608.png?v8",
            smirk: "unicode/1f60f.png?v8",
            smirk_cat: "unicode/1f63c.png?v8",
            smoking: "unicode/1f6ac.png?v8",
            snail: "unicode/1f40c.png?v8",
            snake: "unicode/1f40d.png?v8",
            sneezing_face: "unicode/1f927.png?v8",
            snowboarder: "unicode/1f3c2.png?v8",
            snowflake: "unicode/2744.png?v8",
            snowman: "unicode/26c4.png?v8",
            snowman_with_snow: "unicode/2603.png?v8",
            soap: "unicode/1f9fc.png?v8",
            sob: "unicode/1f62d.png?v8",
            soccer: "unicode/26bd.png?v8",
            socks: "unicode/1f9e6.png?v8",
            softball: "unicode/1f94e.png?v8",
            solomon_islands: "unicode/1f1f8-1f1e7.png?v8",
            somalia: "unicode/1f1f8-1f1f4.png?v8",
            soon: "unicode/1f51c.png?v8",
            sos: "unicode/1f198.png?v8",
            sound: "unicode/1f509.png?v8",
            south_africa: "unicode/1f1ff-1f1e6.png?v8",
            south_georgia_south_sandwich_islands: "unicode/1f1ec-1f1f8.png?v8",
            south_sudan: "unicode/1f1f8-1f1f8.png?v8",
            space_invader: "unicode/1f47e.png?v8",
            spades: "unicode/2660.png?v8",
            spaghetti: "unicode/1f35d.png?v8",
            sparkle: "unicode/2747.png?v8",
            sparkler: "unicode/1f387.png?v8",
            sparkles: "unicode/2728.png?v8",
            sparkling_heart: "unicode/1f496.png?v8",
            speak_no_evil: "unicode/1f64a.png?v8",
            speaker: "unicode/1f508.png?v8",
            speaking_head: "unicode/1f5e3.png?v8",
            speech_balloon: "unicode/1f4ac.png?v8",
            speedboat: "unicode/1f6a4.png?v8",
            spider: "unicode/1f577.png?v8",
            spider_web: "unicode/1f578.png?v8",
            spiral_calendar: "unicode/1f5d3.png?v8",
            spiral_notepad: "unicode/1f5d2.png?v8",
            sponge: "unicode/1f9fd.png?v8",
            spoon: "unicode/1f944.png?v8",
            squid: "unicode/1f991.png?v8",
            sri_lanka: "unicode/1f1f1-1f1f0.png?v8",
            st_barthelemy: "unicode/1f1e7-1f1f1.png?v8",
            st_helena: "unicode/1f1f8-1f1ed.png?v8",
            st_kitts_nevis: "unicode/1f1f0-1f1f3.png?v8",
            st_lucia: "unicode/1f1f1-1f1e8.png?v8",
            st_martin: "unicode/1f1f2-1f1eb.png?v8",
            st_pierre_miquelon: "unicode/1f1f5-1f1f2.png?v8",
            st_vincent_grenadines: "unicode/1f1fb-1f1e8.png?v8",
            stadium: "unicode/1f3df.png?v8",
            standing_man: "unicode/1f9cd-2642.png?v8",
            standing_person: "unicode/1f9cd.png?v8",
            standing_woman: "unicode/1f9cd-2640.png?v8",
            star: "unicode/2b50.png?v8",
            star2: "unicode/1f31f.png?v8",
            star_and_crescent: "unicode/262a.png?v8",
            star_of_david: "unicode/2721.png?v8",
            star_struck: "unicode/1f929.png?v8",
            stars: "unicode/1f320.png?v8",
            station: "unicode/1f689.png?v8",
            statue_of_liberty: "unicode/1f5fd.png?v8",
            steam_locomotive: "unicode/1f682.png?v8",
            stethoscope: "unicode/1fa7a.png?v8",
            stew: "unicode/1f372.png?v8",
            stop_button: "unicode/23f9.png?v8",
            stop_sign: "unicode/1f6d1.png?v8",
            stopwatch: "unicode/23f1.png?v8",
            straight_ruler: "unicode/1f4cf.png?v8",
            strawberry: "unicode/1f353.png?v8",
            stuck_out_tongue: "unicode/1f61b.png?v8",
            stuck_out_tongue_closed_eyes: "unicode/1f61d.png?v8",
            stuck_out_tongue_winking_eye: "unicode/1f61c.png?v8",
            student: "unicode/1f9d1-1f393.png?v8",
            studio_microphone: "unicode/1f399.png?v8",
            stuffed_flatbread: "unicode/1f959.png?v8",
            sudan: "unicode/1f1f8-1f1e9.png?v8",
            sun_behind_large_cloud: "unicode/1f325.png?v8",
            sun_behind_rain_cloud: "unicode/1f326.png?v8",
            sun_behind_small_cloud: "unicode/1f324.png?v8",
            sun_with_face: "unicode/1f31e.png?v8",
            sunflower: "unicode/1f33b.png?v8",
            sunglasses: "unicode/1f60e.png?v8",
            sunny: "unicode/2600.png?v8",
            sunrise: "unicode/1f305.png?v8",
            sunrise_over_mountains: "unicode/1f304.png?v8",
            superhero: "unicode/1f9b8.png?v8",
            superhero_man: "unicode/1f9b8-2642.png?v8",
            superhero_woman: "unicode/1f9b8-2640.png?v8",
            supervillain: "unicode/1f9b9.png?v8",
            supervillain_man: "unicode/1f9b9-2642.png?v8",
            supervillain_woman: "unicode/1f9b9-2640.png?v8",
            surfer: "unicode/1f3c4.png?v8",
            surfing_man: "unicode/1f3c4-2642.png?v8",
            surfing_woman: "unicode/1f3c4-2640.png?v8",
            suriname: "unicode/1f1f8-1f1f7.png?v8",
            sushi: "unicode/1f363.png?v8",
            suspect: "suspect.png?v8",
            suspension_railway: "unicode/1f69f.png?v8",
            svalbard_jan_mayen: "unicode/1f1f8-1f1ef.png?v8",
            swan: "unicode/1f9a2.png?v8",
            swaziland: "unicode/1f1f8-1f1ff.png?v8",
            sweat: "unicode/1f613.png?v8",
            sweat_drops: "unicode/1f4a6.png?v8",
            sweat_smile: "unicode/1f605.png?v8",
            sweden: "unicode/1f1f8-1f1ea.png?v8",
            sweet_potato: "unicode/1f360.png?v8",
            swim_brief: "unicode/1fa72.png?v8",
            swimmer: "unicode/1f3ca.png?v8",
            swimming_man: "unicode/1f3ca-2642.png?v8",
            swimming_woman: "unicode/1f3ca-2640.png?v8",
            switzerland: "unicode/1f1e8-1f1ed.png?v8",
            symbols: "unicode/1f523.png?v8",
            synagogue: "unicode/1f54d.png?v8",
            syria: "unicode/1f1f8-1f1fe.png?v8",
            syringe: "unicode/1f489.png?v8",
            "t-rex": "unicode/1f996.png?v8",
            taco: "unicode/1f32e.png?v8",
            tada: "unicode/1f389.png?v8",
            taiwan: "unicode/1f1f9-1f1fc.png?v8",
            tajikistan: "unicode/1f1f9-1f1ef.png?v8",
            takeout_box: "unicode/1f961.png?v8",
            tamale: "unicode/1fad4.png?v8",
            tanabata_tree: "unicode/1f38b.png?v8",
            tangerine: "unicode/1f34a.png?v8",
            tanzania: "unicode/1f1f9-1f1ff.png?v8",
            taurus: "unicode/2649.png?v8",
            taxi: "unicode/1f695.png?v8",
            tea: "unicode/1f375.png?v8",
            teacher: "unicode/1f9d1-1f3eb.png?v8",
            teapot: "unicode/1fad6.png?v8",
            technologist: "unicode/1f9d1-1f4bb.png?v8",
            teddy_bear: "unicode/1f9f8.png?v8",
            telephone: "unicode/260e.png?v8",
            telephone_receiver: "unicode/1f4de.png?v8",
            telescope: "unicode/1f52d.png?v8",
            tennis: "unicode/1f3be.png?v8",
            tent: "unicode/26fa.png?v8",
            test_tube: "unicode/1f9ea.png?v8",
            thailand: "unicode/1f1f9-1f1ed.png?v8",
            thermometer: "unicode/1f321.png?v8",
            thinking: "unicode/1f914.png?v8",
            thong_sandal: "unicode/1fa74.png?v8",
            thought_balloon: "unicode/1f4ad.png?v8",
            thread: "unicode/1f9f5.png?v8",
            three: "unicode/0033-20e3.png?v8",
            thumbsdown: "unicode/1f44e.png?v8",
            thumbsup: "unicode/1f44d.png?v8",
            ticket: "unicode/1f3ab.png?v8",
            tickets: "unicode/1f39f.png?v8",
            tiger: "unicode/1f42f.png?v8",
            tiger2: "unicode/1f405.png?v8",
            timer_clock: "unicode/23f2.png?v8",
            timor_leste: "unicode/1f1f9-1f1f1.png?v8",
            tipping_hand_man: "unicode/1f481-2642.png?v8",
            tipping_hand_person: "unicode/1f481.png?v8",
            tipping_hand_woman: "unicode/1f481-2640.png?v8",
            tired_face: "unicode/1f62b.png?v8",
            tm: "unicode/2122.png?v8",
            togo: "unicode/1f1f9-1f1ec.png?v8",
            toilet: "unicode/1f6bd.png?v8",
            tokelau: "unicode/1f1f9-1f1f0.png?v8",
            tokyo_tower: "unicode/1f5fc.png?v8",
            tomato: "unicode/1f345.png?v8",
            tonga: "unicode/1f1f9-1f1f4.png?v8",
            tongue: "unicode/1f445.png?v8",
            toolbox: "unicode/1f9f0.png?v8",
            tooth: "unicode/1f9b7.png?v8",
            toothbrush: "unicode/1faa5.png?v8",
            top: "unicode/1f51d.png?v8",
            tophat: "unicode/1f3a9.png?v8",
            tornado: "unicode/1f32a.png?v8",
            tr: "unicode/1f1f9-1f1f7.png?v8",
            trackball: "unicode/1f5b2.png?v8",
            tractor: "unicode/1f69c.png?v8",
            traffic_light: "unicode/1f6a5.png?v8",
            train: "unicode/1f68b.png?v8",
            train2: "unicode/1f686.png?v8",
            tram: "unicode/1f68a.png?v8",
            transgender_flag: "unicode/1f3f3-26a7.png?v8",
            transgender_symbol: "unicode/26a7.png?v8",
            triangular_flag_on_post: "unicode/1f6a9.png?v8",
            triangular_ruler: "unicode/1f4d0.png?v8",
            trident: "unicode/1f531.png?v8",
            trinidad_tobago: "unicode/1f1f9-1f1f9.png?v8",
            tristan_da_cunha: "unicode/1f1f9-1f1e6.png?v8",
            triumph: "unicode/1f624.png?v8",
            troll: "unicode/1f9cc.png?v8",
            trolleybus: "unicode/1f68e.png?v8",
            trollface: "trollface.png?v8",
            trophy: "unicode/1f3c6.png?v8",
            tropical_drink: "unicode/1f379.png?v8",
            tropical_fish: "unicode/1f420.png?v8",
            truck: "unicode/1f69a.png?v8",
            trumpet: "unicode/1f3ba.png?v8",
            tshirt: "unicode/1f455.png?v8",
            tulip: "unicode/1f337.png?v8",
            tumbler_glass: "unicode/1f943.png?v8",
            tunisia: "unicode/1f1f9-1f1f3.png?v8",
            turkey: "unicode/1f983.png?v8",
            turkmenistan: "unicode/1f1f9-1f1f2.png?v8",
            turks_caicos_islands: "unicode/1f1f9-1f1e8.png?v8",
            turtle: "unicode/1f422.png?v8",
            tuvalu: "unicode/1f1f9-1f1fb.png?v8",
            tv: "unicode/1f4fa.png?v8",
            twisted_rightwards_arrows: "unicode/1f500.png?v8",
            two: "unicode/0032-20e3.png?v8",
            two_hearts: "unicode/1f495.png?v8",
            two_men_holding_hands: "unicode/1f46c.png?v8",
            two_women_holding_hands: "unicode/1f46d.png?v8",
            u5272: "unicode/1f239.png?v8",
            u5408: "unicode/1f234.png?v8",
            u55b6: "unicode/1f23a.png?v8",
            u6307: "unicode/1f22f.png?v8",
            u6708: "unicode/1f237.png?v8",
            u6709: "unicode/1f236.png?v8",
            u6e80: "unicode/1f235.png?v8",
            u7121: "unicode/1f21a.png?v8",
            u7533: "unicode/1f238.png?v8",
            u7981: "unicode/1f232.png?v8",
            u7a7a: "unicode/1f233.png?v8",
            uganda: "unicode/1f1fa-1f1ec.png?v8",
            uk: "unicode/1f1ec-1f1e7.png?v8",
            ukraine: "unicode/1f1fa-1f1e6.png?v8",
            umbrella: "unicode/2614.png?v8",
            unamused: "unicode/1f612.png?v8",
            underage: "unicode/1f51e.png?v8",
            unicorn: "unicode/1f984.png?v8",
            united_arab_emirates: "unicode/1f1e6-1f1ea.png?v8",
            united_nations: "unicode/1f1fa-1f1f3.png?v8",
            unlock: "unicode/1f513.png?v8",
            up: "unicode/1f199.png?v8",
            upside_down_face: "unicode/1f643.png?v8",
            uruguay: "unicode/1f1fa-1f1fe.png?v8",
            us: "unicode/1f1fa-1f1f8.png?v8",
            us_outlying_islands: "unicode/1f1fa-1f1f2.png?v8",
            us_virgin_islands: "unicode/1f1fb-1f1ee.png?v8",
            uzbekistan: "unicode/1f1fa-1f1ff.png?v8",
            v: "unicode/270c.png?v8",
            vampire: "unicode/1f9db.png?v8",
            vampire_man: "unicode/1f9db-2642.png?v8",
            vampire_woman: "unicode/1f9db-2640.png?v8",
            vanuatu: "unicode/1f1fb-1f1fa.png?v8",
            vatican_city: "unicode/1f1fb-1f1e6.png?v8",
            venezuela: "unicode/1f1fb-1f1ea.png?v8",
            vertical_traffic_light: "unicode/1f6a6.png?v8",
            vhs: "unicode/1f4fc.png?v8",
            vibration_mode: "unicode/1f4f3.png?v8",
            video_camera: "unicode/1f4f9.png?v8",
            video_game: "unicode/1f3ae.png?v8",
            vietnam: "unicode/1f1fb-1f1f3.png?v8",
            violin: "unicode/1f3bb.png?v8",
            virgo: "unicode/264d.png?v8",
            volcano: "unicode/1f30b.png?v8",
            volleyball: "unicode/1f3d0.png?v8",
            vomiting_face: "unicode/1f92e.png?v8",
            vs: "unicode/1f19a.png?v8",
            vulcan_salute: "unicode/1f596.png?v8",
            waffle: "unicode/1f9c7.png?v8",
            wales: "unicode/1f3f4-e0067-e0062-e0077-e006c-e0073-e007f.png?v8",
            walking: "unicode/1f6b6.png?v8",
            walking_man: "unicode/1f6b6-2642.png?v8",
            walking_woman: "unicode/1f6b6-2640.png?v8",
            wallis_futuna: "unicode/1f1fc-1f1eb.png?v8",
            waning_crescent_moon: "unicode/1f318.png?v8",
            waning_gibbous_moon: "unicode/1f316.png?v8",
            warning: "unicode/26a0.png?v8",
            wastebasket: "unicode/1f5d1.png?v8",
            watch: "unicode/231a.png?v8",
            water_buffalo: "unicode/1f403.png?v8",
            water_polo: "unicode/1f93d.png?v8",
            watermelon: "unicode/1f349.png?v8",
            wave: "unicode/1f44b.png?v8",
            wavy_dash: "unicode/3030.png?v8",
            waxing_crescent_moon: "unicode/1f312.png?v8",
            waxing_gibbous_moon: "unicode/1f314.png?v8",
            wc: "unicode/1f6be.png?v8",
            weary: "unicode/1f629.png?v8",
            wedding: "unicode/1f492.png?v8",
            weight_lifting: "unicode/1f3cb.png?v8",
            weight_lifting_man: "unicode/1f3cb-2642.png?v8",
            weight_lifting_woman: "unicode/1f3cb-2640.png?v8",
            western_sahara: "unicode/1f1ea-1f1ed.png?v8",
            whale: "unicode/1f433.png?v8",
            whale2: "unicode/1f40b.png?v8",
            wheel: "unicode/1f6de.png?v8",
            wheel_of_dharma: "unicode/2638.png?v8",
            wheelchair: "unicode/267f.png?v8",
            white_check_mark: "unicode/2705.png?v8",
            white_circle: "unicode/26aa.png?v8",
            white_flag: "unicode/1f3f3.png?v8",
            white_flower: "unicode/1f4ae.png?v8",
            white_haired_man: "unicode/1f468-1f9b3.png?v8",
            white_haired_woman: "unicode/1f469-1f9b3.png?v8",
            white_heart: "unicode/1f90d.png?v8",
            white_large_square: "unicode/2b1c.png?v8",
            white_medium_small_square: "unicode/25fd.png?v8",
            white_medium_square: "unicode/25fb.png?v8",
            white_small_square: "unicode/25ab.png?v8",
            white_square_button: "unicode/1f533.png?v8",
            wilted_flower: "unicode/1f940.png?v8",
            wind_chime: "unicode/1f390.png?v8",
            wind_face: "unicode/1f32c.png?v8",
            window: "unicode/1fa9f.png?v8",
            wine_glass: "unicode/1f377.png?v8",
            wing: "unicode/1fabd.png?v8",
            wink: "unicode/1f609.png?v8",
            wireless: "unicode/1f6dc.png?v8",
            wolf: "unicode/1f43a.png?v8",
            woman: "unicode/1f469.png?v8",
            woman_artist: "unicode/1f469-1f3a8.png?v8",
            woman_astronaut: "unicode/1f469-1f680.png?v8",
            woman_beard: "unicode/1f9d4-2640.png?v8",
            woman_cartwheeling: "unicode/1f938-2640.png?v8",
            woman_cook: "unicode/1f469-1f373.png?v8",
            woman_dancing: "unicode/1f483.png?v8",
            woman_facepalming: "unicode/1f926-2640.png?v8",
            woman_factory_worker: "unicode/1f469-1f3ed.png?v8",
            woman_farmer: "unicode/1f469-1f33e.png?v8",
            woman_feeding_baby: "unicode/1f469-1f37c.png?v8",
            woman_firefighter: "unicode/1f469-1f692.png?v8",
            woman_health_worker: "unicode/1f469-2695.png?v8",
            woman_in_manual_wheelchair: "unicode/1f469-1f9bd.png?v8",
            woman_in_motorized_wheelchair: "unicode/1f469-1f9bc.png?v8",
            woman_in_tuxedo: "unicode/1f935-2640.png?v8",
            woman_judge: "unicode/1f469-2696.png?v8",
            woman_juggling: "unicode/1f939-2640.png?v8",
            woman_mechanic: "unicode/1f469-1f527.png?v8",
            woman_office_worker: "unicode/1f469-1f4bc.png?v8",
            woman_pilot: "unicode/1f469-2708.png?v8",
            woman_playing_handball: "unicode/1f93e-2640.png?v8",
            woman_playing_water_polo: "unicode/1f93d-2640.png?v8",
            woman_scientist: "unicode/1f469-1f52c.png?v8",
            woman_shrugging: "unicode/1f937-2640.png?v8",
            woman_singer: "unicode/1f469-1f3a4.png?v8",
            woman_student: "unicode/1f469-1f393.png?v8",
            woman_teacher: "unicode/1f469-1f3eb.png?v8",
            woman_technologist: "unicode/1f469-1f4bb.png?v8",
            woman_with_headscarf: "unicode/1f9d5.png?v8",
            woman_with_probing_cane: "unicode/1f469-1f9af.png?v8",
            woman_with_turban: "unicode/1f473-2640.png?v8",
            woman_with_veil: "unicode/1f470-2640.png?v8",
            womans_clothes: "unicode/1f45a.png?v8",
            womans_hat: "unicode/1f452.png?v8",
            women_wrestling: "unicode/1f93c-2640.png?v8",
            womens: "unicode/1f6ba.png?v8",
            wood: "unicode/1fab5.png?v8",
            woozy_face: "unicode/1f974.png?v8",
            world_map: "unicode/1f5fa.png?v8",
            worm: "unicode/1fab1.png?v8",
            worried: "unicode/1f61f.png?v8",
            wrench: "unicode/1f527.png?v8",
            wrestling: "unicode/1f93c.png?v8",
            writing_hand: "unicode/270d.png?v8",
            x: "unicode/274c.png?v8",
            x_ray: "unicode/1fa7b.png?v8",
            yarn: "unicode/1f9f6.png?v8",
            yawning_face: "unicode/1f971.png?v8",
            yellow_circle: "unicode/1f7e1.png?v8",
            yellow_heart: "unicode/1f49b.png?v8",
            yellow_square: "unicode/1f7e8.png?v8",
            yemen: "unicode/1f1fe-1f1ea.png?v8",
            yen: "unicode/1f4b4.png?v8",
            yin_yang: "unicode/262f.png?v8",
            yo_yo: "unicode/1fa80.png?v8",
            yum: "unicode/1f60b.png?v8",
            zambia: "unicode/1f1ff-1f1f2.png?v8",
            zany_face: "unicode/1f92a.png?v8",
            zap: "unicode/26a1.png?v8",
            zebra: "unicode/1f993.png?v8",
            zero: "unicode/0030-20e3.png?v8",
            zimbabwe: "unicode/1f1ff-1f1fc.png?v8",
            zipper_mouth_face: "unicode/1f910.png?v8",
            zombie: "unicode/1f9df.png?v8",
            zombie_man: "unicode/1f9df-2642.png?v8",
            zombie_woman: "unicode/1f9df-2640.png?v8",
            zzz: "unicode/1f4a4.png?v8"
        }
    };

    function on(e, n) {
        return e.replace(/<(code|pre|script|template)[^>]*?>[\s\S]+?<\/(code|pre|script|template)>/g, (e => e.replace(/:/g, "__colon__"))).replace(/<!--[\s\S]+?-->/g, (e => e.replace(/:/g, "__colon__"))).replace(/([a-z]{2,}:)?\/\/[^\s'">)]+/gi, (e => e.replace(/:/g, "__colon__"))).replace(/:([a-z0-9_\-+]+?):/g, ((e, i) => function (e, n, i) {
            const t = tn.data[n];
            let o = e;
            t && (o = i && /unicode/.test(t) ? `<span class="emoji">${t.replace("unicode/", "").replace(/\.png.*/, "").split("-").map((e => `&#x${e};`)).join("&zwj;").concat("&#xFE0E;")}</span>` : `<img src="${tn.baseURL}${t}.png" alt="${n}" class="emoji" loading="lazy">`);
            return o
        }(e, i, n))).replace(/__colon__/g, ":")
    }

    function an() {
        let e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
        const n = {};
        return e && (e = e.replace(/^('|")/, "").replace(/('|")$/, "").replace(/(?:^|\s):([\w-]+:?)=?([\w-%]+)?/g, ((e, i, t) => -1 === i.indexOf(":") ? (n[i] = t && t.replace(/&quot;/g, "") || !0, "") : e)).trim()), {
            str: e,
            config: n
        }
    }

    var cn = "undefined" != typeof globalThis ? globalThis : "undefined" != typeof window ? window : "undefined" != typeof global ? global : "undefined" != typeof self ? self : {};

    function rn(e) {
        return e && e.__esModule && Object.prototype.hasOwnProperty.call(e, "default") ? e.default : e
    }

    var sn = {exports: {}};
    !function (e) {
        var n = function (e) {
            var n = /(?:^|\s)lang(?:uage)?-([\w-]+)(?=\s|$)/i, i = 0, t = {}, o = {
                manual: e.Prism && e.Prism.manual,
                disableWorkerMessageHandler: e.Prism && e.Prism.disableWorkerMessageHandler,
                util: {
                    encode: function e(n) {
                        return n instanceof a ? new a(n.type, e(n.content), n.alias) : Array.isArray(n) ? n.map(e) : n.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/\u00a0/g, " ")
                    }, type: function (e) {
                        return Object.prototype.toString.call(e).slice(8, -1)
                    }, objId: function (e) {
                        return e.__id || Object.defineProperty(e, "__id", {value: ++i}), e.__id
                    }, clone: function e(n, i) {
                        var t, a;
                        switch (i = i || {}, o.util.type(n)) {
                            case"Object":
                                if (a = o.util.objId(n), i[a]) return i[a];
                                for (var c in t = {}, i[a] = t, n) n.hasOwnProperty(c) && (t[c] = e(n[c], i));
                                return t;
                            case"Array":
                                return a = o.util.objId(n), i[a] ? i[a] : (t = [], i[a] = t, n.forEach((function (n, o) {
                                    t[o] = e(n, i)
                                })), t);
                            default:
                                return n
                        }
                    }, getLanguage: function (e) {
                        for (; e;) {
                            var i = n.exec(e.className);
                            if (i) return i[1].toLowerCase();
                            e = e.parentElement
                        }
                        return "none"
                    }, setLanguage: function (e, i) {
                        e.className = e.className.replace(RegExp(n, "gi"), ""), e.classList.add("language-" + i)
                    }, currentScript: function () {
                        if ("undefined" == typeof document) return null;
                        if ("currentScript" in document) return document.currentScript;
                        try {
                            throw new Error
                        } catch (t) {
                            var e = (/at [^(\r\n]*\((.*):[^:]+:[^:]+\)$/i.exec(t.stack) || [])[1];
                            if (e) {
                                var n = document.getElementsByTagName("script");
                                for (var i in n) if (n[i].src == e) return n[i]
                            }
                            return null
                        }
                    }, isActive: function (e, n, i) {
                        for (var t = "no-" + n; e;) {
                            var o = e.classList;
                            if (o.contains(n)) return !0;
                            if (o.contains(t)) return !1;
                            e = e.parentElement
                        }
                        return !!i
                    }
                },
                languages: {
                    plain: t, plaintext: t, text: t, txt: t, extend: function (e, n) {
                        var i = o.util.clone(o.languages[e]);
                        for (var t in n) i[t] = n[t];
                        return i
                    }, insertBefore: function (e, n, i, t) {
                        var a = (t = t || o.languages)[e], c = {};
                        for (var r in a) if (a.hasOwnProperty(r)) {
                            if (r == n) for (var s in i) i.hasOwnProperty(s) && (c[s] = i[s]);
                            i.hasOwnProperty(r) || (c[r] = a[r])
                        }
                        var u = t[e];
                        return t[e] = c, o.languages.DFS(o.languages, (function (n, i) {
                            i === u && n != e && (this[n] = c)
                        })), c
                    }, DFS: function e(n, i, t, a) {
                        a = a || {};
                        var c = o.util.objId;
                        for (var r in n) if (n.hasOwnProperty(r)) {
                            i.call(n, r, n[r], t || r);
                            var s = n[r], u = o.util.type(s);
                            "Object" !== u || a[c(s)] ? "Array" !== u || a[c(s)] || (a[c(s)] = !0, e(s, i, r, a)) : (a[c(s)] = !0, e(s, i, null, a))
                        }
                    }
                },
                plugins: {},
                highlightAll: function (e, n) {
                    o.highlightAllUnder(document, e, n)
                },
                highlightAllUnder: function (e, n, i) {
                    var t = {
                        callback: i,
                        container: e,
                        selector: 'code[class*="language-"], [class*="language-"] code, code[class*="lang-"], [class*="lang-"] code'
                    };
                    o.hooks.run("before-highlightall", t), t.elements = Array.prototype.slice.apply(t.container.querySelectorAll(t.selector)), o.hooks.run("before-all-elements-highlight", t);
                    for (var a, c = 0; a = t.elements[c++];) o.highlightElement(a, !0 === n, t.callback)
                },
                highlightElement: function (n, i, t) {
                    var a = o.util.getLanguage(n), c = o.languages[a];
                    o.util.setLanguage(n, a);
                    var r = n.parentElement;
                    r && "pre" === r.nodeName.toLowerCase() && o.util.setLanguage(r, a);
                    var s = {element: n, language: a, grammar: c, code: n.textContent};

                    function u(e) {
                        s.highlightedCode = e, o.hooks.run("before-insert", s), s.element.innerHTML = s.highlightedCode, o.hooks.run("after-highlight", s), o.hooks.run("complete", s), t && t.call(s.element)
                    }

                    if (o.hooks.run("before-sanity-check", s), (r = s.element.parentElement) && "pre" === r.nodeName.toLowerCase() && !r.hasAttribute("tabindex") && r.setAttribute("tabindex", "0"), !s.code) return o.hooks.run("complete", s), void (t && t.call(s.element));
                    if (o.hooks.run("before-highlight", s), s.grammar) if (i && e.Worker) {
                        var d = new Worker(o.filename);
                        d.onmessage = function (e) {
                            u(e.data)
                        }, d.postMessage(JSON.stringify({language: s.language, code: s.code, immediateClose: !0}))
                    } else u(o.highlight(s.code, s.grammar, s.language)); else u(o.util.encode(s.code))
                },
                highlight: function (e, n, i) {
                    var t = {code: e, grammar: n, language: i};
                    if (o.hooks.run("before-tokenize", t), !t.grammar) throw new Error('The language "' + t.language + '" has no grammar.');
                    return t.tokens = o.tokenize(t.code, t.grammar), o.hooks.run("after-tokenize", t), a.stringify(o.util.encode(t.tokens), t.language)
                },
                tokenize: function (e, n) {
                    var i = n.rest;
                    if (i) {
                        for (var t in i) n[t] = i[t];
                        delete n.rest
                    }
                    var o = new s;
                    return u(o, o.head, e), r(e, o, n, o.head, 0), function (e) {
                        var n = [], i = e.head.next;
                        for (; i !== e.tail;) n.push(i.value), i = i.next;
                        return n
                    }(o)
                },
                hooks: {
                    all: {}, add: function (e, n) {
                        var i = o.hooks.all;
                        i[e] = i[e] || [], i[e].push(n)
                    }, run: function (e, n) {
                        var i = o.hooks.all[e];
                        if (i && i.length) for (var t, a = 0; t = i[a++];) t(n)
                    }
                },
                Token: a
            };

            function a(e, n, i, t) {
                this.type = e, this.content = n, this.alias = i, this.length = 0 | (t || "").length
            }

            function c(e, n, i, t) {
                e.lastIndex = n;
                var o = e.exec(i);
                if (o && t && o[1]) {
                    var a = o[1].length;
                    o.index += a, o[0] = o[0].slice(a)
                }
                return o
            }

            function r(e, n, i, t, s, f) {
                for (var g in i) if (i.hasOwnProperty(g) && i[g]) {
                    var p = i[g];
                    p = Array.isArray(p) ? p : [p];
                    for (var l = 0; l < p.length; ++l) {
                        if (f && f.cause == g + "," + l) return;
                        var v = p[l], h = v.inside, _ = !!v.lookbehind, m = !!v.greedy, b = v.alias;
                        if (m && !v.pattern.global) {
                            var k = v.pattern.toString().match(/[imsuy]*$/)[0];
                            v.pattern = RegExp(v.pattern.source, k + "g")
                        }
                        for (var w = v.pattern || v, y = t.next, x = s; y !== n.tail && !(f && x >= f.reach); x += y.value.length, y = y.next) {
                            var $ = y.value;
                            if (n.length > e.length) return;
                            if (!($ instanceof a)) {
                                var A, S = 1;
                                if (m) {
                                    if (!(A = c(w, x, e, _)) || A.index >= e.length) break;
                                    var E = A.index, z = A.index + A[0].length, T = x;
                                    for (T += y.value.length; E >= T;) T += (y = y.next).value.length;
                                    if (x = T -= y.value.length, y.value instanceof a) continue;
                                    for (var R = y; R !== n.tail && (T < z || "string" == typeof R.value); R = R.next) S++, T += R.value.length;
                                    S--, $ = e.slice(x, T), A.index -= x
                                } else if (!(A = c(w, 0, $, _))) continue;
                                E = A.index;
                                var L = A[0], F = $.slice(0, E), C = $.slice(E + L.length), j = x + $.length;
                                f && j > f.reach && (f.reach = j);
                                var O = y.prev;
                                if (F && (O = u(n, O, F), x += F.length), d(n, O, S), y = u(n, O, new a(g, h ? o.tokenize(L, h) : L, b, L)), C && u(n, y, C), S > 1) {
                                    var q = {cause: g + "," + l, reach: j};
                                    r(e, n, i, y.prev, x, q), f && q.reach > f.reach && (f.reach = q.reach)
                                }
                            }
                        }
                    }
                }
            }

            function s() {
                var e = {value: null, prev: null, next: null}, n = {value: null, prev: e, next: null};
                e.next = n, this.head = e, this.tail = n, this.length = 0
            }

            function u(e, n, i) {
                var t = n.next, o = {value: i, prev: n, next: t};
                return n.next = o, t.prev = o, e.length++, o
            }

            function d(e, n, i) {
                for (var t = n.next, o = 0; o < i && t !== e.tail; o++) t = t.next;
                n.next = t, t.prev = n, e.length -= o
            }

            if (e.Prism = o, a.stringify = function e(n, i) {
                if ("string" == typeof n) return n;
                if (Array.isArray(n)) {
                    var t = "";
                    return n.forEach((function (n) {
                        t += e(n, i)
                    })), t
                }
                var a = {
                    type: n.type,
                    content: e(n.content, i),
                    tag: "span",
                    classes: ["token", n.type],
                    attributes: {},
                    language: i
                }, c = n.alias;
                c && (Array.isArray(c) ? Array.prototype.push.apply(a.classes, c) : a.classes.push(c)), o.hooks.run("wrap", a);
                var r = "";
                for (var s in a.attributes) r += " " + s + '="' + (a.attributes[s] || "").replace(/"/g, "&quot;") + '"';
                return "<" + a.tag + ' class="' + a.classes.join(" ") + '"' + r + ">" + a.content + "</" + a.tag + ">"
            }, !e.document) return e.addEventListener ? (o.disableWorkerMessageHandler || e.addEventListener("message", (function (n) {
                var i = JSON.parse(n.data), t = i.language, a = i.code, c = i.immediateClose;
                e.postMessage(o.highlight(a, o.languages[t], t)), c && e.close()
            }), !1), o) : o;
            var f = o.util.currentScript();

            function g() {
                o.manual || o.highlightAll()
            }

            if (f && (o.filename = f.src, f.hasAttribute("data-manual") && (o.manual = !0)), !o.manual) {
                var p = document.readyState;
                "loading" === p || "interactive" === p && f && f.defer ? document.addEventListener("DOMContentLoaded", g) : window.requestAnimationFrame ? window.requestAnimationFrame(g) : window.setTimeout(g, 16)
            }
            return o
        }("undefined" != typeof window ? window : "undefined" != typeof WorkerGlobalScope && self instanceof WorkerGlobalScope ? self : {});
        e.exports && (e.exports = n), void 0 !== cn && (cn.Prism = n), n.languages.markup = {
            comment: {pattern: /<!--(?:(?!<!--)[\s\S])*?-->/, greedy: !0},
            prolog: {pattern: /<\?[\s\S]+?\?>/, greedy: !0},
            doctype: {
                pattern: /<!DOCTYPE(?:[^>"'[\]]|"[^"]*"|'[^']*')+(?:\[(?:[^<"'\]]|"[^"]*"|'[^']*'|<(?!!--)|<!--(?:[^-]|-(?!->))*-->)*\]\s*)?>/i,
                greedy: !0,
                inside: {
                    "internal-subset": {
                        pattern: /(^[^\[]*\[)[\s\S]+(?=\]>$)/,
                        lookbehind: !0,
                        greedy: !0,
                        inside: null
                    },
                    string: {pattern: /"[^"]*"|'[^']*'/, greedy: !0},
                    punctuation: /^<!|>$|[[\]]/,
                    "doctype-tag": /^DOCTYPE/i,
                    name: /[^\s<>'"]+/
                }
            },
            cdata: {pattern: /<!\[CDATA\[[\s\S]*?\]\]>/i, greedy: !0},
            tag: {
                pattern: /<\/?(?!\d)[^\s>\/=$<%]+(?:\s(?:\s*[^\s>\/=]+(?:\s*=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+(?=[\s>]))|(?=[\s/>])))+)?\s*\/?>/,
                greedy: !0,
                inside: {
                    tag: {pattern: /^<\/?[^\s>\/]+/, inside: {punctuation: /^<\/?/, namespace: /^[^\s>\/:]+:/}},
                    "special-attr": [],
                    "attr-value": {
                        pattern: /=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+)/,
                        inside: {
                            punctuation: [{pattern: /^=/, alias: "attr-equals"}, {
                                pattern: /^(\s*)["']|["']$/,
                                lookbehind: !0
                            }]
                        }
                    },
                    punctuation: /\/?>/,
                    "attr-name": {pattern: /[^\s>\/]+/, inside: {namespace: /^[^\s>\/:]+:/}}
                }
            },
            entity: [{pattern: /&[\da-z]{1,8};/i, alias: "named-entity"}, /&#x?[\da-f]{1,8};/i]
        }, n.languages.markup.tag.inside["attr-value"].inside.entity = n.languages.markup.entity, n.languages.markup.doctype.inside["internal-subset"].inside = n.languages.markup, n.hooks.add("wrap", (function (e) {
            "entity" === e.type && (e.attributes.title = e.content.replace(/&amp;/, "&"))
        })), Object.defineProperty(n.languages.markup.tag, "addInlined", {
            value: function (e, i) {
                var t = {};
                t["language-" + i] = {
                    pattern: /(^<!\[CDATA\[)[\s\S]+?(?=\]\]>$)/i,
                    lookbehind: !0,
                    inside: n.languages[i]
                }, t.cdata = /^<!\[CDATA\[|\]\]>$/i;
                var o = {"included-cdata": {pattern: /<!\[CDATA\[[\s\S]*?\]\]>/i, inside: t}};
                o["language-" + i] = {pattern: /[\s\S]+/, inside: n.languages[i]};
                var a = {};
                a[e] = {
                    pattern: RegExp(/(<__[^>]*>)(?:<!\[CDATA\[(?:[^\]]|\](?!\]>))*\]\]>|(?!<!\[CDATA\[)[\s\S])*?(?=<\/__>)/.source.replace(/__/g, (function () {
                        return e
                    })), "i"), lookbehind: !0, greedy: !0, inside: o
                }, n.languages.insertBefore("markup", "cdata", a)
            }
        }), Object.defineProperty(n.languages.markup.tag, "addAttribute", {
            value: function (e, i) {
                n.languages.markup.tag.inside["special-attr"].push({
                    pattern: RegExp(/(^|["'\s])/.source + "(?:" + e + ")" + /\s*=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+(?=[\s>]))/.source, "i"),
                    lookbehind: !0,
                    inside: {
                        "attr-name": /^[^\s=]+/,
                        "attr-value": {
                            pattern: /=[\s\S]+/,
                            inside: {
                                value: {
                                    pattern: /(^=\s*(["']|(?!["'])))\S[\s\S]*(?=\2$)/,
                                    lookbehind: !0,
                                    alias: [i, "language-" + i],
                                    inside: n.languages[i]
                                }, punctuation: [{pattern: /^=/, alias: "attr-equals"}, /"|'/]
                            }
                        }
                    }
                })
            }
        }), n.languages.html = n.languages.markup, n.languages.mathml = n.languages.markup, n.languages.svg = n.languages.markup, n.languages.xml = n.languages.extend("markup", {}), n.languages.ssml = n.languages.xml, n.languages.atom = n.languages.xml, n.languages.rss = n.languages.xml, function (e) {
            var n = /(?:"(?:\\(?:\r\n|[\s\S])|[^"\\\r\n])*"|'(?:\\(?:\r\n|[\s\S])|[^'\\\r\n])*')/;
            e.languages.css = {
                comment: /\/\*[\s\S]*?\*\//,
                atrule: {
                    pattern: RegExp("@[\\w-](?:" + /[^;{\s"']|\s+(?!\s)/.source + "|" + n.source + ")*?" + /(?:;|(?=\s*\{))/.source),
                    inside: {
                        rule: /^@[\w-]+/,
                        "selector-function-argument": {
                            pattern: /(\bselector\s*\(\s*(?![\s)]))(?:[^()\s]|\s+(?![\s)])|\((?:[^()]|\([^()]*\))*\))+(?=\s*\))/,
                            lookbehind: !0,
                            alias: "selector"
                        },
                        keyword: {pattern: /(^|[^\w-])(?:and|not|only|or)(?![\w-])/, lookbehind: !0}
                    }
                },
                url: {
                    pattern: RegExp("\\burl\\((?:" + n.source + "|" + /(?:[^\\\r\n()"']|\\[\s\S])*/.source + ")\\)", "i"),
                    greedy: !0,
                    inside: {
                        function: /^url/i,
                        punctuation: /^\(|\)$/,
                        string: {pattern: RegExp("^" + n.source + "$"), alias: "url"}
                    }
                },
                selector: {
                    pattern: RegExp("(^|[{}\\s])[^{}\\s](?:[^{};\"'\\s]|\\s+(?![\\s{])|" + n.source + ")*(?=\\s*\\{)"),
                    lookbehind: !0
                },
                string: {pattern: n, greedy: !0},
                property: {
                    pattern: /(^|[^-\w\xA0-\uFFFF])(?!\s)[-_a-z\xA0-\uFFFF](?:(?!\s)[-\w\xA0-\uFFFF])*(?=\s*:)/i,
                    lookbehind: !0
                },
                important: /!important\b/i,
                function: {pattern: /(^|[^-a-z0-9])[-a-z0-9]+(?=\()/i, lookbehind: !0},
                punctuation: /[(){};:,]/
            }, e.languages.css.atrule.inside.rest = e.languages.css;
            var i = e.languages.markup;
            i && (i.tag.addInlined("style", "css"), i.tag.addAttribute("style", "css"))
        }(n), n.languages.clike = {
            comment: [{
                pattern: /(^|[^\\])\/\*[\s\S]*?(?:\*\/|$)/,
                lookbehind: !0,
                greedy: !0
            }, {pattern: /(^|[^\\:])\/\/.*/, lookbehind: !0, greedy: !0}],
            string: {pattern: /(["'])(?:\\(?:\r\n|[\s\S])|(?!\1)[^\\\r\n])*\1/, greedy: !0},
            "class-name": {
                pattern: /(\b(?:class|extends|implements|instanceof|interface|new|trait)\s+|\bcatch\s+\()[\w.\\]+/i,
                lookbehind: !0,
                inside: {punctuation: /[.\\]/}
            },
            keyword: /\b(?:break|catch|continue|do|else|finally|for|function|if|in|instanceof|new|null|return|throw|try|while)\b/,
            boolean: /\b(?:false|true)\b/,
            function: /\b\w+(?=\()/,
            number: /\b0x[\da-f]+\b|(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:e[+-]?\d+)?/i,
            operator: /[<>]=?|[!=]=?=?|--?|\+\+?|&&?|\|\|?|[?*/~^%]/,
            punctuation: /[{}[\];(),.:]/
        }, n.languages.javascript = n.languages.extend("clike", {
            "class-name": [n.languages.clike["class-name"], {
                pattern: /(^|[^$\w\xA0-\uFFFF])(?!\s)[_$A-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\.(?:constructor|prototype))/,
                lookbehind: !0
            }],
            keyword: [{
                pattern: /((?:^|\})\s*)catch\b/,
                lookbehind: !0
            }, {
                pattern: /(^|[^.]|\.\.\.\s*)\b(?:as|assert(?=\s*\{)|async(?=\s*(?:function\b|\(|[$\w\xA0-\uFFFF]|$))|await|break|case|class|const|continue|debugger|default|delete|do|else|enum|export|extends|finally(?=\s*(?:\{|$))|for|from(?=\s*(?:['"]|$))|function|(?:get|set)(?=\s*(?:[#\[$\w\xA0-\uFFFF]|$))|if|implements|import|in|instanceof|interface|let|new|null|of|package|private|protected|public|return|static|super|switch|this|throw|try|typeof|undefined|var|void|while|with|yield)\b/,
                lookbehind: !0
            }],
            function: /#?(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*(?:\.\s*(?:apply|bind|call)\s*)?\()/,
            number: {
                pattern: RegExp(/(^|[^\w$])/.source + "(?:" + /NaN|Infinity/.source + "|" + /0[bB][01]+(?:_[01]+)*n?/.source + "|" + /0[oO][0-7]+(?:_[0-7]+)*n?/.source + "|" + /0[xX][\dA-Fa-f]+(?:_[\dA-Fa-f]+)*n?/.source + "|" + /\d+(?:_\d+)*n/.source + "|" + /(?:\d+(?:_\d+)*(?:\.(?:\d+(?:_\d+)*)?)?|\.\d+(?:_\d+)*)(?:[Ee][+-]?\d+(?:_\d+)*)?/.source + ")" + /(?![\w$])/.source),
                lookbehind: !0
            },
            operator: /--|\+\+|\*\*=?|=>|&&=?|\|\|=?|[!=]==|<<=?|>>>?=?|[-+*/%&|^!=<>]=?|\.{3}|\?\?=?|\?\.?|[~:]/
        }), n.languages.javascript["class-name"][0].pattern = /(\b(?:class|extends|implements|instanceof|interface|new)\s+)[\w.\\]+/, n.languages.insertBefore("javascript", "keyword", {
            regex: {
                pattern: RegExp(/((?:^|[^$\w\xA0-\uFFFF."'\])\s]|\b(?:return|yield))\s*)/.source + /\//.source + "(?:" + /(?:\[(?:[^\]\\\r\n]|\\.)*\]|\\.|[^/\\\[\r\n])+\/[dgimyus]{0,7}/.source + "|" + /(?:\[(?:[^[\]\\\r\n]|\\.|\[(?:[^[\]\\\r\n]|\\.|\[(?:[^[\]\\\r\n]|\\.)*\])*\])*\]|\\.|[^/\\\[\r\n])+\/[dgimyus]{0,7}v[dgimyus]{0,7}/.source + ")" + /(?=(?:\s|\/\*(?:[^*]|\*(?!\/))*\*\/)*(?:$|[\r\n,.;:})\]]|\/\/))/.source),
                lookbehind: !0,
                greedy: !0,
                inside: {
                    "regex-source": {
                        pattern: /^(\/)[\s\S]+(?=\/[a-z]*$)/,
                        lookbehind: !0,
                        alias: "language-regex",
                        inside: n.languages.regex
                    }, "regex-delimiter": /^\/|\/$/, "regex-flags": /^[a-z]+$/
                }
            },
            "function-variable": {
                pattern: /#?(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*[=:]\s*(?:async\s*)?(?:\bfunction\b|(?:\((?:[^()]|\([^()]*\))*\)|(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*)\s*=>))/,
                alias: "function"
            },
            parameter: [{
                pattern: /(function(?:\s+(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*)?\s*\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\))/,
                lookbehind: !0,
                inside: n.languages.javascript
            }, {
                pattern: /(^|[^$\w\xA0-\uFFFF])(?!\s)[_$a-z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*=>)/i,
                lookbehind: !0,
                inside: n.languages.javascript
            }, {
                pattern: /(\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\)\s*=>)/,
                lookbehind: !0,
                inside: n.languages.javascript
            }, {
                pattern: /((?:\b|\s|^)(?!(?:as|async|await|break|case|catch|class|const|continue|debugger|default|delete|do|else|enum|export|extends|finally|for|from|function|get|if|implements|import|in|instanceof|interface|let|new|null|of|package|private|protected|public|return|set|static|super|switch|this|throw|try|typeof|undefined|var|void|while|with|yield)(?![$\w\xA0-\uFFFF]))(?:(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*\s*)\(\s*|\]\s*\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\)\s*\{)/,
                lookbehind: !0,
                inside: n.languages.javascript
            }],
            constant: /\b[A-Z](?:[A-Z_]|\dx?)*\b/
        }), n.languages.insertBefore("javascript", "string", {
            hashbang: {
                pattern: /^#!.*/,
                greedy: !0,
                alias: "comment"
            },
            "template-string": {
                pattern: /`(?:\\[\s\S]|\$\{(?:[^{}]|\{(?:[^{}]|\{[^}]*\})*\})+\}|(?!\$\{)[^\\`])*`/,
                greedy: !0,
                inside: {
                    "template-punctuation": {pattern: /^`|`$/, alias: "string"},
                    interpolation: {
                        pattern: /((?:^|[^\\])(?:\\{2})*)\$\{(?:[^{}]|\{(?:[^{}]|\{[^}]*\})*\})+\}/,
                        lookbehind: !0,
                        inside: {
                            "interpolation-punctuation": {pattern: /^\$\{|\}$/, alias: "punctuation"},
                            rest: n.languages.javascript
                        }
                    },
                    string: /[\s\S]+/
                }
            },
            "string-property": {
                pattern: /((?:^|[,{])[ \t]*)(["'])(?:\\(?:\r\n|[\s\S])|(?!\2)[^\\\r\n])*\2(?=\s*:)/m,
                lookbehind: !0,
                greedy: !0,
                alias: "property"
            }
        }), n.languages.insertBefore("javascript", "operator", {
            "literal-property": {
                pattern: /((?:^|[,{])[ \t]*)(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*:)/m,
                lookbehind: !0,
                alias: "property"
            }
        }), n.languages.markup && (n.languages.markup.tag.addInlined("script", "javascript"), n.languages.markup.tag.addAttribute(/on(?:abort|blur|change|click|composition(?:end|start|update)|dblclick|error|focus(?:in|out)?|key(?:down|up)|load|mouse(?:down|enter|leave|move|out|over|up)|reset|resize|scroll|select|slotchange|submit|unload|wheel)/.source, "javascript")), n.languages.js = n.languages.javascript, function () {
            if (void 0 !== n && "undefined" != typeof document) {
                Element.prototype.matches || (Element.prototype.matches = Element.prototype.msMatchesSelector || Element.prototype.webkitMatchesSelector);
                var e = {
                        js: "javascript",
                        py: "python",
                        rb: "ruby",
                        ps1: "powershell",
                        psm1: "powershell",
                        sh: "bash",
                        bat: "batch",
                        h: "c",
                        tex: "latex"
                    }, i = "data-src-status", t = "loading", o = "loaded",
                    a = "pre[data-src]:not([" + i + '="' + o + '"]):not([' + i + '="' + t + '"])';
                n.hooks.add("before-highlightall", (function (e) {
                    e.selector += ", " + a
                })), n.hooks.add("before-sanity-check", (function (c) {
                    var r = c.element;
                    if (r.matches(a)) {
                        c.code = "", r.setAttribute(i, t);
                        var s = r.appendChild(document.createElement("CODE"));
                        s.textContent = "Loading…";
                        var u = r.getAttribute("data-src"), d = c.language;
                        if ("none" === d) {
                            var f = (/\.(\w+)$/.exec(u) || [, "none"])[1];
                            d = e[f] || f
                        }
                        n.util.setLanguage(s, d), n.util.setLanguage(r, d);
                        var g = n.plugins.autoloader;
                        g && g.loadLanguages(d), function (e, n, i) {
                            var t = new XMLHttpRequest;
                            t.open("GET", e, !0), t.onreadystatechange = function () {
                                4 == t.readyState && (t.status < 400 && t.responseText ? n(t.responseText) : t.status >= 400 ? i("✖ Error " + t.status + " while fetching file: " + t.statusText) : i("✖ Error: File does not exist or is empty"))
                            }, t.send(null)
                        }(u, (function (e) {
                            r.setAttribute(i, o);
                            var t = function (e) {
                                var n = /^\s*(\d+)\s*(?:(,)\s*(?:(\d+)\s*)?)?$/.exec(e || "");
                                if (n) {
                                    var i = Number(n[1]), t = n[2], o = n[3];
                                    return t ? o ? [i, Number(o)] : [i, void 0] : [i, i]
                                }
                            }(r.getAttribute("data-range"));
                            if (t) {
                                var a = e.split(/\r\n?|\n/g), c = t[0], u = null == t[1] ? a.length : t[1];
                                c < 0 && (c += a.length), c = Math.max(0, Math.min(c - 1, a.length)), u < 0 && (u += a.length), u = Math.max(0, Math.min(u, a.length)), e = a.slice(c, u).join("\n"), r.hasAttribute("data-start") || r.setAttribute("data-start", String(c + 1))
                            }
                            s.textContent = e, n.highlightElement(s)
                        }), (function (e) {
                            r.setAttribute(i, "failed"), s.textContent = e
                        }))
                    }
                })), n.plugins.fileHighlight = {
                    highlight: function (e) {
                        for (var i, t = (e || document).querySelectorAll(a), o = 0; i = t[o++];) n.highlightElement(i)
                    }
                };
                var c = !1;
                n.fileHighlight = function () {
                    c || (console.warn("Prism.fileHighlight is deprecated. Use `Prism.plugins.fileHighlight.highlight` instead."), c = !0), n.plugins.fileHighlight.highlight.apply(this, arguments)
                }
            }
        }()
    }(sn);
    var un = rn(sn.exports);
    !function (e) {
        function n(e, n) {
            return "___" + e.toUpperCase() + n + "___"
        }

        Object.defineProperties(e.languages["markup-templating"] = {}, {
            buildPlaceholders: {
                value: function (i, t, o, a) {
                    if (i.language === t) {
                        var c = i.tokenStack = [];
                        i.code = i.code.replace(o, (function (e) {
                            if ("function" == typeof a && !a(e)) return e;
                            for (var o, r = c.length; -1 !== i.code.indexOf(o = n(t, r));) ++r;
                            return c[r] = e, o
                        })), i.grammar = e.languages.markup
                    }
                }
            }, tokenizePlaceholders: {
                value: function (i, t) {
                    if (i.language === t && i.tokenStack) {
                        i.grammar = e.languages[t];
                        var o = 0, a = Object.keys(i.tokenStack);
                        !function c(r) {
                            for (var s = 0; s < r.length && !(o >= a.length); s++) {
                                var u = r[s];
                                if ("string" == typeof u || u.content && "string" == typeof u.content) {
                                    var d = a[o], f = i.tokenStack[d], g = "string" == typeof u ? u : u.content,
                                        p = n(t, d), l = g.indexOf(p);
                                    if (l > -1) {
                                        ++o;
                                        var v = g.substring(0, l),
                                            h = new e.Token(t, e.tokenize(f, i.grammar), "language-" + t, f),
                                            _ = g.substring(l + p.length), m = [];
                                        v && m.push.apply(m, c([v])), m.push(h), _ && m.push.apply(m, c([_])), "string" == typeof u ? r.splice.apply(r, [s, 1].concat(m)) : u.content = m
                                    }
                                } else u.content && c(u.content)
                            }
                            return r
                        }(i.tokens)
                    }
                }
            }
        })
    }(Prism);
    const dn = {}, fn = {
        markdown: e => ({url: e}),
        mermaid: e => ({url: e}),
        iframe: (e, n) => ({html: `<iframe src="${e}" ${n || "width=100% height=400"}></iframe>`}),
        video: (e, n) => ({html: `<video src="${e}" ${n || "controls"}>Not Support</video>`}),
        audio: (e, n) => ({html: `<audio src="${e}" ${n || "controls"}>Not Support</audio>`}),
        code(e, n) {
            let i = e.match(/\.(\w+)$/);
            return i = n || i && i[1], "md" === i && (i = "markdown"), {url: e, lang: i}
        }
    };

    class gn {
        constructor(n, t) {
            this.config = n, this.router = t, this.cacheTree = {}, this.toc = [], this.cacheTOC = {}, this.linkTarget = n.externalLinkTarget || "_blank", this.linkRel = "_blank" === this.linkTarget ? n.externalLinkRel || "noopener" : "", this.contentBase = t.getBasePath();
            const a = this._initRenderer();
            let c;
            this.heading = a.heading;
            const r = n.markdown || {};
            o(r) ? c = r(We, a) : (We.setOptions(Object.assign(r, {renderer: Object.assign(a, r.renderer)})), c = We), this._marked = c, this.compile = t => {
                let o = !0;
                const a = e((e => {
                    o = !1;
                    let a = "";
                    return t ? (a = i(t) ? c(t) : c.parser(t), a = n.noEmoji ? a : on(a, n.nativeEmoji), nn.clear(), a) : t
                }))(t), r = this.router.parse().file;
                return o ? this.toc = this.cacheTOC[r] : this.cacheTOC[r] = [...this.toc], a
            }
        }

        compileEmbed(e, n) {
            const {str: i, config: t} = an(n);
            let o;
            if (n = i, t.include) {
                let i;
                if (A(e) || (e = L(this.contentBase, E(this.router.getCurrentPath()), e)), t.type && (i = fn[t.type])) o = i.call(this, e, n), o.type = t.type; else {
                    let i = "code";
                    /\.(md|markdown)/.test(e) ? i = "markdown" : /\.mmd/.test(e) ? i = "mermaid" : /\.html?/.test(e) ? i = "iframe" : /\.(mp4|ogg)/.test(e) ? i = "video" : /\.mp3/.test(e) && (i = "audio"), o = fn[i].call(this, e, n), o.type = i
                }
                return o.fragment = t.fragment, o
            }
        }

        _matchNotCompileLink(e) {
            const n = this.config.noCompileLinks || [];
            for (const i of n) {
                if ((dn[i] || (dn[i] = new RegExp(`^${i}$`))).test(e)) return e
            }
        }

        _initRenderer() {
            const e = new We.Renderer, {linkTarget: n, linkRel: i, router: t, contentBase: o} = this, a = this, c = {};
            return c.heading = e.heading = function (e, n) {
                let {str: i, config: o} = an(e);
                const c = {level: n, title: i}, {content: r, ignoreAllSubs: s, ignoreSubHeading: u} = function () {
                    let e, n, i = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return /<!-- {docsify-ignore} -->/g.test(i) && (i = i.replace("\x3c!-- {docsify-ignore} --\x3e", ""), n = !0), /{docsify-ignore}/g.test(i) && (i = i.replace("{docsify-ignore}", ""), n = !0), /<!-- {docsify-ignore-all} -->/g.test(i) && (i = i.replace("\x3c!-- {docsify-ignore-all} --\x3e", ""), e = !0), /{docsify-ignore-all}/g.test(i) && (i = i.replace("{docsify-ignore-all}", ""), e = !0), {
                        content: i,
                        ignoreAllSubs: e,
                        ignoreSubHeading: n
                    }
                }(i);
                i = r.trim(), c.title = function () {
                    return (arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "").replace(/(<\/?a.*?>)/gi, "")
                }(i), c.ignoreAllSubs = s, c.ignoreSubHeading = u;
                const d = nn(o.id || i), f = t.toURL(t.getCurrentPath(), {id: d});
                return c.slug = f, a.toc.push(c), `<h${n} id="${d}" tabindex="-1"><a href="${f}" data-id="${d}" class="anchor"><span>${i}</span></a></h${n}>`
            }, c.code = (e => {
                let {renderer: n} = e;
                return n.code = function (e) {
                    let n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : "markup";
                    const i = un.languages[n] || un.languages.markup;
                    return `<pre data-lang="${n}"><code class="lang-${n}" tabindex="0">${un.highlight(e.replace(/@DOCSIFY_QM@/g, "`"), i, n)}</code></pre>`
                }
            })({renderer: e}), c.link = (e => {
                let {renderer: n, router: i, linkTarget: t, linkRel: o, compilerClass: a} = e;
                return n.link = function (e) {
                    let n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : "",
                        c = arguments.length > 2 ? arguments[2] : void 0;
                    const r = [], {str: s, config: u} = an(n);
                    return t = u.target || t, o = "_blank" === t ? a.config.externalLinkRel || "noopener" : "", n = s, A(e) || a._matchNotCompileLink(e) || u.ignore ? (A(e) || "./" !== e.slice(0, 2) || (e = document.URL.replace(/\/(?!.*\/).*/, "/").replace("#/./", "") + e), r.push(0 === e.indexOf("mailto:") ? "" : `target="${t}"`), r.push(0 === e.indexOf("mailto:") ? "" : "" !== o ? ` rel="${o}"` : "")) : (e === a.config.homepage && (e = "README"), e = i.toURL(e, null, i.getCurrentPath()), u.target && 0 !== e.indexOf("mailto:") && r.push(`target="${t}"`)), u.disabled && (r.push("disabled"), e = "javascript:void(0)"), u.class && r.push(`class="${u.class}"`), u.id && r.push(`id="${u.id}"`), n && r.push(`title="${n}"`), `<a href="${e}" ${r.join(" ")}>${c}</a>`
                }
            })({renderer: e, router: t, linkTarget: n, linkRel: i, compilerClass: a}), c.paragraph = (e => {
                let {renderer: n} = e;
                return n.paragraph = e => {
                    let n;
                    return n = /^!&gt;/.test(e) ? Ye("tip", e) : /^\?&gt;/.test(e) ? Ye("warn", e) : `<p>${e}</p>`, n
                }
            })({renderer: e}), c.image = (e => {
                let {renderer: n, contentBase: i, router: t} = e;
                return n.image = (e, n, o) => {
                    let a = e;
                    const c = [], {str: r, config: s} = an(n);
                    if (n = r, s["no-zoom"] && c.push("data-no-zoom"), n && c.push(`title="${n}"`), s.size) {
                        const [e, n] = s.size.split("x");
                        n ? c.push(`width="${e}" height="${n}"`) : c.push(`width="${e}"`)
                    }
                    return s.class && c.push(`class="${s.class}"`), s.id && c.push(`id="${s.id}"`), A(e) || (a = L(i, E(t.getCurrentPath()), e)), `<img src="${a}" data-origin="${e}" alt="${o}" ${c.join(" ")} />`
                }
            })({renderer: e, contentBase: o, router: t}), c.list = (e => {
                let {renderer: n} = e;
                return n.list = (e, n, i) => {
                    const t = n ? "ol" : "ul";
                    return `<${t} ${[/<li class="task-list-item">/.test(e.split('class="task-list"')[0]) ? 'class="task-list"' : "", i && i > 1 ? `start="${i}"` : ""].join(" ").trim()}>${e}</${t}>`
                }
            })({renderer: e}), c.listitem = (e => {
                let {renderer: n} = e;
                return n.listitem = e => /^(<input.*type="checkbox"[^>]*>)/.test(e) ? `<li class="task-list-item"><label>${e}</label></li>` : `<li>${e}</li>`
            })({renderer: e}), e.origin = c, e
        }

        sidebar(e, n) {
            const {toc: i} = this, t = this.router.getCurrentPath();
            let o = "";
            if (e) o = this.compile(e); else {
                for (let e = 0; e < i.length; e++) if (i[e].ignoreSubHeading) {
                    const n = i[e].level;
                    i.splice(e, 1);
                    for (let t = e; t < i.length && n < i[t].level; t++) i.splice(t, 1) && t-- && e++;
                    e--
                }
                const e = this.cacheTree[t] || Xe(i, n);
                o = Ge(e, "<ul>{inner}</ul>"), this.cacheTree[t] = e
            }
            return o
        }

        subSidebar(e) {
            if (!e) return void (this.toc = []);
            const n = this.router.getCurrentPath(), {cacheTree: i, toc: t} = this;
            t[0] && t[0].ignoreAllSubs && t.splice(0), t[0] && 1 === t[0].level && t.shift();
            for (let e = 0; e < t.length; e++) t[e].ignoreSubHeading && t.splice(e, 1) && e--;
            const o = i[n] || Xe(t, e);
            return i[n] = o, this.toc = [], Ge(o)
        }

        header(e, n) {
            return this.heading(e, n)
        }

        article(e) {
            return this.compile(e)
        }

        cover(e) {
            const n = this.toc.slice(), i = this.compile(e);
            return this.toc = n.slice(), i
        }
    }

    var pn, ln, vn = function () {
            function e(e, n) {
                for (var i = 0; i < n.length; i++) {
                    var t = n[i];
                    t.enumerable = t.enumerable || !1, t.configurable = !0, "value" in t && (t.writable = !0), Object.defineProperty(e, t.key, t)
                }
            }

            return function (n, i, t) {
                return i && e(n.prototype, i), t && e(n, t), n
            }
        }(),
        hn = (pn = ["", ""], ln = ["", ""], Object.freeze(Object.defineProperties(pn, {raw: {value: Object.freeze(ln)}})));
    var _n = function () {
        function e() {
            for (var n = this, i = arguments.length, t = Array(i), o = 0; o < i; o++) t[o] = arguments[o];
            return function (e, n) {
                if (!(e instanceof n)) throw new TypeError("Cannot call a class as a function")
            }(this, e), this.tag = function (e) {
                for (var i = arguments.length, t = Array(i > 1 ? i - 1 : 0), o = 1; o < i; o++) t[o - 1] = arguments[o];
                return "function" == typeof e ? n.interimTag.bind(n, e) : "string" == typeof e ? n.transformEndResult(e) : (e = e.map(n.transformString.bind(n)), n.transformEndResult(e.reduce(n.processSubstitutions.bind(n, t))))
            }, t.length > 0 && Array.isArray(t[0]) && (t = t[0]), this.transformers = t.map((function (e) {
                return "function" == typeof e ? e() : e
            })), this.tag
        }

        return vn(e, [{
            key: "interimTag", value: function (e, n) {
                for (var i = arguments.length, t = Array(i > 2 ? i - 2 : 0), o = 2; o < i; o++) t[o - 2] = arguments[o];
                return this.tag(hn, e.apply(void 0, [n].concat(t)))
            }
        }, {
            key: "processSubstitutions", value: function (e, n, i) {
                var t = this.transformSubstitution(e.shift(), n);
                return "".concat(n, t, i)
            }
        }, {
            key: "transformString", value: function (e) {
                return this.transformers.reduce((function (e, n) {
                    return n.onString ? n.onString(e) : e
                }), e)
            }
        }, {
            key: "transformSubstitution", value: function (e, n) {
                return this.transformers.reduce((function (e, i) {
                    return i.onSubstitution ? i.onSubstitution(e, n) : e
                }), e)
            }
        }, {
            key: "transformEndResult", value: function (e) {
                return this.transformers.reduce((function (e, n) {
                    return n.onEndResult ? n.onEndResult(e) : e
                }), e)
            }
        }]), e
    }(), mn = function () {
        var e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
        return {
            onEndResult: function (n) {
                if ("" === e) return n.trim();
                if ("start" === (e = e.toLowerCase()) || "left" === e) return n.replace(/^\s*/, "");
                if ("end" === e || "right" === e) return n.replace(/\s*$/, "");
                throw new Error("Side not supported: " + e)
            }
        }
    };
    var bn = function () {
        var e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "initial";
        return {
            onEndResult: function (n) {
                if ("initial" === e) {
                    var i = n.match(/^[^\S\n]*(?=\S)/gm), t = i && Math.min.apply(Math, function (e) {
                        if (Array.isArray(e)) {
                            for (var n = 0, i = Array(e.length); n < e.length; n++) i[n] = e[n];
                            return i
                        }
                        return Array.from(e)
                    }(i.map((function (e) {
                        return e.length
                    }))));
                    if (t) {
                        var o = new RegExp("^.{" + t + "}", "gm");
                        return n.replace(o, "")
                    }
                    return n
                }
                if ("all" === e) return n.replace(/^[^\S\n]+/gm, "");
                throw new Error("Unknown type: " + e)
            }
        }
    }, kn = function (e, n) {
        return {
            onEndResult: function (i) {
                if (null == e || null == n) throw new Error("replaceResultTransformer requires at least 2 arguments.");
                return i.replace(e, n)
            }
        }
    }, wn = function (e, n) {
        return {
            onSubstitution: function (i, t) {
                if (null == e || null == n) throw new Error("replaceSubstitutionTransformer requires at least 2 arguments.");
                return null == i ? i : i.toString().replace(e, n)
            }
        }
    }, yn = {separator: "", conjunction: "", serial: !1}, xn = function () {
        var e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : yn;
        return {
            onSubstitution: function (n, i) {
                if (Array.isArray(n)) {
                    var t = n.length, o = e.separator, a = e.conjunction, c = e.serial, r = i.match(/(\n?[^\S\n]+)$/);
                    if (n = r ? n.join(o + r[1]) : n.join(o + " "), a && t > 1) {
                        var s = n.lastIndexOf(o);
                        n = n.slice(0, s) + (c ? o : "") + " " + a + n.slice(s + 1)
                    }
                }
                return n
            }
        }
    }, $n = function (e) {
        return {
            onSubstitution: function (n, i) {
                return "string" == typeof n && n.includes(e) && (n = n.split(e)), n
            }
        }
    }, An = function (e) {
        return null != e && !Number.isNaN(e) && "boolean" != typeof e
    };
    new _n(xn({separator: ","}), bn, mn), new _n(xn({
        separator: ",",
        conjunction: "and"
    }), bn, mn), new _n(xn({separator: ",", conjunction: "or"}), bn, mn), new _n($n("\n"), (function () {
        return {
            onSubstitution: function (e) {
                return Array.isArray(e) ? e.filter(An) : An(e) ? e : ""
            }
        }
    }), xn, bn, mn), new _n($n("\n"), xn, bn, mn, wn(/&/g, "&amp;"), wn(/</g, "&lt;"), wn(/>/g, "&gt;"), wn(/"/g, "&quot;"), wn(/'/g, "&#x27;"), wn(/`/g, "&#x60;")), new _n(kn(/(?:\n(?:\s*))+/g, " "), mn), new _n(kn(/(?:\n\s*)/g, ""), mn), new _n(xn({separator: ","}), kn(/(?:\s+)/g, " "), mn), new _n(xn({
        separator: ",",
        conjunction: "or"
    }), kn(/(?:\s+)/g, " "), mn), new _n(xn({
        separator: ",",
        conjunction: "and"
    }), kn(/(?:\s+)/g, " "), mn), new _n(xn, bn, mn), new _n(xn, kn(/(?:\s+)/g, " "), mn);
    var Sn = new _n(bn, mn);
    let En, zn;

    function Tn(e) {
        const {loaded: n, total: i, step: t} = e;
        let o;
        !En && function () {
            const e = l("div");
            e.classList.add("progress"), e.setAttribute("role", "progressbar"), e.setAttribute("aria-valuemin", "0"), e.setAttribute("aria-valuemax", "100"), e.setAttribute("aria-label", "Loading..."), v(d, e), En = e
        }(), void 0 !== t ? (o = parseInt(En.style.width || 0, 10) + t, o = o > 80 ? 80 : o) : o = Math.floor(n / i * 100), En.style.opacity = 1, En.style.width = o >= 95 ? "100%" : o + "%", En.setAttribute("aria-valuenow", o >= 95 ? 100 : o), o >= 95 && (clearTimeout(zn), zn = setTimeout((e => {
            En.style.opacity = 0, En.style.width = "0%", En.removeAttribute("aria-valuenow")
        }), 200))
    }

    new _n(bn("all"), mn);
    const Rn = {};

    function Ln(e) {
        let n = arguments.length > 1 && void 0 !== arguments[1] && arguments[1],
            i = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : {};
        const o = new XMLHttpRequest, a = Rn[e];
        if (a) return {then: e => e(a.content, a.opt), abort: t};
        o.open("GET", e);
        for (const e of Object.keys(i)) o.setRequestHeader(e, i[e]);
        return o.send(), {
            then(i) {
                let a = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : t;
                const c = e => ({
                    ok: e.target.status >= 200 && e.target.status < 300,
                    status: e.target.status,
                    statusText: e.target.statusText
                });
                if (n) {
                    const e = setInterval((e => Tn({step: Math.floor(5 * Math.random() + 1)})), 500);
                    o.addEventListener("progress", Tn), o.addEventListener("loadend", (n => {
                        Tn(n), clearInterval(e)
                    }))
                }
                o.addEventListener("error", (e => {
                    a(e, c(e))
                })), o.addEventListener("load", (n => {
                    const t = n.target;
                    if (t.status >= 400) a(n, c(n)); else {
                        if ("string" != typeof t.response) throw new TypeError("Unsupported content type.");
                        const a = Rn[e] = {
                            content: t.response,
                            opt: {updatedAt: o.getResponseHeader("last-modified") ?? ""}
                        };
                        i(a.content, a.opt, c(n))
                    }
                }))
            }, abort: e => 4 !== o.readyState && o.abort()
        }
    }

    const Fn = {};

    function Cn(e, n) {
        let {compiler: i, raw: t = "", fetch: o} = e;
        const a = Fn[t];
        if (a) {
            const e = a.slice();
            return e.links = a.links, n(e)
        }
        const c = i._marked;
        let r = c.lexer(t);
        const s = [], u = c.Lexer.rules.inline.normal.link, d = r.links;
        r.forEach(((e, n) => {
            "paragraph" === e.type && (e.text = e.text.replace(new RegExp(u.source, "g"), ((e, t, o, a) => {
                const c = i.compileEmbed(o, a);
                return c && s.push({index: n, embed: c}), e
            })))
        }));
        const f = [];
        !function (e, n) {
            let i, {embedTokens: t, compile: o, fetch: a} = e, c = 0, r = 0;
            if (!t.length) return n({});
            for (; i = t[c++];) {
                const e = i, a = i => {
                    let a;
                    if (i) if ("markdown" === e.embed.type) {
                        let n = e.embed.url.split("/");
                        n.pop(), n = n.join("/"), i = i.replace(/\[([^[\]]+)\]\(([^)]+)\)/g, (e => {
                            const i = e.indexOf("(");
                            return "(." === e.slice(i, i + 2) ? e.substring(0, i) + `(${window.location.protocol}//${window.location.host}${n}/` + e.substring(i + 1, e.length - 1) + ")" : e
                        })), !0 === (($docsify.frontMatter || {}).installed || !1) && (i = $docsify.frontMatter.parseMarkdown(i)), a = o.lexer(i)
                    } else if ("code" === e.embed.type) {
                        if (e.embed.fragment) {
                            const n = e.embed.fragment,
                                t = new RegExp(`(?:###|\\/\\/\\/)\\s*\\[${n}\\]([\\s\\S]*)(?:###|\\/\\/\\/)\\s*\\[${n}\\]`);
                            i = Sn((i.match(t) || [])[1] || "").trim()
                        }
                        a = o.lexer("```" + e.embed.lang + "\n" + i.replace(/`/g, "@DOCSIFY_QM@") + "\n```\n")
                    } else "mermaid" === e.embed.type ? (a = [{
                        type: "html",
                        text: `<div class="mermaid">\n${i}\n</div>`
                    }], a.links = {}) : (a = [{type: "html", text: i}], a.links = {});
                    n({token: e, embedToken: a}), ++r >= t.length && n({})
                };
                i.embed.url ? Ln(i.embed.url).then(a) : a(i.embed.html)
            }
        }({compile: c, embedTokens: s, fetch: o}, (e => {
            let {embedToken: i, token: o} = e;
            if (o) {
                let e = o.index;
                f.forEach((n => {
                    e > n.start && (e += n.length)
                })), Object.assign(d, i.links), r = r.slice(0, e).concat(i, r.slice(e + 1)), f.push({
                    start: e,
                    length: i.length - 1
                })
            } else Fn[t] = r.concat(), r.links = Fn[t].links = d, n(r)
        }))
    }

    function jn(e) {
        return class extends e {
            #a;

            #c(e) {
                p(e).forEach((e => {
                    !e.title && e.innerText && (e.title = e.innerText)
                }))
            }

            #r() {
                const e = p(".markdown-section>script").filter((e => !/template/.test(e.type)))[0];
                if (!e) return !1;
                const n = e.innerText.trim();
                if (!n) return !1;
                new Function(n)()
            }

            #s(e, n, i) {
                var t, o, a;
                return n = "function" == typeof i ? i(n) : "string" == typeof i ? (o = [], a = 0, (t = i).replace(M, (function (e, n, i) {
                    o.push(t.substring(a, i - 1)), a = i += e.length + 1, o.push((function (n) {
                        return ("00" + ("string" == typeof I[e] ? n[I[e]]() : I[e](n))).slice(-e.length)
                    }))
                })), a !== t.length && o.push(t.substring(a)), function (e) {
                    for (var n = "", i = 0, t = e || new Date; i < o.length; i++) n += "string" == typeof o[i] ? o[i] : o[i](t);
                    return n
                })(new Date(n)) : n, e.replace(/{docsify-updated}/g, n)
            }

            #u(e) {
                const n = this.config, i = g(".markdown-section"),
                    t = "Vue" in window && window.Vue.version && Number(window.Vue.version.charAt(0)), o = e => {
                        const n = Boolean(e.__vue__ && e.__vue__._isVue), i = Boolean(e._vnode && e._vnode.__v_skip);
                        return n || i
                    };
                if ("Vue" in window) {
                    const e = p(".markdown-section > *").filter((e => o(e)));
                    for (const n of e) 2 === t ? n.__vue__.$destroy() : 3 === t && n.__vue_app__.unmount()
                }
                if (this._renderTo(i, e), !n.loadSidebar && this._renderSidebar(), (n.executeScript || "Vue" in window && !1 !== n.executeScript) && this.#r(), "Vue" in window) {
                    const e = n.vueGlobalOptions || {}, a = [], c = Object.keys(n.vueComponents || {});
                    2 === t && c.length && c.forEach((e => {
                        !window.Vue.options.components[e] && window.Vue.component(e, n.vueComponents[e])
                    })), !this.#a && e.data && "function" == typeof e.data && (this.#a = e.data()), a.push(...Object.keys(n.vueMounts || {}).map((e => [g(i, e), n.vueMounts[e]])).filter((e => {
                        let [n, i] = e;
                        return n
                    })));
                    const r = /{{2}[^{}]*}{2}/, s = /<[^>/]+\s([@:]|v-)[\w-:.[\]]+[=>\s]/;
                    if (a.push(...p(".markdown-section > *").filter((e => !a.some((n => {
                        let [i, t] = n;
                        return i === e
                    })))).filter((e => e.tagName.toLowerCase() in (n.vueComponents || {}) || e.querySelector(c.join(",") || null) || r.test(e.outerHTML) || s.test(e.outerHTML))).map((n => {
                        const i = {...e};
                        return this.#a && (i.data = () => this.#a), [n, i]
                    }))), 0 === a.length) return;
                    for (const [e, i] of a) {
                        const a = "data-isvue";
                        if (!(e.matches("pre, :not([v-template]):has(pre), script") || o(e) || e.querySelector(`[${a}]`))) if (e.setAttribute(a, ""), 2 === t) i.el = void 0, new window.Vue(i).$mount(e); else if (3 === t) {
                            const t = window.Vue.createApp(i);
                            c.forEach((e => {
                                const i = n.vueComponents[e];
                                t.component(e, i)
                            })), t.mount(e)
                        }
                    }
                }
            }

            #d(e) {
                const n = s(".app-name-link"), t = e.config.nameLink, o = e.route.path;
                if (n) if (i(e.config.nameLink)) n.setAttribute("href", t); else if ("object" == typeof t) {
                    const e = Object.keys(t).filter((e => o.indexOf(e) > -1))[0];
                    n.setAttribute("href", t[e])
                }
            }

            #f(e) {
                const {skipLink: n} = e.config;
                if (!1 !== n) {
                    const i = s("#skip-to-content");
                    let t = "string" == typeof n ? n : "Skip to main content";
                    if (n?.constructor === Object) {
                        const i = Object.keys(n).find((n => e.route.path.startsWith(n.startsWith("/") ? n : `/${n}`)));
                        t = i && n[i] || t
                    }
                    if (i) i.innerHTML = t; else {
                        const e = `<button id="skip-to-content">${t}</button>`;
                        d.insertAdjacentHTML("afterbegin", e)
                    }
                }
            }

            _renderTo(e, n, i) {
                const t = s(e);
                t && (t[i ? "outerHTML" : "innerHTML"] = n)
            }

            _renderSidebar(e) {
                const {maxLevel: n, subMaxLevel: i, loadSidebar: t, hideSidebar: o} = this.config,
                    a = s("aside.sidebar"), r = s("button.sidebar-toggle");
                if (o) return d.classList.add("hidesidebar"), a?.remove(a), r?.remove(r), null;
                this._renderTo(".sidebar-nav", this.compiler.sidebar(e, n)), r.setAttribute("aria-expanded", !c);
                const u = g(`.sidebar-nav a[href="${this.router.toURL(this.route.path)}"]`);
                this.#c(".sidebar-nav a"), t && u ? u.parentNode.innerHTML += this.compiler.subSidebar(i) || "" : this.compiler.subSidebar(), this._bindEventOnRendered(u)
            }

            _bindEventOnRendered(e) {
                const {autoHeader: n} = this.config;
                if (this.onRender(), n && e) {
                    const n = s("#main"), i = n.children[0];
                    if (i && "H1" !== i.tagName) {
                        h(n, l("div", this.compiler.header(e.innerText, 1)).children[0])
                    }
                }
            }

            _renderNav(e) {
                e && this._renderTo("nav", this.compiler.compile(e)), this.#c("nav a")
            }

            _renderMain(e) {
                let n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {},
                    i = arguments.length > 2 ? arguments[2] : void 0;
                const {response: t} = this.route;
                !t || t.ok || e && 404 === t.status || (e = `# ${t.status} - ${t.statusText}`), this.callHook("beforeEach", e, (t => {
                    let o;
                    const a = () => {
                        n.updatedAt && (o = this.#s(o, n.updatedAt, this.config.formatUpdated)), this.callHook("afterEach", o, (e => {
                            this.#u(e), i()
                        }))
                    };
                    this.isHTML ? (o = this.result = e, a()) : Cn({compiler: this.compiler, raw: t}, (e => {
                        o = this.compiler.compile(e), a()
                    }))
                }))
            }

            _renderCover(e, n) {
                const i = s(".cover");
                if (m(s("main"), n ? "add" : "remove", "hidden"), !e) return void m(i, "remove", "show");
                m(i, "add", "show");
                let t = this.coverIsHTML ? e : this.compiler.cover(e);
                const o = t.trim().match('<p><img.*?data-origin="(.*?)"[^a]+alt="(.*?)">([^<]*?)</p>$');
                if (o) {
                    if ("color" === o[2]) i.style.background = o[1] + (o[3] || ""); else {
                        let e = o[1];
                        m(i, "add", "has-mask"), A(o[1]) || (e = L(this.router.getBasePath(), o[1])), i.style.backgroundImage = `url(${e})`, i.style.backgroundSize = "cover", i.style.backgroundPosition = "center center"
                    }
                    t = t.replace(o[0], "")
                }
                this._renderTo(".cover-main", t)
            }

            _updateRender() {
                this.#d(this), this.#f(this)
            }

            initRender() {
                const e = this.config;
                this.compiler = new gn(e, this.router), window.__current_docsify_compiler__ = this.compiler;
                const n = g(e.el || "#app");
                if (n) {
                    let o = "";
                    if (e.repo && (o += (i = e.repo, t = e.cornerExternalLinkTarget, i ? (/\/\//.test(i) || (i = "https://github.com/" + i), `\n    <a href="${i = i.replace(/^git\+/, "")}" target="${t = t || "_blank"}" class="github-corner" aria-label="View source on Github">\n      <svg viewBox="0 0 250 250" aria-hidden="true">\n        <path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path>\n        <path d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2" fill="currentColor" style="transform-origin: 130px 106px;" class="octo-arm"></path>\n        <path d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z" fill="currentColor" class="octo-body"></path>\n      </svg>\n    </a>\n  `) : "")), e.coverpage && (o += function () {
                        const e = ", 100%, 85%";
                        return `\n    <section class="cover show" role="complementary" aria-label="cover" style="background: \n    linear-gradient(\n      to left bottom,\n      hsl(${Math.floor(255 * Math.random()) + e}) 0%,\n      hsl(${Math.floor(255 * Math.random()) + e}) 100%\n    )\n  ">\n      <div class="mask"></div>\n      <div class="cover-main">\x3c!--cover--\x3e</div>\n    </section>\n  `
                    }()), e.logo) {
                        const n = /^data:image/.test(e.logo), i = /(?:http[s]?:)?\/\//.test(e.logo),
                            t = /^\./.test(e.logo);
                        n || i || t || (e.logo = L(this.router.getBasePath(), e.logo))
                    }
                    o += function (e) {
                        const n = e.name ? e.name : "";
                        return `\n    <main role="presentation">\n    <button class="sidebar-toggle" title="Press \\ to toggle" aria-label="Toggle primary navigation" aria-keyshortcuts="\\" aria-controls="__sidebar">\n      <div class="sidebar-toggle-button" aria-hidden="true">\n        <span></span><span></span><span></span>\n      </div>\n    </button>\n    <aside id="__sidebar" class="sidebar" role="none">\n      ${e.name ? `\n            <h1 class="app-name"><a class="app-name-link" data-nosearch>${e.logo ? `<img alt="${n}" src=${e.logo} />` : n}</a></h1>\n          ` : ""}\n      <div class="sidebar-nav" role="navigation" aria-label="primary">\x3c!--sidebar--\x3e</div>\n    </aside>\n  \n      <section class="content">\n        <article id="main" class="markdown-section" role="main" tabindex="-1">\x3c!--main--\x3e</article>\n      </section>\n    </main>\n  `
                    }(e), this._renderTo(n, o, !0)
                } else this.rendered = !0;
                var i, t, o;
                if (e.loadNavbar) {
                    const n = g("nav") || l("nav"), i = e.mergeNavbar && c;
                    n.setAttribute("aria-label", "secondary"), i ? g(".sidebar").prepend(n) : (d.prepend(n), n.classList.add("app-nav"), n.classList.toggle("no-badge", !e.repo))
                }
                e.themeColor && u.head.appendChild(l("div", (o = e.themeColor, `<style>:root{--theme-color: ${o};}</style>`)).firstElementChild), this._updateRender(), m(d, "ready")
            }
        }
    }

    function On(e) {
        return class extends e {
            routes() {
                return this.config.routes || {}
            }

            matchVirtualRoute(e) {
                const n = this.routes(), i = Object.keys(n);
                let t = () => null;

                function o() {
                    const a = i.shift();
                    if (!a) return t(null);
                    const c = function (e) {
                        const n = e.startsWith("^") ? e : `^${e}`;
                        return n.endsWith("$") ? n : `${n}$`
                    }(a), r = e.match(c);
                    if (!r) return o();
                    const s = n[a];
                    if ("string" == typeof s) {
                        return t(s)
                    }
                    if ("function" == typeof s) {
                        const n = s, [i, a] = function () {
                            let e = () => null;
                            return [function (n) {
                                e(n)
                            }, function (n) {
                                e = n
                            }]
                        }();
                        if (a((e => "string" == typeof e ? t(e) : !1 === e ? t(null) : o())), n.length <= 2) {
                            return i(n(e, r))
                        }
                        return n(e, r, i)
                    }
                    return o()
                }

                return {
                    then(e) {
                        t = e, o()
                    }
                }
            }
        }
    }

    const qn = document.currentScript;

    class Pn extends (function (e) {
        return class extends e {
            #g(e, n, i, t, o, a) {
                e = a ? e : e.replace(/\/$/, ""), (e = E(e)) && Ln(o.router.getFile(e + i) + n, !1, o.config.requestHeaders).then(t, (a => this.#g(e, n, i, t, o)))
            }

            #p;
            #l = () => this.#p && this.#p.abort && this.#p.abort();
            #v = (e, n) => (this.#l(), this.#p = Ln(e, !0, n), this.#p);
            #h = (e, n) => {
                const {notFoundPage: i, ext: t} = n, o = "_404" + (t || ".md");
                let a, c;
                switch (typeof i) {
                    case"boolean":
                        c = o;
                        break;
                    case"string":
                        c = i;
                        break;
                    case"object":
                        a = Object.keys(i).sort(((e, n) => n.length - e.length)).filter((n => e.match(new RegExp("^" + n))))[0], c = a && i[a] || o
                }
                return c
            };

            _loadSideAndNav(e, n, i, t) {
                return () => {
                    if (!i) return t();
                    this.#g(e, n, i, (e => {
                        this._renderSidebar(e), t()
                    }), this, !0)
                }
            }

            _fetch() {
                let e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : t;
                const {query: n} = this.route, {path: i} = this.route;
                if (a(i)) history.replaceState(null, "", "#"), this.router.normalize(); else {
                    const t = $(n, ["id"]), {loadNavbar: o, requestHeaders: c, loadSidebar: r} = this.config,
                        s = this.router.getFile(i);
                    this.isRemoteUrl = a(s), this.isHTML = /\.html$/g.test(s);
                    const u = (n, o, a) => {
                        this.route.response = a, this._renderMain(n, o, this._loadSideAndNav(i, t, r, e))
                    }, d = (n, o) => {
                        this.route.response = o, this._fetchFallbackPage(i, t, e) || this._fetch404(s, t, e)
                    };
                    this.isRemoteUrl ? this.#v(s + t, c).then(u, d) : this.matchVirtualRoute(i).then((e => {
                        "string" == typeof e ? u(e) : this.#v(s + t, c).then(u, d)
                    })), o && this.#g(i, t, o, (e => this._renderNav(e)), this, !0)
                }
            }

            _fetchCover() {
                const {coverpage: e, requestHeaders: n} = this.config, i = this.route.query, t = E(this.route.path);
                if (e) {
                    let o = null;
                    const a = this.route.path;
                    if ("string" == typeof e) "/" === a && (o = e); else if (Array.isArray(e)) o = e.indexOf(a) > -1 && "_coverpage"; else {
                        const n = e[a];
                        o = !0 === n ? "_coverpage" : n
                    }
                    const c = Boolean(o) && this.config.onlyCover;
                    return o ? (o = this.router.getFile(t + o), this.coverIsHTML = /\.html$/g.test(o), Ln(o + $(i, ["id"]), !1, n).then((e => this._renderCover(e, c)))) : this._renderCover(null, c), c
                }
            }

            $fetch() {
                let e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : t,
                    n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : this.onNavigate.bind(this);
                const i = () => {
                    this.callHook("doneEach"), e()
                };
                this._fetchCover() ? i() : this._fetch((() => {
                    n(), i()
                }))
            }

            _fetchFallbackPage(e, n) {
                let i = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : t;
                const {requestHeaders: o, fallbackLanguages: a, loadSidebar: c} = this.config;
                if (!a) return !1;
                const r = e.split("/")[1];
                if (-1 === a.indexOf(r)) return !1;
                const s = this.router.getFile(e.replace(new RegExp(`^/${r}`), ""));
                return this.#v(s + n, o).then(((t, o) => this._renderMain(t, o, this._loadSideAndNav(e, n, c, i))), (t => this._fetch404(e, n, i))), !0
            }

            _fetch404(e, n) {
                let i = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : t;
                const {loadSidebar: o, requestHeaders: a, notFoundPage: c} = this.config,
                    r = this._loadSideAndNav(e, n, o, i);
                if (c) {
                    const n = this.#h(e, this.config);
                    return this.#v(this.router.getFile(n), a).then(((e, n) => this._renderMain(e, n, r)), (e => this._renderMain(null, {}, r))), !0
                }
                return this._renderMain(null, {}, r), !1
            }

            initFetch() {
                this.config, this.$fetch((e => this.callHook("ready")))
            }
        }
    }(function (e) {
        return class extends e {
            #_;
            #m;
            #b = u.title;

            initEvent() {
                const {topMargin: e} = this.config;
                e && document.documentElement.style.setProperty("scroll-padding-top", `${e}px`), this.#k(), this.#w("#skip-to-content"), this.#y(".sidebar"), this.#x("button.sidebar-toggle"), this.#$()
            }

            #k() {
                const e = g("section.cover");
                if (!e) return void m(d, "add", "sticky");
                new IntersectionObserver((e => {
                    const n = e[0].isIntersecting;
                    m(d, n ? "remove" : "add", "sticky")
                })).observe(e)
            }

            #A() {
                const e = p("#main :where(h1, h2, h3, h4, h5)"), n = new Set;
                this.#_?.disconnect(), this.#_ = new IntersectionObserver((e => {
                    if (this.#m) return;
                    for (const i of e) {
                        const e = i.isIntersecting ? "add" : "delete";
                        n[e](i.target)
                    }
                    const i = n.size > 1 ? Array.from(n).sort(((e, n) => e.compareDocumentPosition(n) & Node.DOCUMENT_POSITION_FOLLOWING ? -1 : 1))[0] : n.values().next().value;
                    if (i) {
                        const e = i.getAttribute("id"), n = this.router.toURL(this.router.getCurrentPath(), {id: e}),
                            t = this.#S(n);
                        t?.scrollIntoView({behavior: "instant", block: "nearest", inline: "nearest"})
                    }
                }), {rootMargin: "0% 0% -50% 0%"}), e.forEach((e => {
                    this.#_.observe(e)
                }))
            }

            #$() {
                const {keyBindings: e} = this.config, n = ["alt", "ctrl", "meta", "shift"];
                e && e.constructor === Object && (Object.values(e || []).forEach((e => {
                    const {bindings: i} = e;
                    i && (e.bindings = Array.isArray(i) ? i : [i], e.bindings = e.bindings.map((e => {
                        const i = [[], []];
                        return "string" == typeof e && (e = e.split("+")), e.forEach((e => {
                            const t = n.includes(e), o = i[t ? 0 : 1], a = e.trim().toLowerCase();
                            o.push(a)
                        })), i.forEach((e => e.sort())), i.flat()
                    })))
                })), _("keydown", (i => {
                    if (document.activeElement.matches("input, select, textarea")) return;
                    Object.values(e || []).filter((e => {
                        let {bindings: t} = e;
                        return t && t.some((e => e.every((e => n.includes(e) && i[e + "Key"] || i.key === e || i.code.toLowerCase() === e || i.code.toLowerCase() === `key${e}`))))
                    })).forEach((e => {
                        let {callback: n} = e;
                        i.preventDefault(), n(i)
                    }))
                })))
            }

            #y(e) {
                (e = "string" == typeof e ? document.querySelector(e) : e) && _(e, "click", (e => {
                    let {target: n} = e;
                    "A" === n.nodeName && n.nextSibling && n.nextSibling.classList && n.nextSibling.classList.contains("app-sub-sidebar") && m(n.parentNode, "collapse")
                }))
            }

            #x(e) {
                if (!(e = "string" == typeof e ? document.querySelector(e) : e)) return;
                const n = () => {
                    d.classList.toggle("close");
                    const n = c ? d.classList.contains("close") : !d.classList.contains("close");
                    e.setAttribute("aria-expanded", n)
                };
                _(e, "click", (e => {
                    e.stopPropagation(), n()
                })), c && _(d, "click", (() => d.classList.contains("close") && n()))
            }

            #w(e) {
                (e = "string" == typeof e ? document.querySelector(e) : e) && e.addEventListener("click", (e => {
                    e.preventDefault(), g("main")?.scrollIntoView({behavior: "smooth"}), this.#E()
                }))
            }

            onRender() {
                const e = this.router.toURL(this.router.getCurrentPath()), n = g(`.sidebar a[href='${e}']`)?.innerText;
                u.title = n || this.#b, this.#z(), this.#T(), this.#A()
            }

            onNavigate(e) {
                const {auto2top: n, topMargin: i} = this.config, {query: t} = this.route;
                if (this.#S(), "history" !== e) if (t.id) {
                    const e = g(`.markdown-section :where(h1, h2, h3, h4, h5)[id="${t.id}"]`);
                    e && (this.#R(), e.scrollIntoView({behavior: "smooth", block: "start"}))
                } else "navigate" === e && n && (document.scrollingElement.scrollTop = i ?? 0);
                (t.id || "navigate" === e) && this.#E()
            }

            #E() {
                const e = {preventScroll: !0, ...arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {}}, {query: n} = this.route,
                    i = n.id ? g(`#${n.id}`) : g("#main :where(h1, h2, h3, h4, h5, h6)") || g("#main");
                i?.focus(e)
            }

            #z() {
                const e = decodeURIComponent(this.router.toURL(this.route.path)), n = g("nav.app-nav");
                if (!n) return;
                const i = p(n, "a").sort(((e, n) => n.href.length - e.href.length)).find((n => e.includes(n.getAttribute("href")) || e.includes(decodeURI(n.getAttribute("href"))))),
                    t = g(n, "li.active");
                return i && i !== t && (t?.classList.remove("active"), i.classList.add("active")), i
            }

            #S(e) {
                e ??= this.router.toURL(this.router.getCurrentPath());
                const n = g(".sidebar");
                if (!n) return;
                const i = g(n, "li.active"),
                    t = g(n, `a[href="${e}"], a[href="${decodeURIComponent(e)}"]`)?.closest("li");
                return t && t !== i && (i?.classList.remove("active"), t.classList.add("active")), t
            }

            #T(e) {
                e ??= this.router.toURL(this.route.path);
                const n = g(".sidebar");
                if (!n) return;
                const i = e?.split("?")[0], t = g(n, "li[aria-current]"),
                    o = g(n, `a[href="${i}"], a[href="${decodeURIComponent(i)}"]`)?.closest("li");
                return o && o !== t && (t?.removeAttribute("aria-current"), o.setAttribute("aria-current", "page")), o
            }

            #R() {
                document.addEventListener("scroll", (() => {
                    if (this.#m = !0, "onscrollend" in window) document.addEventListener("scrollend", (() => this.#m = !1), {once: !0}); else {
                        const e = () => {
                            clearTimeout(n), n = setTimeout((() => {
                                document.removeEventListener("scroll", e), this.#m = !1
                            }), 100)
                        };
                        let n;
                        document.addEventListener("scroll", e, !1)
                    }
                }), {once: !0})
            }
        }
    }(jn(On(function (e) {
        return class extends e {
            route = {};

            updateRender() {
                this.router.normalize(), this.route = this.router.parse(), d.setAttribute("data-page", this.route.file)
            }

            initRouter() {
                const e = this.config;
                let n;
                n = "history" === (e.routerMode || "hash") ? new q(e) : new O(e), this.router = n, this.updateRender(), P = this.route, n.onchange((e => {
                    this.updateRender(), this._updateRender(), P.path !== this.route.path ? (this.$fetch(t, this.onNavigate.bind(this, e.source)), P = this.route) : this.onNavigate(e.source)
                }))
            }
        }
    }(function (e) {
        return class extends e {
            _hooks = {};
            _lifecycle = {};

            initLifecycle() {
                ["init", "mounted", "beforeEach", "afterEach", "doneEach", "ready"].forEach((e => {
                    const n = this._hooks[e] = [];
                    this._lifecycle[e] = e => n.push(e)
                }))
            }

            callHook(e, n) {
                let i = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : t;
                const o = this._hooks[e], a = this.config.catchPluginErrors, c = function (e) {
                    const t = o[e];
                    if (e >= o.length) i(n); else if ("function" == typeof t) {
                        const i = "Docsify plugin error";
                        if (2 === t.length) try {
                            t(n, (i => {
                                n = i, c(e + 1)
                            }))
                        } catch (n) {
                            if (!a) throw n;
                            console.error(i, n), c(e + 1)
                        } else try {
                            const i = t(n);
                            n = void 0 === i ? n : i, c(e + 1)
                        } catch (n) {
                            if (!a) throw n;
                            console.error(i, n), c(e + 1)
                        }
                    } else c(e + 1)
                };
                c(0)
            }
        }
    }(Object))))))) {
        config = function (e) {
            const t = Object.assign({
                auto2top: !1,
                autoHeader: !1,
                basePath: "",
                catchPluginErrors: !0,
                cornerExternalLinkTarget: "_blank",
                coverpage: "",
                el: "#app",
                executeScript: null,
                ext: ".md",
                externalLinkRel: "noopener",
                externalLinkTarget: "_blank",
                formatUpdated: "",
                ga: "",
                homepage: "README.md",
                loadNavbar: null,
                loadSidebar: null,
                maxLevel: 6,
                mergeNavbar: !1,
                name: "",
                nameLink: window.location.pathname,
                nativeEmoji: !1,
                noCompileLinks: [],
                noEmoji: !1,
                notFoundPage: !1,
                plugins: [],
                relativePath: !1,
                repo: "",
                routes: {},
                routerMode: "hash",
                subMaxLevel: 0,
                topMargin: 0,
                __themeColor: "",
                get themeColor() {
                    return this.__themeColor
                },
                set themeColor(e) {
                    e && (this.__themeColor = e, console.warn(Sn("\n              $docsify.themeColor is deprecated. Use a --theme-color property in your style sheet. Example:\n              <style>\n                :root {\n                  --theme-color: deeppink;\n                }\n              </style>\n            ").trim()))
                }
            }, "function" == typeof window.$docsify ? window.$docsify(e) : window.$docsify);
            !1 !== t.keyBindings && (t.keyBindings = Object.assign({
                toggleSidebar: {
                    bindings: ["\\"], callback(e) {
                        const n = document.querySelector(".sidebar-toggle");
                        n && (n.click(), n.focus())
                    }
                }
            }, t.keyBindings));
            const o = qn || Array.from(document.getElementsByTagName("script")).filter((e => /docsify\./.test(e.src)))[0];
            if (o) for (const e of Object.keys(t)) {
                const a = o.getAttribute("data-" + n(e));
                i(a) && (t[e] = "" === a || a)
            }
            return !0 === t.loadSidebar && (t.loadSidebar = "_sidebar" + t.ext), !0 === t.loadNavbar && (t.loadNavbar = "_navbar" + t.ext), !0 === t.coverpage && (t.coverpage = "_coverpage" + t.ext), !0 === t.repo && (t.repo = ""), !0 === t.name && (t.name = ""), window.$docsify = t, t
        }(this);

        constructor() {
            super(), this.initLifecycle(), this.initPlugin(), this.callHook("init"), this.initRouter(), this.initRender(), this.initEvent(), this.initFetch(), this.callHook("mounted")
        }

        initPlugin() {
            this.config.plugins.forEach((e => {
                try {
                    o(e) && e(this._lifecycle, this)
                } catch (e) {
                    if (!this.config.catchPluginErrors) throw e;
                    {
                        const n = "Docsify plugin error";
                        console.error(n, e)
                    }
                }
            }))
        }
    }

    var Mn = Object.freeze({
        __proto__: null,
        cached: e,
        cleanPath: z,
        getParentPath: E,
        getPath: L,
        hyphenate: n,
        inBrowser: !0,
        isAbsolutePath: A,
        isExternal: a,
        isFn: o,
        isMobile: c,
        isPrimitive: i,
        noop: t,
        parseQuery: x,
        removeParams: S,
        replaceSlug: F,
        resolvePath: T,
        stringifyQuery: $
    });
    window.Docsify = {
        util: Mn,
        dom: k,
        get: Ln,
        slugify: nn,
        version: "4.13.0"
    }, window.DocsifyCompiler = gn, window.marked = We, window.Prism = un, b((() => new Pn))
}();
//# sourceMappingURL=docsify.min.js.map

(function() {
    "use strict";
    function cached$1(fn) {
        const cache = Object.create(null);
        return function(str) {
            const key = isPrimitive(str) ? str : JSON.stringify(str);
            const hit = cache[key];
            return hit || (cache[key] = fn(str));
        };
    }
    const hyphenate = cached$1((str => str.replace(/([A-Z])/g, (m => "-" + m.toLowerCase()))));
    function isPrimitive(value) {
        return typeof value === "string" || typeof value === "number";
    }
    function noop() {}
    function isFn(obj) {
        return typeof obj === "function";
    }
    function isExternal(url) {
        const match = url.match(/^([^:/?#]+:)?(?:\/{2,}([^/?#]*))?([^?#]+)?(\?[^#]*)?(#.*)?/);
        if (typeof match[1] === "string" && match[1].length > 0 && match[1].toLowerCase() !== location.protocol) {
            return true;
        }
        if (typeof match[2] === "string" && match[2].length > 0 && match[2].replace(new RegExp(":(" + {
            "http:": 80,
            "https:": 443
        }[location.protocol] + ")?$"), "") !== location.host) {
            return true;
        }
        if (/^\/\\/.test(url)) {
            return true;
        }
        return false;
    }
    const inBrowser = true;
    const isMobile = document.body.clientWidth <= 600;
    const cacheNode = {};
    function getNode(el) {
        let noCache = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : false;
        if (typeof el === "string") {
            if (typeof window.Vue !== "undefined") {
                return find(el);
            }
            el = noCache ? find(el) : cacheNode[el] || (cacheNode[el] = find(el));
        }
        return el;
    }
    const $ = document;
    const body = $.body;
    const head = $.head;
    function find(el, node) {
        return node ? el.querySelector(node) : $.querySelector(el);
    }
    function findAll(el, node) {
        return Array.from(node ? el.querySelectorAll(node) : $.querySelectorAll(el));
    }
    function create(node, tpl) {
        node = $.createElement(node);
        if (tpl) {
            node.innerHTML = tpl;
        }
        return node;
    }
    function appendTo(target, el) {
        return target.appendChild(el);
    }
    function before(target, el) {
        return target.insertBefore(el, target.children[0]);
    }
    function on(el, type, handler) {
        isFn(type) ? window.addEventListener(el, type) : el.addEventListener(type, handler);
    }
    function off(el, type, handler) {
        isFn(type) ? window.removeEventListener(el, type) : el.removeEventListener(type, handler);
    }
    function toggleClass(el, type, val) {
        el && el.classList[val ? type : "toggle"](val || type);
    }
    function style(content) {
        appendTo(head, create("style", content));
    }
    function documentReady(callback) {
        let doc = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : document;
        const state = doc.readyState;
        if (state === "complete" || state === "interactive") {
            return setTimeout(callback, 0);
        }
        doc.addEventListener("DOMContentLoaded", callback);
    }
    var dom = Object.freeze({
        __proto__: null,
        $: $,
        appendTo: appendTo,
        before: before,
        body: body,
        create: create,
        documentReady: documentReady,
        find: find,
        findAll: findAll,
        getNode: getNode,
        head: head,
        off: off,
        on: on,
        style: style,
        toggleClass: toggleClass
    });
    const decode = decodeURIComponent;
    const encode = encodeURIComponent;
    function parseQuery(query) {
        const res = {};
        query = query.trim().replace(/^(\?|#|&)/, "");
        if (!query) {
            return res;
        }
        query.split("&").forEach((param => {
            const parts = param.replace(/\+/g, " ").split("=");
            res[parts[0]] = parts[1] && decode(parts[1]);
        }));
        return res;
    }
    function stringifyQuery(obj) {
        let ignores = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : [];
        const qs = [];
        for (const key in obj) {
            if (ignores.indexOf(key) > -1) {
                continue;
            }
            qs.push(obj[key] ? `${encode(key)}=${encode(obj[key])}`.toLowerCase() : encode(key));
        }
        return qs.length ? `?${qs.join("&")}` : "";
    }
    const isAbsolutePath = cached$1((path => /(:|(\/{2}))/g.test(path)));
    const removeParams = cached$1((path => path.split(/[?#]/)[0]));
    const getParentPath = cached$1((path => {
        if (/\/$/g.test(path)) {
            return path;
        }
        const matchingParts = path.match(/(\S*\/)[^/]+$/);
        return matchingParts ? matchingParts[1] : "";
    }));
    const cleanPath = cached$1((path => path.replace(/^\/+/, "/").replace(/([^:])\/{2,}/g, "$1/")));
    const resolvePath = cached$1((path => {
        const segments = path.replace(/^\//, "").split("/");
        const resolved = [];
        for (const segment of segments) {
            if (segment === "..") {
                resolved.pop();
            } else if (segment !== ".") {
                resolved.push(segment);
            }
        }
        return "/" + resolved.join("/");
    }));
    function normaliseFragment(path) {
        return path.split("/").filter((p => p.indexOf("#") === -1)).join("/");
    }
    function getPath() {
        for (var _len = arguments.length, args = new Array(_len), _key = 0; _key < _len; _key++) {
            args[_key] = arguments[_key];
        }
        return cleanPath(args.map(normaliseFragment).join("/"));
    }
    const replaceSlug = cached$1((path => path.replace("#", "?id=")));
    class History {
        #cached={};
        constructor(config) {
            this.config = config;
        }
        #getAlias(path, alias, last) {
            const match = Object.keys(alias).filter((key => {
                const re = this.#cached[key] || (this.#cached[key] = new RegExp(`^${key}$`));
                return re.test(path) && path !== last;
            }))[0];
            return match ? this.#getAlias(path.replace(this.#cached[match], alias[match]), alias, path) : path;
        }
        #getFileName(path, ext) {
            return new RegExp(`\\.(${ext.replace(/^\./, "")}|html)$`, "g").test(path) ? path : /\/$/g.test(path) ? `${path}README${ext}` : `${path}${ext}`;
        }
        getBasePath() {
            return this.config.basePath;
        }
        getFile() {
            let path = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : this.getCurrentPath();
            let isRelative = arguments.length > 1 ? arguments[1] : undefined;
            const {config: config} = this;
            const base = this.getBasePath();
            const ext = typeof config.ext === "string" ? config.ext : ".md";
            path = config.alias ? this.#getAlias(path, config.alias) : path;
            path = this.#getFileName(path, ext);
            path = path === `/README${ext}` ? config.homepage || path : path;
            path = isAbsolutePath(path) ? path : getPath(base, path);
            if (isRelative) {
                path = path.replace(new RegExp(`^${base}`), "");
            }
            return path;
        }
        onchange() {
            let cb = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : noop;
            cb();
        }
        getCurrentPath() {}
        normalize() {}
        parse() {}
        toURL(path, params, currentRoute) {
            const local = currentRoute && path[0] === "#";
            const route = this.parse(replaceSlug(path));
            route.query = {
                ...route.query,
                ...params
            };
            path = route.path + stringifyQuery(route.query);
            path = path.replace(/\.md(\?)|\.md$/, "$1");
            if (local) {
                const idIndex = currentRoute.indexOf("?");
                path = (idIndex > 0 ? currentRoute.substring(0, idIndex) : currentRoute) + path;
            }
            if (this.config.relativePath && path.indexOf("/") !== 0) {
                const currentDir = currentRoute.substring(0, currentRoute.lastIndexOf("/") + 1);
                return cleanPath(resolvePath(currentDir + path));
            }
            return cleanPath("/" + path);
        }
    }
    function replaceHash(path) {
        const i = location.href.indexOf("#");
        location.replace(location.href.slice(0, i >= 0 ? i : 0) + "#" + path);
    }
    class HashHistory extends History {
        mode="hash";
        getBasePath() {
            const path = window.location.pathname || "";
            const base = this.config.basePath;
            const basePath = path.endsWith(".html") ? path + "#/" + base : path + "/" + base;
            return /^(\/|https?:)/g.test(base) ? base : cleanPath(basePath);
        }
        getCurrentPath() {
            const href = location.href;
            const index = href.indexOf("#");
            return index === -1 ? "" : href.slice(index + 1);
        }
        onchange() {
            let cb = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : noop;
            let navigating = false;
            on("click", (e => {
                const el = e.target.tagName === "A" ? e.target : e.target.parentNode;
                if (el && el.tagName === "A" && !isExternal(el.href)) {
                    navigating = true;
                }
            }));
            on("hashchange", (e => {
                const source = navigating ? "navigate" : "history";
                navigating = false;
                cb({
                    event: e,
                    source: source
                });
            }));
        }
        normalize() {
            let path = this.getCurrentPath();
            path = replaceSlug(path);
            if (path.charAt(0) === "/") {
                return replaceHash(path);
            }
            replaceHash("/" + path);
        }
        parse() {
            let path = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : location.href;
            let query = "";
            const hashIndex = path.indexOf("#");
            if (hashIndex >= 0) {
                path = path.slice(hashIndex + 1);
            }
            const queryIndex = path.indexOf("?");
            if (queryIndex >= 0) {
                query = path.slice(queryIndex + 1);
                path = path.slice(0, queryIndex);
            }
            return {
                path: path,
                file: this.getFile(path, true),
                query: parseQuery(query),
                response: {}
            };
        }
        toURL(path, params, currentRoute) {
            return "#" + super.toURL(path, params, currentRoute);
        }
    }
    class HTML5History extends History {
        mode="history";
        getCurrentPath() {
            const base = this.getBasePath();
            let path = window.location.pathname;
            if (base && path.indexOf(base) === 0) {
                path = path.slice(base.length);
            }
            return (path || "/") + window.location.search + window.location.hash;
        }
        onchange() {
            let cb = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : noop;
            on("click", (e => {
                const el = e.target.tagName === "A" ? e.target : e.target.parentNode;
                if (el && el.tagName === "A" && !isExternal(el.href)) {
                    e.preventDefault();
                    const url = el.href;
                    window.history.pushState({
                        key: url
                    }, "", url);
                    cb({
                        event: e,
                        source: "navigate"
                    });
                }
            }));
            on("popstate", (e => {
                cb({
                    event: e,
                    source: "history"
                });
            }));
        }
        parse() {
            let path = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : location.href;
            let query = "";
            const queryIndex = path.indexOf("?");
            if (queryIndex >= 0) {
                query = path.slice(queryIndex + 1);
                path = path.slice(0, queryIndex);
            }
            const base = getPath(location.origin);
            const baseIndex = path.indexOf(base);
            if (baseIndex > -1) {
                path = path.slice(baseIndex + base.length);
            }
            return {
                path: path,
                file: this.getFile(path),
                query: parseQuery(query),
                response: {}
            };
        }
    }
    let lastRoute = {};
    function Router(Base) {
        return class Router extends Base {
            route={};
            updateRender() {
                this.router.normalize();
                this.route = this.router.parse();
                body.setAttribute("data-page", this.route.file);
            }
            initRouter() {
                const config = this.config;
                const mode = config.routerMode || "hash";
                let router;
                if (mode === "history") {
                    router = new HTML5History(config);
                } else {
                    router = new HashHistory(config);
                }
                this.router = router;
                this.updateRender();
                lastRoute = this.route;
                router.onchange((params => {
                    this.updateRender();
                    this._updateRender();
                    if (lastRoute.path === this.route.path) {
                        this.onNavigate(params.source);
                        return;
                    }
                    this.$fetch(noop, this.onNavigate.bind(this, params.source));
                    lastRoute = this.route;
                }));
            }
        };
    }
    var RGX = /([^{]*?)\w(?=\})/g;
    var MAP = {
        YYYY: "getFullYear",
        YY: "getYear",
        MM: function(d) {
            return d.getMonth() + 1;
        },
        DD: "getDate",
        HH: "getHours",
        mm: "getMinutes",
        ss: "getSeconds",
        fff: "getMilliseconds"
    };
    function tinydate(str, custom) {
        var parts = [], offset = 0;
        str.replace(RGX, (function(key, _, idx) {
            parts.push(str.substring(offset, idx - 1));
            offset = idx += key.length + 1;
            parts.push((function(d) {
                return ("00" + (typeof MAP[key] === "string" ? d[MAP[key]]() : MAP[key](d))).slice(-key.length);
            }));
        }));
        if (offset !== str.length) {
            parts.push(str.substring(offset));
        }
        return function(arg) {
            var out = "", i = 0, d = arg || new Date;
            for (;i < parts.length; i++) {
                out += typeof parts[i] === "string" ? parts[i] : parts[i](d);
            }
            return out;
        };
    }
    function _getDefaults() {
        return {
            async: false,
            breaks: false,
            extensions: null,
            gfm: true,
            hooks: null,
            pedantic: false,
            renderer: null,
            silent: false,
            tokenizer: null,
            walkTokens: null
        };
    }
    let _defaults = _getDefaults();
    function changeDefaults(newDefaults) {
        _defaults = newDefaults;
    }
    const escapeTest = /[&<>"']/;
    const escapeReplace = new RegExp(escapeTest.source, "g");
    const escapeTestNoEncode = /[<>"']|&(?!(#\d{1,7}|#[Xx][a-fA-F0-9]{1,6}|\w+);)/;
    const escapeReplaceNoEncode = new RegExp(escapeTestNoEncode.source, "g");
    const escapeReplacements = {
        "&": "&amp;",
        "<": "&lt;",
        ">": "&gt;",
        '"': "&quot;",
        "'": "&#39;"
    };
    const getEscapeReplacement = ch => escapeReplacements[ch];
    function escape$1(html, encode) {
        if (encode) {
            if (escapeTest.test(html)) {
                return html.replace(escapeReplace, getEscapeReplacement);
            }
        } else {
            if (escapeTestNoEncode.test(html)) {
                return html.replace(escapeReplaceNoEncode, getEscapeReplacement);
            }
        }
        return html;
    }
    const unescapeTest = /&(#(?:\d+)|(?:#x[0-9A-Fa-f]+)|(?:\w+));?/gi;
    function unescape(html) {
        return html.replace(unescapeTest, ((_, n) => {
            n = n.toLowerCase();
            if (n === "colon") return ":";
            if (n.charAt(0) === "#") {
                return n.charAt(1) === "x" ? String.fromCharCode(parseInt(n.substring(2), 16)) : String.fromCharCode(+n.substring(1));
            }
            return "";
        }));
    }
    const caret = /(^|[^\[])\^/g;
    function edit(regex, opt) {
        let source = typeof regex === "string" ? regex : regex.source;
        opt = opt || "";
        const obj = {
            replace: (name, val) => {
                let valSource = typeof val === "string" ? val : val.source;
                valSource = valSource.replace(caret, "$1");
                source = source.replace(name, valSource);
                return obj;
            },
            getRegex: () => new RegExp(source, opt)
        };
        return obj;
    }
    function cleanUrl(href) {
        try {
            href = encodeURI(href).replace(/%25/g, "%");
        } catch (e) {
            return null;
        }
        return href;
    }
    const noopTest = {
        exec: () => null
    };
    function splitCells(tableRow, count) {
        const row = tableRow.replace(/\|/g, ((match, offset, str) => {
            let escaped = false;
            let curr = offset;
            while (--curr >= 0 && str[curr] === "\\") escaped = !escaped;
            if (escaped) {
                return "|";
            } else {
                return " |";
            }
        })), cells = row.split(/ \|/);
        let i = 0;
        if (!cells[0].trim()) {
            cells.shift();
        }
        if (cells.length > 0 && !cells[cells.length - 1].trim()) {
            cells.pop();
        }
        if (count) {
            if (cells.length > count) {
                cells.splice(count);
            } else {
                while (cells.length < count) cells.push("");
            }
        }
        for (;i < cells.length; i++) {
            cells[i] = cells[i].trim().replace(/\\\|/g, "|");
        }
        return cells;
    }
    function rtrim(str, c, invert) {
        const l = str.length;
        if (l === 0) {
            return "";
        }
        let suffLen = 0;
        while (suffLen < l) {
            const currChar = str.charAt(l - suffLen - 1);
            if (currChar === c && !invert) {
                suffLen++;
            } else if (currChar !== c && invert) {
                suffLen++;
            } else {
                break;
            }
        }
        return str.slice(0, l - suffLen);
    }
    function findClosingBracket(str, b) {
        if (str.indexOf(b[1]) === -1) {
            return -1;
        }
        let level = 0;
        for (let i = 0; i < str.length; i++) {
            if (str[i] === "\\") {
                i++;
            } else if (str[i] === b[0]) {
                level++;
            } else if (str[i] === b[1]) {
                level--;
                if (level < 0) {
                    return i;
                }
            }
        }
        return -1;
    }
    function outputLink(cap, link, raw, lexer) {
        const href = link.href;
        const title = link.title ? escape$1(link.title) : null;
        const text = cap[1].replace(/\\([\[\]])/g, "$1");
        if (cap[0].charAt(0) !== "!") {
            lexer.state.inLink = true;
            const token = {
                type: "link",
                raw: raw,
                href: href,
                title: title,
                text: text,
                tokens: lexer.inlineTokens(text)
            };
            lexer.state.inLink = false;
            return token;
        }
        return {
            type: "image",
            raw: raw,
            href: href,
            title: title,
            text: escape$1(text)
        };
    }
    function indentCodeCompensation(raw, text) {
        const matchIndentToCode = raw.match(/^(\s+)(?:```)/);
        if (matchIndentToCode === null) {
            return text;
        }
        const indentToCode = matchIndentToCode[1];
        return text.split("\n").map((node => {
            const matchIndentInNode = node.match(/^\s+/);
            if (matchIndentInNode === null) {
                return node;
            }
            const [indentInNode] = matchIndentInNode;
            if (indentInNode.length >= indentToCode.length) {
                return node.slice(indentToCode.length);
            }
            return node;
        })).join("\n");
    }
    class _Tokenizer {
        options;
        rules;
        lexer;
        constructor(options) {
            this.options = options || _defaults;
        }
        space(src) {
            const cap = this.rules.block.newline.exec(src);
            if (cap && cap[0].length > 0) {
                return {
                    type: "space",
                    raw: cap[0]
                };
            }
        }
        code(src) {
            const cap = this.rules.block.code.exec(src);
            if (cap) {
                const text = cap[0].replace(/^ {1,4}/gm, "");
                return {
                    type: "code",
                    raw: cap[0],
                    codeBlockStyle: "indented",
                    text: !this.options.pedantic ? rtrim(text, "\n") : text
                };
            }
        }
        fences(src) {
            const cap = this.rules.block.fences.exec(src);
            if (cap) {
                const raw = cap[0];
                const text = indentCodeCompensation(raw, cap[3] || "");
                return {
                    type: "code",
                    raw: raw,
                    lang: cap[2] ? cap[2].trim().replace(this.rules.inline.anyPunctuation, "$1") : cap[2],
                    text: text
                };
            }
        }
        heading(src) {
            const cap = this.rules.block.heading.exec(src);
            if (cap) {
                let text = cap[2].trim();
                if (/#$/.test(text)) {
                    const trimmed = rtrim(text, "#");
                    if (this.options.pedantic) {
                        text = trimmed.trim();
                    } else if (!trimmed || / $/.test(trimmed)) {
                        text = trimmed.trim();
                    }
                }
                return {
                    type: "heading",
                    raw: cap[0],
                    depth: cap[1].length,
                    text: text,
                    tokens: this.lexer.inline(text)
                };
            }
        }
        hr(src) {
            const cap = this.rules.block.hr.exec(src);
            if (cap) {
                return {
                    type: "hr",
                    raw: cap[0]
                };
            }
        }
        blockquote(src) {
            const cap = this.rules.block.blockquote.exec(src);
            if (cap) {
                let text = cap[0].replace(/\n {0,3}((?:=+|-+) *)(?=\n|$)/g, "\n    $1");
                text = rtrim(text.replace(/^ *>[ \t]?/gm, ""), "\n");
                const top = this.lexer.state.top;
                this.lexer.state.top = true;
                const tokens = this.lexer.blockTokens(text);
                this.lexer.state.top = top;
                return {
                    type: "blockquote",
                    raw: cap[0],
                    tokens: tokens,
                    text: text
                };
            }
        }
        list(src) {
            let cap = this.rules.block.list.exec(src);
            if (cap) {
                let bull = cap[1].trim();
                const isordered = bull.length > 1;
                const list = {
                    type: "list",
                    raw: "",
                    ordered: isordered,
                    start: isordered ? +bull.slice(0, -1) : "",
                    loose: false,
                    items: []
                };
                bull = isordered ? `\\d{1,9}\\${bull.slice(-1)}` : `\\${bull}`;
                if (this.options.pedantic) {
                    bull = isordered ? bull : "[*+-]";
                }
                const itemRegex = new RegExp(`^( {0,3}${bull})((?:[\t ][^\\n]*)?(?:\\n|$))`);
                let raw = "";
                let itemContents = "";
                let endsWithBlankLine = false;
                while (src) {
                    let endEarly = false;
                    if (!(cap = itemRegex.exec(src))) {
                        break;
                    }
                    if (this.rules.block.hr.test(src)) {
                        break;
                    }
                    raw = cap[0];
                    src = src.substring(raw.length);
                    let line = cap[2].split("\n", 1)[0].replace(/^\t+/, (t => " ".repeat(3 * t.length)));
                    let nextLine = src.split("\n", 1)[0];
                    let indent = 0;
                    if (this.options.pedantic) {
                        indent = 2;
                        itemContents = line.trimStart();
                    } else {
                        indent = cap[2].search(/[^ ]/);
                        indent = indent > 4 ? 1 : indent;
                        itemContents = line.slice(indent);
                        indent += cap[1].length;
                    }
                    let blankLine = false;
                    if (!line && /^ *$/.test(nextLine)) {
                        raw += nextLine + "\n";
                        src = src.substring(nextLine.length + 1);
                        endEarly = true;
                    }
                    if (!endEarly) {
                        const nextBulletRegex = new RegExp(`^ {0,${Math.min(3, indent - 1)}}(?:[*+-]|\\d{1,9}[.)])((?:[ \t][^\\n]*)?(?:\\n|$))`);
                        const hrRegex = new RegExp(`^ {0,${Math.min(3, indent - 1)}}((?:- *){3,}|(?:_ *){3,}|(?:\\* *){3,})(?:\\n+|$)`);
                        const fencesBeginRegex = new RegExp(`^ {0,${Math.min(3, indent - 1)}}(?:\`\`\`|~~~)`);
                        const headingBeginRegex = new RegExp(`^ {0,${Math.min(3, indent - 1)}}#`);
                        while (src) {
                            const rawLine = src.split("\n", 1)[0];
                            nextLine = rawLine;
                            if (this.options.pedantic) {
                                nextLine = nextLine.replace(/^ {1,4}(?=( {4})*[^ ])/g, "  ");
                            }
                            if (fencesBeginRegex.test(nextLine)) {
                                break;
                            }
                            if (headingBeginRegex.test(nextLine)) {
                                break;
                            }
                            if (nextBulletRegex.test(nextLine)) {
                                break;
                            }
                            if (hrRegex.test(src)) {
                                break;
                            }
                            if (nextLine.search(/[^ ]/) >= indent || !nextLine.trim()) {
                                itemContents += "\n" + nextLine.slice(indent);
                            } else {
                                if (blankLine) {
                                    break;
                                }
                                if (line.search(/[^ ]/) >= 4) {
                                    break;
                                }
                                if (fencesBeginRegex.test(line)) {
                                    break;
                                }
                                if (headingBeginRegex.test(line)) {
                                    break;
                                }
                                if (hrRegex.test(line)) {
                                    break;
                                }
                                itemContents += "\n" + nextLine;
                            }
                            if (!blankLine && !nextLine.trim()) {
                                blankLine = true;
                            }
                            raw += rawLine + "\n";
                            src = src.substring(rawLine.length + 1);
                            line = nextLine.slice(indent);
                        }
                    }
                    if (!list.loose) {
                        if (endsWithBlankLine) {
                            list.loose = true;
                        } else if (/\n *\n *$/.test(raw)) {
                            endsWithBlankLine = true;
                        }
                    }
                    let istask = null;
                    let ischecked;
                    if (this.options.gfm) {
                        istask = /^\[[ xX]\] /.exec(itemContents);
                        if (istask) {
                            ischecked = istask[0] !== "[ ] ";
                            itemContents = itemContents.replace(/^\[[ xX]\] +/, "");
                        }
                    }
                    list.items.push({
                        type: "list_item",
                        raw: raw,
                        task: !!istask,
                        checked: ischecked,
                        loose: false,
                        text: itemContents,
                        tokens: []
                    });
                    list.raw += raw;
                }
                list.items[list.items.length - 1].raw = raw.trimEnd();
                list.items[list.items.length - 1].text = itemContents.trimEnd();
                list.raw = list.raw.trimEnd();
                for (let i = 0; i < list.items.length; i++) {
                    this.lexer.state.top = false;
                    list.items[i].tokens = this.lexer.blockTokens(list.items[i].text, []);
                    if (!list.loose) {
                        const spacers = list.items[i].tokens.filter((t => t.type === "space"));
                        const hasMultipleLineBreaks = spacers.length > 0 && spacers.some((t => /\n.*\n/.test(t.raw)));
                        list.loose = hasMultipleLineBreaks;
                    }
                }
                if (list.loose) {
                    for (let i = 0; i < list.items.length; i++) {
                        list.items[i].loose = true;
                    }
                }
                return list;
            }
        }
        html(src) {
            const cap = this.rules.block.html.exec(src);
            if (cap) {
                const token = {
                    type: "html",
                    block: true,
                    raw: cap[0],
                    pre: cap[1] === "pre" || cap[1] === "script" || cap[1] === "style",
                    text: cap[0]
                };
                return token;
            }
        }
        def(src) {
            const cap = this.rules.block.def.exec(src);
            if (cap) {
                const tag = cap[1].toLowerCase().replace(/\s+/g, " ");
                const href = cap[2] ? cap[2].replace(/^<(.*)>$/, "$1").replace(this.rules.inline.anyPunctuation, "$1") : "";
                const title = cap[3] ? cap[3].substring(1, cap[3].length - 1).replace(this.rules.inline.anyPunctuation, "$1") : cap[3];
                return {
                    type: "def",
                    tag: tag,
                    raw: cap[0],
                    href: href,
                    title: title
                };
            }
        }
        table(src) {
            const cap = this.rules.block.table.exec(src);
            if (!cap) {
                return;
            }
            if (!/[:|]/.test(cap[2])) {
                return;
            }
            const headers = splitCells(cap[1]);
            const aligns = cap[2].replace(/^\||\| *$/g, "").split("|");
            const rows = cap[3] && cap[3].trim() ? cap[3].replace(/\n[ \t]*$/, "").split("\n") : [];
            const item = {
                type: "table",
                raw: cap[0],
                header: [],
                align: [],
                rows: []
            };
            if (headers.length !== aligns.length) {
                return;
            }
            for (const align of aligns) {
                if (/^ *-+: *$/.test(align)) {
                    item.align.push("right");
                } else if (/^ *:-+: *$/.test(align)) {
                    item.align.push("center");
                } else if (/^ *:-+ *$/.test(align)) {
                    item.align.push("left");
                } else {
                    item.align.push(null);
                }
            }
            for (const header of headers) {
                item.header.push({
                    text: header,
                    tokens: this.lexer.inline(header)
                });
            }
            for (const row of rows) {
                item.rows.push(splitCells(row, item.header.length).map((cell => ({
                    text: cell,
                    tokens: this.lexer.inline(cell)
                }))));
            }
            return item;
        }
        lheading(src) {
            const cap = this.rules.block.lheading.exec(src);
            if (cap) {
                return {
                    type: "heading",
                    raw: cap[0],
                    depth: cap[2].charAt(0) === "=" ? 1 : 2,
                    text: cap[1],
                    tokens: this.lexer.inline(cap[1])
                };
            }
        }
        paragraph(src) {
            const cap = this.rules.block.paragraph.exec(src);
            if (cap) {
                const text = cap[1].charAt(cap[1].length - 1) === "\n" ? cap[1].slice(0, -1) : cap[1];
                return {
                    type: "paragraph",
                    raw: cap[0],
                    text: text,
                    tokens: this.lexer.inline(text)
                };
            }
        }
        text(src) {
            const cap = this.rules.block.text.exec(src);
            if (cap) {
                return {
                    type: "text",
                    raw: cap[0],
                    text: cap[0],
                    tokens: this.lexer.inline(cap[0])
                };
            }
        }
        escape(src) {
            const cap = this.rules.inline.escape.exec(src);
            if (cap) {
                return {
                    type: "escape",
                    raw: cap[0],
                    text: escape$1(cap[1])
                };
            }
        }
        tag(src) {
            const cap = this.rules.inline.tag.exec(src);
            if (cap) {
                if (!this.lexer.state.inLink && /^<a /i.test(cap[0])) {
                    this.lexer.state.inLink = true;
                } else if (this.lexer.state.inLink && /^<\/a>/i.test(cap[0])) {
                    this.lexer.state.inLink = false;
                }
                if (!this.lexer.state.inRawBlock && /^<(pre|code|kbd|script)(\s|>)/i.test(cap[0])) {
                    this.lexer.state.inRawBlock = true;
                } else if (this.lexer.state.inRawBlock && /^<\/(pre|code|kbd|script)(\s|>)/i.test(cap[0])) {
                    this.lexer.state.inRawBlock = false;
                }
                return {
                    type: "html",
                    raw: cap[0],
                    inLink: this.lexer.state.inLink,
                    inRawBlock: this.lexer.state.inRawBlock,
                    block: false,
                    text: cap[0]
                };
            }
        }
        link(src) {
            const cap = this.rules.inline.link.exec(src);
            if (cap) {
                const trimmedUrl = cap[2].trim();
                if (!this.options.pedantic && /^</.test(trimmedUrl)) {
                    if (!/>$/.test(trimmedUrl)) {
                        return;
                    }
                    const rtrimSlash = rtrim(trimmedUrl.slice(0, -1), "\\");
                    if ((trimmedUrl.length - rtrimSlash.length) % 2 === 0) {
                        return;
                    }
                } else {
                    const lastParenIndex = findClosingBracket(cap[2], "()");
                    if (lastParenIndex > -1) {
                        const start = cap[0].indexOf("!") === 0 ? 5 : 4;
                        const linkLen = start + cap[1].length + lastParenIndex;
                        cap[2] = cap[2].substring(0, lastParenIndex);
                        cap[0] = cap[0].substring(0, linkLen).trim();
                        cap[3] = "";
                    }
                }
                let href = cap[2];
                let title = "";
                if (this.options.pedantic) {
                    const link = /^([^'"]*[^\s])\s+(['"])(.*)\2/.exec(href);
                    if (link) {
                        href = link[1];
                        title = link[3];
                    }
                } else {
                    title = cap[3] ? cap[3].slice(1, -1) : "";
                }
                href = href.trim();
                if (/^</.test(href)) {
                    if (this.options.pedantic && !/>$/.test(trimmedUrl)) {
                        href = href.slice(1);
                    } else {
                        href = href.slice(1, -1);
                    }
                }
                return outputLink(cap, {
                    href: href ? href.replace(this.rules.inline.anyPunctuation, "$1") : href,
                    title: title ? title.replace(this.rules.inline.anyPunctuation, "$1") : title
                }, cap[0], this.lexer);
            }
        }
        reflink(src, links) {
            let cap;
            if ((cap = this.rules.inline.reflink.exec(src)) || (cap = this.rules.inline.nolink.exec(src))) {
                const linkString = (cap[2] || cap[1]).replace(/\s+/g, " ");
                const link = links[linkString.toLowerCase()];
                if (!link) {
                    const text = cap[0].charAt(0);
                    return {
                        type: "text",
                        raw: text,
                        text: text
                    };
                }
                return outputLink(cap, link, cap[0], this.lexer);
            }
        }
        emStrong(src, maskedSrc) {
            let prevChar = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : "";
            let match = this.rules.inline.emStrongLDelim.exec(src);
            if (!match) return;
            if (match[3] && prevChar.match(/[\p{L}\p{N}]/u)) return;
            const nextChar = match[1] || match[2] || "";
            if (!nextChar || !prevChar || this.rules.inline.punctuation.exec(prevChar)) {
                const lLength = [ ...match[0] ].length - 1;
                let rDelim, rLength, delimTotal = lLength, midDelimTotal = 0;
                const endReg = match[0][0] === "*" ? this.rules.inline.emStrongRDelimAst : this.rules.inline.emStrongRDelimUnd;
                endReg.lastIndex = 0;
                maskedSrc = maskedSrc.slice(-1 * src.length + lLength);
                while ((match = endReg.exec(maskedSrc)) != null) {
                    rDelim = match[1] || match[2] || match[3] || match[4] || match[5] || match[6];
                    if (!rDelim) continue;
                    rLength = [ ...rDelim ].length;
                    if (match[3] || match[4]) {
                        delimTotal += rLength;
                        continue;
                    } else if (match[5] || match[6]) {
                        if (lLength % 3 && !((lLength + rLength) % 3)) {
                            midDelimTotal += rLength;
                            continue;
                        }
                    }
                    delimTotal -= rLength;
                    if (delimTotal > 0) continue;
                    rLength = Math.min(rLength, rLength + delimTotal + midDelimTotal);
                    const lastCharLength = [ ...match[0] ][0].length;
                    const raw = src.slice(0, lLength + match.index + lastCharLength + rLength);
                    if (Math.min(lLength, rLength) % 2) {
                        const text = raw.slice(1, -1);
                        return {
                            type: "em",
                            raw: raw,
                            text: text,
                            tokens: this.lexer.inlineTokens(text)
                        };
                    }
                    const text = raw.slice(2, -2);
                    return {
                        type: "strong",
                        raw: raw,
                        text: text,
                        tokens: this.lexer.inlineTokens(text)
                    };
                }
            }
        }
        codespan(src) {
            const cap = this.rules.inline.code.exec(src);
            if (cap) {
                let text = cap[2].replace(/\n/g, " ");
                const hasNonSpaceChars = /[^ ]/.test(text);
                const hasSpaceCharsOnBothEnds = /^ /.test(text) && / $/.test(text);
                if (hasNonSpaceChars && hasSpaceCharsOnBothEnds) {
                    text = text.substring(1, text.length - 1);
                }
                text = escape$1(text, true);
                return {
                    type: "codespan",
                    raw: cap[0],
                    text: text
                };
            }
        }
        br(src) {
            const cap = this.rules.inline.br.exec(src);
            if (cap) {
                return {
                    type: "br",
                    raw: cap[0]
                };
            }
        }
        del(src) {
            const cap = this.rules.inline.del.exec(src);
            if (cap) {
                return {
                    type: "del",
                    raw: cap[0],
                    text: cap[2],
                    tokens: this.lexer.inlineTokens(cap[2])
                };
            }
        }
        autolink(src) {
            const cap = this.rules.inline.autolink.exec(src);
            if (cap) {
                let text, href;
                if (cap[2] === "@") {
                    text = escape$1(cap[1]);
                    href = "mailto:" + text;
                } else {
                    text = escape$1(cap[1]);
                    href = text;
                }
                return {
                    type: "link",
                    raw: cap[0],
                    text: text,
                    href: href,
                    tokens: [ {
                        type: "text",
                        raw: text,
                        text: text
                    } ]
                };
            }
        }
        url(src) {
            let cap;
            if (cap = this.rules.inline.url.exec(src)) {
                let text, href;
                if (cap[2] === "@") {
                    text = escape$1(cap[0]);
                    href = "mailto:" + text;
                } else {
                    let prevCapZero;
                    do {
                        prevCapZero = cap[0];
                        cap[0] = this.rules.inline._backpedal.exec(cap[0])?.[0] ?? "";
                    } while (prevCapZero !== cap[0]);
                    text = escape$1(cap[0]);
                    if (cap[1] === "www.") {
                        href = "http://" + cap[0];
                    } else {
                        href = cap[0];
                    }
                }
                return {
                    type: "link",
                    raw: cap[0],
                    text: text,
                    href: href,
                    tokens: [ {
                        type: "text",
                        raw: text,
                        text: text
                    } ]
                };
            }
        }
        inlineText(src) {
            const cap = this.rules.inline.text.exec(src);
            if (cap) {
                let text;
                if (this.lexer.state.inRawBlock) {
                    text = cap[0];
                } else {
                    text = escape$1(cap[0]);
                }
                return {
                    type: "text",
                    raw: cap[0],
                    text: text
                };
            }
        }
    }
    const newline = /^(?: *(?:\n|$))+/;
    const blockCode = /^( {4}[^\n]+(?:\n(?: *(?:\n|$))*)?)+/;
    const fences = /^ {0,3}(`{3,}(?=[^`\n]*(?:\n|$))|~{3,})([^\n]*)(?:\n|$)(?:|([\s\S]*?)(?:\n|$))(?: {0,3}\1[~`]* *(?=\n|$)|$)/;
    const hr = /^ {0,3}((?:-[\t ]*){3,}|(?:_[ \t]*){3,}|(?:\*[ \t]*){3,})(?:\n+|$)/;
    const heading = /^ {0,3}(#{1,6})(?=\s|$)(.*)(?:\n+|$)/;
    const bullet = /(?:[*+-]|\d{1,9}[.)])/;
    const lheading = edit(/^(?!bull |blockCode|fences|blockquote|heading|html)((?:.|\n(?!\s*?\n|bull |blockCode|fences|blockquote|heading|html))+?)\n {0,3}(=+|-+) *(?:\n+|$)/).replace(/bull/g, bullet).replace(/blockCode/g, / {4}/).replace(/fences/g, / {0,3}(?:`{3,}|~{3,})/).replace(/blockquote/g, / {0,3}>/).replace(/heading/g, / {0,3}#{1,6}/).replace(/html/g, / {0,3}<[^\n>]+>\n/).getRegex();
    const _paragraph = /^([^\n]+(?:\n(?!hr|heading|lheading|blockquote|fences|list|html|table| +\n)[^\n]+)*)/;
    const blockText = /^[^\n]+/;
    const _blockLabel = /(?!\s*\])(?:\\.|[^\[\]\\])+/;
    const def = edit(/^ {0,3}\[(label)\]: *(?:\n *)?([^<\s][^\s]*|<.*?>)(?:(?: +(?:\n *)?| *\n *)(title))? *(?:\n+|$)/).replace("label", _blockLabel).replace("title", /(?:"(?:\\"?|[^"\\])*"|'[^'\n]*(?:\n[^'\n]+)*\n?'|\([^()]*\))/).getRegex();
    const list = edit(/^( {0,3}bull)([ \t][^\n]+?)?(?:\n|$)/).replace(/bull/g, bullet).getRegex();
    const _tag = "address|article|aside|base|basefont|blockquote|body|caption" + "|center|col|colgroup|dd|details|dialog|dir|div|dl|dt|fieldset|figcaption" + "|figure|footer|form|frame|frameset|h[1-6]|head|header|hr|html|iframe" + "|legend|li|link|main|menu|menuitem|meta|nav|noframes|ol|optgroup|option" + "|p|param|search|section|summary|table|tbody|td|tfoot|th|thead|title" + "|tr|track|ul";
    const _comment = /<!--(?:-?>|[\s\S]*?(?:-->|$))/;
    const html = edit("^ {0,3}(?:" + "<(script|pre|style|textarea)[\\s>][\\s\\S]*?(?:</\\1>[^\\n]*\\n+|$)" + "|comment[^\\n]*(\\n+|$)" + "|<\\?[\\s\\S]*?(?:\\?>\\n*|$)" + "|<![A-Z][\\s\\S]*?(?:>\\n*|$)" + "|<!\\[CDATA\\[[\\s\\S]*?(?:\\]\\]>\\n*|$)" + "|</?(tag)(?: +|\\n|/?>)[\\s\\S]*?(?:(?:\\n *)+\\n|$)" + "|<(?!script|pre|style|textarea)([a-z][\\w-]*)(?:attribute)*? */?>(?=[ \\t]*(?:\\n|$))[\\s\\S]*?(?:(?:\\n *)+\\n|$)" + "|</(?!script|pre|style|textarea)[a-z][\\w-]*\\s*>(?=[ \\t]*(?:\\n|$))[\\s\\S]*?(?:(?:\\n *)+\\n|$)" + ")", "i").replace("comment", _comment).replace("tag", _tag).replace("attribute", / +[a-zA-Z:_][\w.:-]*(?: *= *"[^"\n]*"| *= *'[^'\n]*'| *= *[^\s"'=<>`]+)?/).getRegex();
    const paragraph = edit(_paragraph).replace("hr", hr).replace("heading", " {0,3}#{1,6}(?:\\s|$)").replace("|lheading", "").replace("|table", "").replace("blockquote", " {0,3}>").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", _tag).getRegex();
    const blockquote = edit(/^( {0,3}> ?(paragraph|[^\n]*)(?:\n|$))+/).replace("paragraph", paragraph).getRegex();
    const blockNormal = {
        blockquote: blockquote,
        code: blockCode,
        def: def,
        fences: fences,
        heading: heading,
        hr: hr,
        html: html,
        lheading: lheading,
        list: list,
        newline: newline,
        paragraph: paragraph,
        table: noopTest,
        text: blockText
    };
    const gfmTable = edit("^ *([^\\n ].*)\\n" + " {0,3}((?:\\| *)?:?-+:? *(?:\\| *:?-+:? *)*(?:\\| *)?)" + "(?:\\n((?:(?! *\\n|hr|heading|blockquote|code|fences|list|html).*(?:\\n|$))*)\\n*|$)").replace("hr", hr).replace("heading", " {0,3}#{1,6}(?:\\s|$)").replace("blockquote", " {0,3}>").replace("code", " {4}[^\\n]").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", _tag).getRegex();
    const blockGfm = {
        ...blockNormal,
        table: gfmTable,
        paragraph: edit(_paragraph).replace("hr", hr).replace("heading", " {0,3}#{1,6}(?:\\s|$)").replace("|lheading", "").replace("table", gfmTable).replace("blockquote", " {0,3}>").replace("fences", " {0,3}(?:`{3,}(?=[^`\\n]*\\n)|~{3,})[^\\n]*\\n").replace("list", " {0,3}(?:[*+-]|1[.)]) ").replace("html", "</?(?:tag)(?: +|\\n|/?>)|<(?:script|pre|style|textarea|!--)").replace("tag", _tag).getRegex()
    };
    const blockPedantic = {
        ...blockNormal,
        html: edit("^ *(?:comment *(?:\\n|\\s*$)" + "|<(tag)[\\s\\S]+?</\\1> *(?:\\n{2,}|\\s*$)" + "|<tag(?:\"[^\"]*\"|'[^']*'|\\s[^'\"/>\\s]*)*?/?> *(?:\\n{2,}|\\s*$))").replace("comment", _comment).replace(/tag/g, "(?!(?:" + "a|em|strong|small|s|cite|q|dfn|abbr|data|time|code|var|samp|kbd|sub" + "|sup|i|b|u|mark|ruby|rt|rp|bdi|bdo|span|br|wbr|ins|del|img)" + "\\b)\\w+(?!:|[^\\w\\s@]*@)\\b").getRegex(),
        def: /^ *\[([^\]]+)\]: *<?([^\s>]+)>?(?: +(["(][^\n]+[")]))? *(?:\n+|$)/,
        heading: /^(#{1,6})(.*)(?:\n+|$)/,
        fences: noopTest,
        lheading: /^(.+?)\n {0,3}(=+|-+) *(?:\n+|$)/,
        paragraph: edit(_paragraph).replace("hr", hr).replace("heading", " *#{1,6} *[^\n]").replace("lheading", lheading).replace("|table", "").replace("blockquote", " {0,3}>").replace("|fences", "").replace("|list", "").replace("|html", "").replace("|tag", "").getRegex()
    };
    const escape = /^\\([!"#$%&'()*+,\-./:;<=>?@\[\]\\^_`{|}~])/;
    const inlineCode = /^(`+)([^`]|[^`][\s\S]*?[^`])\1(?!`)/;
    const br = /^( {2,}|\\)\n(?!\s*$)/;
    const inlineText = /^(`+|[^`])(?:(?= {2,}\n)|[\s\S]*?(?:(?=[\\<!\[`*_]|\b_|$)|[^ ](?= {2,}\n)))/;
    const _punctuation = "\\p{P}\\p{S}";
    const punctuation = edit(/^((?![*_])[\spunctuation])/, "u").replace(/punctuation/g, _punctuation).getRegex();
    const blockSkip = /\[[^[\]]*?\]\([^\(\)]*?\)|`[^`]*?`|<[^<>]*?>/g;
    const emStrongLDelim = edit(/^(?:\*+(?:((?!\*)[punct])|[^\s*]))|^_+(?:((?!_)[punct])|([^\s_]))/, "u").replace(/punct/g, _punctuation).getRegex();
    const emStrongRDelimAst = edit("^[^_*]*?__[^_*]*?\\*[^_*]*?(?=__)" + "|[^*]+(?=[^*])" + "|(?!\\*)[punct](\\*+)(?=[\\s]|$)" + "|[^punct\\s](\\*+)(?!\\*)(?=[punct\\s]|$)" + "|(?!\\*)[punct\\s](\\*+)(?=[^punct\\s])" + "|[\\s](\\*+)(?!\\*)(?=[punct])" + "|(?!\\*)[punct](\\*+)(?!\\*)(?=[punct])" + "|[^punct\\s](\\*+)(?=[^punct\\s])", "gu").replace(/punct/g, _punctuation).getRegex();
    const emStrongRDelimUnd = edit("^[^_*]*?\\*\\*[^_*]*?_[^_*]*?(?=\\*\\*)" + "|[^_]+(?=[^_])" + "|(?!_)[punct](_+)(?=[\\s]|$)" + "|[^punct\\s](_+)(?!_)(?=[punct\\s]|$)" + "|(?!_)[punct\\s](_+)(?=[^punct\\s])" + "|[\\s](_+)(?!_)(?=[punct])" + "|(?!_)[punct](_+)(?!_)(?=[punct])", "gu").replace(/punct/g, _punctuation).getRegex();
    const anyPunctuation = edit(/\\([punct])/, "gu").replace(/punct/g, _punctuation).getRegex();
    const autolink = edit(/^<(scheme:[^\s\x00-\x1f<>]*|email)>/).replace("scheme", /[a-zA-Z][a-zA-Z0-9+.-]{1,31}/).replace("email", /[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+(@)[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+(?![-_])/).getRegex();
    const _inlineComment = edit(_comment).replace("(?:--\x3e|$)", "--\x3e").getRegex();
    const tag = edit("^comment" + "|^</[a-zA-Z][\\w:-]*\\s*>" + "|^<[a-zA-Z][\\w-]*(?:attribute)*?\\s*/?>" + "|^<\\?[\\s\\S]*?\\?>" + "|^<![a-zA-Z]+\\s[\\s\\S]*?>" + "|^<!\\[CDATA\\[[\\s\\S]*?\\]\\]>").replace("comment", _inlineComment).replace("attribute", /\s+[a-zA-Z:_][\w.:-]*(?:\s*=\s*"[^"]*"|\s*=\s*'[^']*'|\s*=\s*[^\s"'=<>`]+)?/).getRegex();
    const _inlineLabel = /(?:\[(?:\\.|[^\[\]\\])*\]|\\.|`[^`]*`|[^\[\]\\`])*?/;
    const link = edit(/^!?\[(label)\]\(\s*(href)(?:\s+(title))?\s*\)/).replace("label", _inlineLabel).replace("href", /<(?:\\.|[^\n<>\\])+>|[^\s\x00-\x1f]*/).replace("title", /"(?:\\"?|[^"\\])*"|'(?:\\'?|[^'\\])*'|\((?:\\\)?|[^)\\])*\)/).getRegex();
    const reflink = edit(/^!?\[(label)\]\[(ref)\]/).replace("label", _inlineLabel).replace("ref", _blockLabel).getRegex();
    const nolink = edit(/^!?\[(ref)\](?:\[\])?/).replace("ref", _blockLabel).getRegex();
    const reflinkSearch = edit("reflink|nolink(?!\\()", "g").replace("reflink", reflink).replace("nolink", nolink).getRegex();
    const inlineNormal = {
        _backpedal: noopTest,
        anyPunctuation: anyPunctuation,
        autolink: autolink,
        blockSkip: blockSkip,
        br: br,
        code: inlineCode,
        del: noopTest,
        emStrongLDelim: emStrongLDelim,
        emStrongRDelimAst: emStrongRDelimAst,
        emStrongRDelimUnd: emStrongRDelimUnd,
        escape: escape,
        link: link,
        nolink: nolink,
        punctuation: punctuation,
        reflink: reflink,
        reflinkSearch: reflinkSearch,
        tag: tag,
        text: inlineText,
        url: noopTest
    };
    const inlinePedantic = {
        ...inlineNormal,
        link: edit(/^!?\[(label)\]\((.*?)\)/).replace("label", _inlineLabel).getRegex(),
        reflink: edit(/^!?\[(label)\]\s*\[([^\]]*)\]/).replace("label", _inlineLabel).getRegex()
    };
    const inlineGfm = {
        ...inlineNormal,
        escape: edit(escape).replace("])", "~|])").getRegex(),
        url: edit(/^((?:ftp|https?):\/\/|www\.)(?:[a-zA-Z0-9\-]+\.?)+[^\s<]*|^email/, "i").replace("email", /[A-Za-z0-9._+-]+(@)[a-zA-Z0-9-_]+(?:\.[a-zA-Z0-9-_]*[a-zA-Z0-9])+(?![-_])/).getRegex(),
        _backpedal: /(?:[^?!.,:;*_'"~()&]+|\([^)]*\)|&(?![a-zA-Z0-9]+;$)|[?!.,:;*_'"~)]+(?!$))+/,
        del: /^(~~?)(?=[^\s~])([\s\S]*?[^\s~])\1(?=[^~]|$)/,
        text: /^([`~]+|[^`~])(?:(?= {2,}\n)|(?=[a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-]+@)|[\s\S]*?(?:(?=[\\<!\[`*~_]|\b_|https?:\/\/|ftp:\/\/|www\.|$)|[^ ](?= {2,}\n)|[^a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-](?=[a-zA-Z0-9.!#$%&'*+\/=?_`{\|}~-]+@)))/
    };
    const inlineBreaks = {
        ...inlineGfm,
        br: edit(br).replace("{2,}", "*").getRegex(),
        text: edit(inlineGfm.text).replace("\\b_", "\\b_| {2,}\\n").replace(/\{2,\}/g, "*").getRegex()
    };
    const block = {
        normal: blockNormal,
        gfm: blockGfm,
        pedantic: blockPedantic
    };
    const inline = {
        normal: inlineNormal,
        gfm: inlineGfm,
        breaks: inlineBreaks,
        pedantic: inlinePedantic
    };
    class _Lexer {
        tokens;
        options;
        state;
        tokenizer;
        inlineQueue;
        constructor(options) {
            this.tokens = [];
            this.tokens.links = Object.create(null);
            this.options = options || _defaults;
            this.options.tokenizer = this.options.tokenizer || new _Tokenizer;
            this.tokenizer = this.options.tokenizer;
            this.tokenizer.options = this.options;
            this.tokenizer.lexer = this;
            this.inlineQueue = [];
            this.state = {
                inLink: false,
                inRawBlock: false,
                top: true
            };
            const rules = {
                block: block.normal,
                inline: inline.normal
            };
            if (this.options.pedantic) {
                rules.block = block.pedantic;
                rules.inline = inline.pedantic;
            } else if (this.options.gfm) {
                rules.block = block.gfm;
                if (this.options.breaks) {
                    rules.inline = inline.breaks;
                } else {
                    rules.inline = inline.gfm;
                }
            }
            this.tokenizer.rules = rules;
        }
        static get rules() {
            return {
                block: block,
                inline: inline
            };
        }
        static lex(src, options) {
            const lexer = new _Lexer(options);
            return lexer.lex(src);
        }
        static lexInline(src, options) {
            const lexer = new _Lexer(options);
            return lexer.inlineTokens(src);
        }
        lex(src) {
            src = src.replace(/\r\n|\r/g, "\n");
            this.blockTokens(src, this.tokens);
            for (let i = 0; i < this.inlineQueue.length; i++) {
                const next = this.inlineQueue[i];
                this.inlineTokens(next.src, next.tokens);
            }
            this.inlineQueue = [];
            return this.tokens;
        }
        blockTokens(src) {
            let tokens = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : [];
            if (this.options.pedantic) {
                src = src.replace(/\t/g, "    ").replace(/^ +$/gm, "");
            } else {
                src = src.replace(/^( *)(\t+)/gm, ((_, leading, tabs) => leading + "    ".repeat(tabs.length)));
            }
            let token;
            let lastToken;
            let cutSrc;
            let lastParagraphClipped;
            while (src) {
                if (this.options.extensions && this.options.extensions.block && this.options.extensions.block.some((extTokenizer => {
                    if (token = extTokenizer.call({
                        lexer: this
                    }, src, tokens)) {
                        src = src.substring(token.raw.length);
                        tokens.push(token);
                        return true;
                    }
                    return false;
                }))) {
                    continue;
                }
                if (token = this.tokenizer.space(src)) {
                    src = src.substring(token.raw.length);
                    if (token.raw.length === 1 && tokens.length > 0) {
                        tokens[tokens.length - 1].raw += "\n";
                    } else {
                        tokens.push(token);
                    }
                    continue;
                }
                if (token = this.tokenizer.code(src)) {
                    src = src.substring(token.raw.length);
                    lastToken = tokens[tokens.length - 1];
                    if (lastToken && (lastToken.type === "paragraph" || lastToken.type === "text")) {
                        lastToken.raw += "\n" + token.raw;
                        lastToken.text += "\n" + token.text;
                        this.inlineQueue[this.inlineQueue.length - 1].src = lastToken.text;
                    } else {
                        tokens.push(token);
                    }
                    continue;
                }
                if (token = this.tokenizer.fences(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.heading(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.hr(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.blockquote(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.list(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.html(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.def(src)) {
                    src = src.substring(token.raw.length);
                    lastToken = tokens[tokens.length - 1];
                    if (lastToken && (lastToken.type === "paragraph" || lastToken.type === "text")) {
                        lastToken.raw += "\n" + token.raw;
                        lastToken.text += "\n" + token.raw;
                        this.inlineQueue[this.inlineQueue.length - 1].src = lastToken.text;
                    } else if (!this.tokens.links[token.tag]) {
                        this.tokens.links[token.tag] = {
                            href: token.href,
                            title: token.title
                        };
                    }
                    continue;
                }
                if (token = this.tokenizer.table(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.lheading(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                cutSrc = src;
                if (this.options.extensions && this.options.extensions.startBlock) {
                    let startIndex = Infinity;
                    const tempSrc = src.slice(1);
                    let tempStart;
                    this.options.extensions.startBlock.forEach((getStartIndex => {
                        tempStart = getStartIndex.call({
                            lexer: this
                        }, tempSrc);
                        if (typeof tempStart === "number" && tempStart >= 0) {
                            startIndex = Math.min(startIndex, tempStart);
                        }
                    }));
                    if (startIndex < Infinity && startIndex >= 0) {
                        cutSrc = src.substring(0, startIndex + 1);
                    }
                }
                if (this.state.top && (token = this.tokenizer.paragraph(cutSrc))) {
                    lastToken = tokens[tokens.length - 1];
                    if (lastParagraphClipped && lastToken.type === "paragraph") {
                        lastToken.raw += "\n" + token.raw;
                        lastToken.text += "\n" + token.text;
                        this.inlineQueue.pop();
                        this.inlineQueue[this.inlineQueue.length - 1].src = lastToken.text;
                    } else {
                        tokens.push(token);
                    }
                    lastParagraphClipped = cutSrc.length !== src.length;
                    src = src.substring(token.raw.length);
                    continue;
                }
                if (token = this.tokenizer.text(src)) {
                    src = src.substring(token.raw.length);
                    lastToken = tokens[tokens.length - 1];
                    if (lastToken && lastToken.type === "text") {
                        lastToken.raw += "\n" + token.raw;
                        lastToken.text += "\n" + token.text;
                        this.inlineQueue.pop();
                        this.inlineQueue[this.inlineQueue.length - 1].src = lastToken.text;
                    } else {
                        tokens.push(token);
                    }
                    continue;
                }
                if (src) {
                    const errMsg = "Infinite loop on byte: " + src.charCodeAt(0);
                    if (this.options.silent) {
                        console.error(errMsg);
                        break;
                    } else {
                        throw new Error(errMsg);
                    }
                }
            }
            this.state.top = true;
            return tokens;
        }
        inline(src) {
            let tokens = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : [];
            this.inlineQueue.push({
                src: src,
                tokens: tokens
            });
            return tokens;
        }
        inlineTokens(src) {
            let tokens = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : [];
            let token, lastToken, cutSrc;
            let maskedSrc = src;
            let match;
            let keepPrevChar, prevChar;
            if (this.tokens.links) {
                const links = Object.keys(this.tokens.links);
                if (links.length > 0) {
                    while ((match = this.tokenizer.rules.inline.reflinkSearch.exec(maskedSrc)) != null) {
                        if (links.includes(match[0].slice(match[0].lastIndexOf("[") + 1, -1))) {
                            maskedSrc = maskedSrc.slice(0, match.index) + "[" + "a".repeat(match[0].length - 2) + "]" + maskedSrc.slice(this.tokenizer.rules.inline.reflinkSearch.lastIndex);
                        }
                    }
                }
            }
            while ((match = this.tokenizer.rules.inline.blockSkip.exec(maskedSrc)) != null) {
                maskedSrc = maskedSrc.slice(0, match.index) + "[" + "a".repeat(match[0].length - 2) + "]" + maskedSrc.slice(this.tokenizer.rules.inline.blockSkip.lastIndex);
            }
            while ((match = this.tokenizer.rules.inline.anyPunctuation.exec(maskedSrc)) != null) {
                maskedSrc = maskedSrc.slice(0, match.index) + "++" + maskedSrc.slice(this.tokenizer.rules.inline.anyPunctuation.lastIndex);
            }
            while (src) {
                if (!keepPrevChar) {
                    prevChar = "";
                }
                keepPrevChar = false;
                if (this.options.extensions && this.options.extensions.inline && this.options.extensions.inline.some((extTokenizer => {
                    if (token = extTokenizer.call({
                        lexer: this
                    }, src, tokens)) {
                        src = src.substring(token.raw.length);
                        tokens.push(token);
                        return true;
                    }
                    return false;
                }))) {
                    continue;
                }
                if (token = this.tokenizer.escape(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.tag(src)) {
                    src = src.substring(token.raw.length);
                    lastToken = tokens[tokens.length - 1];
                    if (lastToken && token.type === "text" && lastToken.type === "text") {
                        lastToken.raw += token.raw;
                        lastToken.text += token.text;
                    } else {
                        tokens.push(token);
                    }
                    continue;
                }
                if (token = this.tokenizer.link(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.reflink(src, this.tokens.links)) {
                    src = src.substring(token.raw.length);
                    lastToken = tokens[tokens.length - 1];
                    if (lastToken && token.type === "text" && lastToken.type === "text") {
                        lastToken.raw += token.raw;
                        lastToken.text += token.text;
                    } else {
                        tokens.push(token);
                    }
                    continue;
                }
                if (token = this.tokenizer.emStrong(src, maskedSrc, prevChar)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.codespan(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.br(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.del(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (token = this.tokenizer.autolink(src)) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                if (!this.state.inLink && (token = this.tokenizer.url(src))) {
                    src = src.substring(token.raw.length);
                    tokens.push(token);
                    continue;
                }
                cutSrc = src;
                if (this.options.extensions && this.options.extensions.startInline) {
                    let startIndex = Infinity;
                    const tempSrc = src.slice(1);
                    let tempStart;
                    this.options.extensions.startInline.forEach((getStartIndex => {
                        tempStart = getStartIndex.call({
                            lexer: this
                        }, tempSrc);
                        if (typeof tempStart === "number" && tempStart >= 0) {
                            startIndex = Math.min(startIndex, tempStart);
                        }
                    }));
                    if (startIndex < Infinity && startIndex >= 0) {
                        cutSrc = src.substring(0, startIndex + 1);
                    }
                }
                if (token = this.tokenizer.inlineText(cutSrc)) {
                    src = src.substring(token.raw.length);
                    if (token.raw.slice(-1) !== "_") {
                        prevChar = token.raw.slice(-1);
                    }
                    keepPrevChar = true;
                    lastToken = tokens[tokens.length - 1];
                    if (lastToken && lastToken.type === "text") {
                        lastToken.raw += token.raw;
                        lastToken.text += token.text;
                    } else {
                        tokens.push(token);
                    }
                    continue;
                }
                if (src) {
                    const errMsg = "Infinite loop on byte: " + src.charCodeAt(0);
                    if (this.options.silent) {
                        console.error(errMsg);
                        break;
                    } else {
                        throw new Error(errMsg);
                    }
                }
            }
            return tokens;
        }
    }
    class _Renderer {
        options;
        constructor(options) {
            this.options = options || _defaults;
        }
        code(code, infostring, escaped) {
            const lang = (infostring || "").match(/^\S*/)?.[0];
            code = code.replace(/\n$/, "") + "\n";
            if (!lang) {
                return "<pre><code>" + (escaped ? code : escape$1(code, true)) + "</code></pre>\n";
            }
            return '<pre><code class="language-' + escape$1(lang) + '">' + (escaped ? code : escape$1(code, true)) + "</code></pre>\n";
        }
        blockquote(quote) {
            return `<blockquote>\n${quote}</blockquote>\n`;
        }
        html(html, block) {
            return html;
        }
        heading(text, level, raw) {
            return `<h${level}>${text}</h${level}>\n`;
        }
        hr() {
            return "<hr>\n";
        }
        list(body, ordered, start) {
            const type = ordered ? "ol" : "ul";
            const startatt = ordered && start !== 1 ? ' start="' + start + '"' : "";
            return "<" + type + startatt + ">\n" + body + "</" + type + ">\n";
        }
        listitem(text, task, checked) {
            return `<li>${text}</li>\n`;
        }
        checkbox(checked) {
            return "<input " + (checked ? 'checked="" ' : "") + 'disabled="" type="checkbox">';
        }
        paragraph(text) {
            return `<p>${text}</p>\n`;
        }
        table(header, body) {
            if (body) body = `<tbody>${body}</tbody>`;
            return "<table>\n" + "<thead>\n" + header + "</thead>\n" + body + "</table>\n";
        }
        tablerow(content) {
            return `<tr>\n${content}</tr>\n`;
        }
        tablecell(content, flags) {
            const type = flags.header ? "th" : "td";
            const tag = flags.align ? `<${type} align="${flags.align}">` : `<${type}>`;
            return tag + content + `</${type}>\n`;
        }
        strong(text) {
            return `<strong>${text}</strong>`;
        }
        em(text) {
            return `<em>${text}</em>`;
        }
        codespan(text) {
            return `<code>${text}</code>`;
        }
        br() {
            return "<br>";
        }
        del(text) {
            return `<del>${text}</del>`;
        }
        link(href, title, text) {
            const cleanHref = cleanUrl(href);
            if (cleanHref === null) {
                return text;
            }
            href = cleanHref;
            let out = '<a href="' + href + '"';
            if (title) {
                out += ' title="' + title + '"';
            }
            out += ">" + text + "</a>";
            return out;
        }
        image(href, title, text) {
            const cleanHref = cleanUrl(href);
            if (cleanHref === null) {
                return text;
            }
            href = cleanHref;
            let out = `<img src="${href}" alt="${text}"`;
            if (title) {
                out += ` title="${title}"`;
            }
            out += ">";
            return out;
        }
        text(text) {
            return text;
        }
    }
    class _TextRenderer {
        strong(text) {
            return text;
        }
        em(text) {
            return text;
        }
        codespan(text) {
            return text;
        }
        del(text) {
            return text;
        }
        html(text) {
            return text;
        }
        text(text) {
            return text;
        }
        link(href, title, text) {
            return "" + text;
        }
        image(href, title, text) {
            return "" + text;
        }
        br() {
            return "";
        }
    }
    class _Parser {
        options;
        renderer;
        textRenderer;
        constructor(options) {
            this.options = options || _defaults;
            this.options.renderer = this.options.renderer || new _Renderer;
            this.renderer = this.options.renderer;
            this.renderer.options = this.options;
            this.textRenderer = new _TextRenderer;
        }
        static parse(tokens, options) {
            const parser = new _Parser(options);
            return parser.parse(tokens);
        }
        static parseInline(tokens, options) {
            const parser = new _Parser(options);
            return parser.parseInline(tokens);
        }
        parse(tokens) {
            let top = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : true;
            let out = "";
            for (let i = 0; i < tokens.length; i++) {
                const token = tokens[i];
                if (this.options.extensions && this.options.extensions.renderers && this.options.extensions.renderers[token.type]) {
                    const genericToken = token;
                    const ret = this.options.extensions.renderers[genericToken.type].call({
                        parser: this
                    }, genericToken);
                    if (ret !== false || ![ "space", "hr", "heading", "code", "table", "blockquote", "list", "html", "paragraph", "text" ].includes(genericToken.type)) {
                        out += ret || "";
                        continue;
                    }
                }
                switch (token.type) {
                    case "space":
                    {
                        continue;
                    }

                    case "hr":
                    {
                        out += this.renderer.hr();
                        continue;
                    }

                    case "heading":
                    {
                        const headingToken = token;
                        out += this.renderer.heading(this.parseInline(headingToken.tokens), headingToken.depth, unescape(this.parseInline(headingToken.tokens, this.textRenderer)));
                        continue;
                    }

                    case "code":
                    {
                        const codeToken = token;
                        out += this.renderer.code(codeToken.text, codeToken.lang, !!codeToken.escaped);
                        continue;
                    }

                    case "table":
                    {
                        const tableToken = token;
                        let header = "";
                        let cell = "";
                        for (let j = 0; j < tableToken.header.length; j++) {
                            cell += this.renderer.tablecell(this.parseInline(tableToken.header[j].tokens), {
                                header: true,
                                align: tableToken.align[j]
                            });
                        }
                        header += this.renderer.tablerow(cell);
                        let body = "";
                        for (let j = 0; j < tableToken.rows.length; j++) {
                            const row = tableToken.rows[j];
                            cell = "";
                            for (let k = 0; k < row.length; k++) {
                                cell += this.renderer.tablecell(this.parseInline(row[k].tokens), {
                                    header: false,
                                    align: tableToken.align[k]
                                });
                            }
                            body += this.renderer.tablerow(cell);
                        }
                        out += this.renderer.table(header, body);
                        continue;
                    }

                    case "blockquote":
                    {
                        const blockquoteToken = token;
                        const body = this.parse(blockquoteToken.tokens);
                        out += this.renderer.blockquote(body);
                        continue;
                    }

                    case "list":
                    {
                        const listToken = token;
                        const ordered = listToken.ordered;
                        const start = listToken.start;
                        const loose = listToken.loose;
                        let body = "";
                        for (let j = 0; j < listToken.items.length; j++) {
                            const item = listToken.items[j];
                            const checked = item.checked;
                            const task = item.task;
                            let itemBody = "";
                            if (item.task) {
                                const checkbox = this.renderer.checkbox(!!checked);
                                if (loose) {
                                    if (item.tokens.length > 0 && item.tokens[0].type === "paragraph") {
                                        item.tokens[0].text = checkbox + " " + item.tokens[0].text;
                                        if (item.tokens[0].tokens && item.tokens[0].tokens.length > 0 && item.tokens[0].tokens[0].type === "text") {
                                            item.tokens[0].tokens[0].text = checkbox + " " + item.tokens[0].tokens[0].text;
                                        }
                                    } else {
                                        item.tokens.unshift({
                                            type: "text",
                                            text: checkbox + " "
                                        });
                                    }
                                } else {
                                    itemBody += checkbox + " ";
                                }
                            }
                            itemBody += this.parse(item.tokens, loose);
                            body += this.renderer.listitem(itemBody, task, !!checked);
                        }
                        out += this.renderer.list(body, ordered, start);
                        continue;
                    }

                    case "html":
                    {
                        const htmlToken = token;
                        out += this.renderer.html(htmlToken.text, htmlToken.block);
                        continue;
                    }

                    case "paragraph":
                    {
                        const paragraphToken = token;
                        out += this.renderer.paragraph(this.parseInline(paragraphToken.tokens));
                        continue;
                    }

                    case "text":
                    {
                        let textToken = token;
                        let body = textToken.tokens ? this.parseInline(textToken.tokens) : textToken.text;
                        while (i + 1 < tokens.length && tokens[i + 1].type === "text") {
                            textToken = tokens[++i];
                            body += "\n" + (textToken.tokens ? this.parseInline(textToken.tokens) : textToken.text);
                        }
                        out += top ? this.renderer.paragraph(body) : body;
                        continue;
                    }

                    default:
                    {
                        const errMsg = 'Token with "' + token.type + '" type was not found.';
                        if (this.options.silent) {
                            console.error(errMsg);
                            return "";
                        } else {
                            throw new Error(errMsg);
                        }
                    }
                }
            }
            return out;
        }
        parseInline(tokens, renderer) {
            renderer = renderer || this.renderer;
            let out = "";
            for (let i = 0; i < tokens.length; i++) {
                const token = tokens[i];
                if (this.options.extensions && this.options.extensions.renderers && this.options.extensions.renderers[token.type]) {
                    const ret = this.options.extensions.renderers[token.type].call({
                        parser: this
                    }, token);
                    if (ret !== false || ![ "escape", "html", "link", "image", "strong", "em", "codespan", "br", "del", "text" ].includes(token.type)) {
                        out += ret || "";
                        continue;
                    }
                }
                switch (token.type) {
                    case "escape":
                    {
                        const escapeToken = token;
                        out += renderer.text(escapeToken.text);
                        break;
                    }

                    case "html":
                    {
                        const tagToken = token;
                        out += renderer.html(tagToken.text);
                        break;
                    }

                    case "link":
                    {
                        const linkToken = token;
                        out += renderer.link(linkToken.href, linkToken.title, this.parseInline(linkToken.tokens, renderer));
                        break;
                    }

                    case "image":
                    {
                        const imageToken = token;
                        out += renderer.image(imageToken.href, imageToken.title, imageToken.text);
                        break;
                    }

                    case "strong":
                    {
                        const strongToken = token;
                        out += renderer.strong(this.parseInline(strongToken.tokens, renderer));
                        break;
                    }

                    case "em":
                    {
                        const emToken = token;
                        out += renderer.em(this.parseInline(emToken.tokens, renderer));
                        break;
                    }

                    case "codespan":
                    {
                        const codespanToken = token;
                        out += renderer.codespan(codespanToken.text);
                        break;
                    }

                    case "br":
                    {
                        out += renderer.br();
                        break;
                    }

                    case "del":
                    {
                        const delToken = token;
                        out += renderer.del(this.parseInline(delToken.tokens, renderer));
                        break;
                    }

                    case "text":
                    {
                        const textToken = token;
                        out += renderer.text(textToken.text);
                        break;
                    }

                    default:
                    {
                        const errMsg = 'Token with "' + token.type + '" type was not found.';
                        if (this.options.silent) {
                            console.error(errMsg);
                            return "";
                        } else {
                            throw new Error(errMsg);
                        }
                    }
                }
            }
            return out;
        }
    }
    class _Hooks {
        options;
        constructor(options) {
            this.options = options || _defaults;
        }
        static passThroughHooks=new Set([ "preprocess", "postprocess", "processAllTokens" ]);
        preprocess(markdown) {
            return markdown;
        }
        postprocess(html) {
            return html;
        }
        processAllTokens(tokens) {
            return tokens;
        }
    }
    class Marked {
        defaults=_getDefaults();
        options=this.setOptions;
        parse=this.#parseMarkdown(_Lexer.lex, _Parser.parse);
        parseInline=this.#parseMarkdown(_Lexer.lexInline, _Parser.parseInline);
        Parser=_Parser;
        Renderer=_Renderer;
        TextRenderer=_TextRenderer;
        Lexer=_Lexer;
        Tokenizer=_Tokenizer;
        Hooks=_Hooks;
        constructor() {
            this.use(...arguments);
        }
        walkTokens(tokens, callback) {
            let values = [];
            for (const token of tokens) {
                values = values.concat(callback.call(this, token));
                switch (token.type) {
                    case "table":
                    {
                        const tableToken = token;
                        for (const cell of tableToken.header) {
                            values = values.concat(this.walkTokens(cell.tokens, callback));
                        }
                        for (const row of tableToken.rows) {
                            for (const cell of row) {
                                values = values.concat(this.walkTokens(cell.tokens, callback));
                            }
                        }
                        break;
                    }

                    case "list":
                    {
                        const listToken = token;
                        values = values.concat(this.walkTokens(listToken.items, callback));
                        break;
                    }

                    default:
                    {
                        const genericToken = token;
                        if (this.defaults.extensions?.childTokens?.[genericToken.type]) {
                            this.defaults.extensions.childTokens[genericToken.type].forEach((childTokens => {
                                const tokens = genericToken[childTokens].flat(Infinity);
                                values = values.concat(this.walkTokens(tokens, callback));
                            }));
                        } else if (genericToken.tokens) {
                            values = values.concat(this.walkTokens(genericToken.tokens, callback));
                        }
                    }
                }
            }
            return values;
        }
        use() {
            const extensions = this.defaults.extensions || {
                renderers: {},
                childTokens: {}
            };
            for (var _len = arguments.length, args = new Array(_len), _key = 0; _key < _len; _key++) {
                args[_key] = arguments[_key];
            }
            args.forEach((pack => {
                const opts = {
                    ...pack
                };
                opts.async = this.defaults.async || opts.async || false;
                if (pack.extensions) {
                    pack.extensions.forEach((ext => {
                        if (!ext.name) {
                            throw new Error("extension name required");
                        }
                        if ("renderer" in ext) {
                            const prevRenderer = extensions.renderers[ext.name];
                            if (prevRenderer) {
                                extensions.renderers[ext.name] = function() {
                                    for (var _len2 = arguments.length, args = new Array(_len2), _key2 = 0; _key2 < _len2; _key2++) {
                                        args[_key2] = arguments[_key2];
                                    }
                                    let ret = ext.renderer.apply(this, args);
                                    if (ret === false) {
                                        ret = prevRenderer.apply(this, args);
                                    }
                                    return ret;
                                };
                            } else {
                                extensions.renderers[ext.name] = ext.renderer;
                            }
                        }
                        if ("tokenizer" in ext) {
                            if (!ext.level || ext.level !== "block" && ext.level !== "inline") {
                                throw new Error("extension level must be 'block' or 'inline'");
                            }
                            const extLevel = extensions[ext.level];
                            if (extLevel) {
                                extLevel.unshift(ext.tokenizer);
                            } else {
                                extensions[ext.level] = [ ext.tokenizer ];
                            }
                            if (ext.start) {
                                if (ext.level === "block") {
                                    if (extensions.startBlock) {
                                        extensions.startBlock.push(ext.start);
                                    } else {
                                        extensions.startBlock = [ ext.start ];
                                    }
                                } else if (ext.level === "inline") {
                                    if (extensions.startInline) {
                                        extensions.startInline.push(ext.start);
                                    } else {
                                        extensions.startInline = [ ext.start ];
                                    }
                                }
                            }
                        }
                        if ("childTokens" in ext && ext.childTokens) {
                            extensions.childTokens[ext.name] = ext.childTokens;
                        }
                    }));
                    opts.extensions = extensions;
                }
                if (pack.renderer) {
                    const renderer = this.defaults.renderer || new _Renderer(this.defaults);
                    for (const prop in pack.renderer) {
                        if (!(prop in renderer)) {
                            throw new Error(`renderer '${prop}' does not exist`);
                        }
                        if (prop === "options") {
                            continue;
                        }
                        const rendererProp = prop;
                        const rendererFunc = pack.renderer[rendererProp];
                        const prevRenderer = renderer[rendererProp];
                        renderer[rendererProp] = function() {
                            for (var _len3 = arguments.length, args = new Array(_len3), _key3 = 0; _key3 < _len3; _key3++) {
                                args[_key3] = arguments[_key3];
                            }
                            let ret = rendererFunc.apply(renderer, args);
                            if (ret === false) {
                                ret = prevRenderer.apply(renderer, args);
                            }
                            return ret || "";
                        };
                    }
                    opts.renderer = renderer;
                }
                if (pack.tokenizer) {
                    const tokenizer = this.defaults.tokenizer || new _Tokenizer(this.defaults);
                    for (const prop in pack.tokenizer) {
                        if (!(prop in tokenizer)) {
                            throw new Error(`tokenizer '${prop}' does not exist`);
                        }
                        if ([ "options", "rules", "lexer" ].includes(prop)) {
                            continue;
                        }
                        const tokenizerProp = prop;
                        const tokenizerFunc = pack.tokenizer[tokenizerProp];
                        const prevTokenizer = tokenizer[tokenizerProp];
                        tokenizer[tokenizerProp] = function() {
                            for (var _len4 = arguments.length, args = new Array(_len4), _key4 = 0; _key4 < _len4; _key4++) {
                                args[_key4] = arguments[_key4];
                            }
                            let ret = tokenizerFunc.apply(tokenizer, args);
                            if (ret === false) {
                                ret = prevTokenizer.apply(tokenizer, args);
                            }
                            return ret;
                        };
                    }
                    opts.tokenizer = tokenizer;
                }
                if (pack.hooks) {
                    const hooks = this.defaults.hooks || new _Hooks;
                    for (const prop in pack.hooks) {
                        if (!(prop in hooks)) {
                            throw new Error(`hook '${prop}' does not exist`);
                        }
                        if (prop === "options") {
                            continue;
                        }
                        const hooksProp = prop;
                        const hooksFunc = pack.hooks[hooksProp];
                        const prevHook = hooks[hooksProp];
                        if (_Hooks.passThroughHooks.has(prop)) {
                            hooks[hooksProp] = arg => {
                                if (this.defaults.async) {
                                    return Promise.resolve(hooksFunc.call(hooks, arg)).then((ret => prevHook.call(hooks, ret)));
                                }
                                const ret = hooksFunc.call(hooks, arg);
                                return prevHook.call(hooks, ret);
                            };
                        } else {
                            hooks[hooksProp] = function() {
                                for (var _len5 = arguments.length, args = new Array(_len5), _key5 = 0; _key5 < _len5; _key5++) {
                                    args[_key5] = arguments[_key5];
                                }
                                let ret = hooksFunc.apply(hooks, args);
                                if (ret === false) {
                                    ret = prevHook.apply(hooks, args);
                                }
                                return ret;
                            };
                        }
                    }
                    opts.hooks = hooks;
                }
                if (pack.walkTokens) {
                    const walkTokens = this.defaults.walkTokens;
                    const packWalktokens = pack.walkTokens;
                    opts.walkTokens = function(token) {
                        let values = [];
                        values.push(packWalktokens.call(this, token));
                        if (walkTokens) {
                            values = values.concat(walkTokens.call(this, token));
                        }
                        return values;
                    };
                }
                this.defaults = {
                    ...this.defaults,
                    ...opts
                };
            }));
            return this;
        }
        setOptions(opt) {
            this.defaults = {
                ...this.defaults,
                ...opt
            };
            return this;
        }
        lexer(src, options) {
            return _Lexer.lex(src, options ?? this.defaults);
        }
        parser(tokens, options) {
            return _Parser.parse(tokens, options ?? this.defaults);
        }
        #parseMarkdown(lexer, parser) {
            return (src, options) => {
                const origOpt = {
                    ...options
                };
                const opt = {
                    ...this.defaults,
                    ...origOpt
                };
                if (this.defaults.async === true && origOpt.async === false) {
                    if (!opt.silent) {
                        console.warn("marked(): The async option was set to true by an extension. The async: false option sent to parse will be ignored.");
                    }
                    opt.async = true;
                }
                const throwError = this.#onError(!!opt.silent, !!opt.async);
                if (typeof src === "undefined" || src === null) {
                    return throwError(new Error("marked(): input parameter is undefined or null"));
                }
                if (typeof src !== "string") {
                    return throwError(new Error("marked(): input parameter is of type " + Object.prototype.toString.call(src) + ", string expected"));
                }
                if (opt.hooks) {
                    opt.hooks.options = opt;
                }
                if (opt.async) {
                    return Promise.resolve(opt.hooks ? opt.hooks.preprocess(src) : src).then((src => lexer(src, opt))).then((tokens => opt.hooks ? opt.hooks.processAllTokens(tokens) : tokens)).then((tokens => opt.walkTokens ? Promise.all(this.walkTokens(tokens, opt.walkTokens)).then((() => tokens)) : tokens)).then((tokens => parser(tokens, opt))).then((html => opt.hooks ? opt.hooks.postprocess(html) : html)).catch(throwError);
                }
                try {
                    if (opt.hooks) {
                        src = opt.hooks.preprocess(src);
                    }
                    let tokens = lexer(src, opt);
                    if (opt.hooks) {
                        tokens = opt.hooks.processAllTokens(tokens);
                    }
                    if (opt.walkTokens) {
                        this.walkTokens(tokens, opt.walkTokens);
                    }
                    let html = parser(tokens, opt);
                    if (opt.hooks) {
                        html = opt.hooks.postprocess(html);
                    }
                    return html;
                } catch (e) {
                    return throwError(e);
                }
            };
        }
        #onError(silent, async) {
            return e => {
                e.message += "\nPlease report this to https://github.com/markedjs/marked.";
                if (silent) {
                    const msg = "<p>An error occurred:</p><pre>" + escape$1(e.message + "", true) + "</pre>";
                    if (async) {
                        return Promise.resolve(msg);
                    }
                    return msg;
                }
                if (async) {
                    return Promise.reject(e);
                }
                throw e;
            };
        }
    }
    const markedInstance = new Marked;
    function marked(src, opt) {
        return markedInstance.parse(src, opt);
    }
    marked.options = marked.setOptions = function(options) {
        markedInstance.setOptions(options);
        marked.defaults = markedInstance.defaults;
        changeDefaults(marked.defaults);
        return marked;
    };
    marked.getDefaults = _getDefaults;
    marked.defaults = _defaults;
    marked.use = function() {
        markedInstance.use(...arguments);
        marked.defaults = markedInstance.defaults;
        changeDefaults(marked.defaults);
        return marked;
    };
    marked.walkTokens = function(tokens, callback) {
        return markedInstance.walkTokens(tokens, callback);
    };
    marked.parseInline = markedInstance.parseInline;
    marked.Parser = _Parser;
    marked.parser = _Parser.parse;
    marked.Renderer = _Renderer;
    marked.TextRenderer = _TextRenderer;
    marked.Lexer = _Lexer;
    marked.lexer = _Lexer.lex;
    marked.Tokenizer = _Tokenizer;
    marked.Hooks = _Hooks;
    marked.parse = marked;
    marked.options;
    marked.setOptions;
    marked.use;
    marked.walkTokens;
    marked.parseInline;
    _Parser.parse;
    _Lexer.lex;
    function corner(data, cornerExternalLinkTarget) {
        if (!data) {
            return "";
        }
        if (!/\/\//.test(data)) {
            data = "https://github.com/" + data;
        }
        data = data.replace(/^git\+/, "");
        cornerExternalLinkTarget = cornerExternalLinkTarget || "_blank";
        return `\n    <a href="${data}" target="${cornerExternalLinkTarget}" class="github-corner" aria-label="View source on Github">\n      <svg viewBox="0 0 250 250" aria-hidden="true">\n        <path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path>\n        <path d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2" fill="currentColor" style="transform-origin: 130px 106px;" class="octo-arm"></path>\n        <path d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z" fill="currentColor" class="octo-body"></path>\n      </svg>\n    </a>\n  `;
    }
    function main(config) {
        const name = config.name ? config.name : "";
        const aside = `\n    <button class="sidebar-toggle" title="Press \\ to toggle" aria-label="Toggle primary navigation" aria-keyshortcuts="\\" aria-controls="__sidebar">\n      <div class="sidebar-toggle-button" aria-hidden="true">\n        <span></span><span></span><span></span>\n      </div>\n    </button>\n    <aside id="__sidebar" class="sidebar" role="none">\n      ${config.name ? `\n            <h1 class="app-name"><a class="app-name-link" data-nosearch>${config.logo ? `<img alt="${name}" src=${config.logo} />` : name}</a></h1>\n          ` : ""}\n      <div class="sidebar-nav" role="navigation" aria-label="primary">\x3c!--sidebar--\x3e</div>\n    </aside>\n  `;
        return `\n    <main role="presentation">${aside}\n      <section class="content">\n        <article id="main" class="markdown-section" role="main" tabindex="-1">\x3c!--main--\x3e</article>\n      </section>\n    </main>\n  `;
    }
    function cover() {
        const SL = ", 100%, 85%";
        const bgc = `\n    linear-gradient(\n      to left bottom,\n      hsl(${Math.floor(Math.random() * 255) + SL}) 0%,\n      hsl(${Math.floor(Math.random() * 255) + SL}) 100%\n    )\n  `;
        return `\n    <section class="cover show" role="complementary" aria-label="cover" style="background: ${bgc}">\n      <div class="mask"></div>\n      <div class="cover-main">\x3c!--cover--\x3e</div>\n    </section>\n  `;
    }
    function tree(toc) {
        let tpl = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : '<ul class="app-sub-sidebar">{inner}</ul>';
        if (!toc || !toc.length) {
            return "";
        }
        let innerHTML = "";
        toc.forEach((node => {
            const title = node.title.replace(/(<([^>]+)>)/g, "");
            innerHTML += `<li><a class="section-link" href="${node.slug}" title="${title}">${node.title}</a></li>`;
            if (node.children) {
                innerHTML += tree(node.children, tpl);
            }
        }));
        return tpl.replace("{inner}", innerHTML);
    }
    function helper(className, content) {
        return `<p class="${className}">${content.slice(5).trim()}</p>`;
    }
    function theme(color) {
        return `<style>:root{--theme-color: ${color};}</style>`;
    }
    function genTree(toc, maxLevel) {
        const headlines = [];
        const last = {};
        toc.forEach((headline => {
            const level = headline.level || 1;
            const len = level - 1;
            if (level > maxLevel) {
                return;
            }
            if (last[len]) {
                last[len].children = [ ...last[len].children || [], headline ];
            } else {
                headlines.push(headline);
            }
            last[level] = headline;
        }));
        return headlines;
    }
    let cache$1 = {};
    const re = /[\u2000-\u206F\u2E00-\u2E7F\\'!"#$%&()*+,./:;<=>?@[\]^`{|}~]/g;
    function lower(string) {
        return string.toLowerCase();
    }
    function slugify(str) {
        if (typeof str !== "string") {
            return "";
        }
        let slug = str.trim().replace(/[A-Z]+/g, lower).replace(/<[^>]+>/g, "").replace(re, "").replace(/\s/g, "-").replace(/-+/g, "-").replace(/^(\d)/, "_$1");
        let count = cache$1[slug];
        count = Object.keys(cache$1).includes(slug) ? count + 1 : 0;
        cache$1[slug] = count;
        if (count) {
            slug = slug + "-" + count;
        }
        return slug;
    }
    slugify.clear = function() {
        cache$1 = {};
    };
    var emojiData = {
        baseURL: "https://github.githubassets.com/images/icons/emoji/",
        data: {
            100: "unicode/1f4af.png?v8",
            1234: "unicode/1f522.png?v8",
            "+1": "unicode/1f44d.png?v8",
            "-1": "unicode/1f44e.png?v8",
            "1st_place_medal": "unicode/1f947.png?v8",
            "2nd_place_medal": "unicode/1f948.png?v8",
            "3rd_place_medal": "unicode/1f949.png?v8",
            "8ball": "unicode/1f3b1.png?v8",
            a: "unicode/1f170.png?v8",
            ab: "unicode/1f18e.png?v8",
            abacus: "unicode/1f9ee.png?v8",
            abc: "unicode/1f524.png?v8",
            abcd: "unicode/1f521.png?v8",
            accept: "unicode/1f251.png?v8",
            accessibility: "accessibility.png?v8",
            accordion: "unicode/1fa97.png?v8",
            adhesive_bandage: "unicode/1fa79.png?v8",
            adult: "unicode/1f9d1.png?v8",
            aerial_tramway: "unicode/1f6a1.png?v8",
            afghanistan: "unicode/1f1e6-1f1eb.png?v8",
            airplane: "unicode/2708.png?v8",
            aland_islands: "unicode/1f1e6-1f1fd.png?v8",
            alarm_clock: "unicode/23f0.png?v8",
            albania: "unicode/1f1e6-1f1f1.png?v8",
            alembic: "unicode/2697.png?v8",
            algeria: "unicode/1f1e9-1f1ff.png?v8",
            alien: "unicode/1f47d.png?v8",
            ambulance: "unicode/1f691.png?v8",
            american_samoa: "unicode/1f1e6-1f1f8.png?v8",
            amphora: "unicode/1f3fa.png?v8",
            anatomical_heart: "unicode/1fac0.png?v8",
            anchor: "unicode/2693.png?v8",
            andorra: "unicode/1f1e6-1f1e9.png?v8",
            angel: "unicode/1f47c.png?v8",
            anger: "unicode/1f4a2.png?v8",
            angola: "unicode/1f1e6-1f1f4.png?v8",
            angry: "unicode/1f620.png?v8",
            anguilla: "unicode/1f1e6-1f1ee.png?v8",
            anguished: "unicode/1f627.png?v8",
            ant: "unicode/1f41c.png?v8",
            antarctica: "unicode/1f1e6-1f1f6.png?v8",
            antigua_barbuda: "unicode/1f1e6-1f1ec.png?v8",
            apple: "unicode/1f34e.png?v8",
            aquarius: "unicode/2652.png?v8",
            argentina: "unicode/1f1e6-1f1f7.png?v8",
            aries: "unicode/2648.png?v8",
            armenia: "unicode/1f1e6-1f1f2.png?v8",
            arrow_backward: "unicode/25c0.png?v8",
            arrow_double_down: "unicode/23ec.png?v8",
            arrow_double_up: "unicode/23eb.png?v8",
            arrow_down: "unicode/2b07.png?v8",
            arrow_down_small: "unicode/1f53d.png?v8",
            arrow_forward: "unicode/25b6.png?v8",
            arrow_heading_down: "unicode/2935.png?v8",
            arrow_heading_up: "unicode/2934.png?v8",
            arrow_left: "unicode/2b05.png?v8",
            arrow_lower_left: "unicode/2199.png?v8",
            arrow_lower_right: "unicode/2198.png?v8",
            arrow_right: "unicode/27a1.png?v8",
            arrow_right_hook: "unicode/21aa.png?v8",
            arrow_up: "unicode/2b06.png?v8",
            arrow_up_down: "unicode/2195.png?v8",
            arrow_up_small: "unicode/1f53c.png?v8",
            arrow_upper_left: "unicode/2196.png?v8",
            arrow_upper_right: "unicode/2197.png?v8",
            arrows_clockwise: "unicode/1f503.png?v8",
            arrows_counterclockwise: "unicode/1f504.png?v8",
            art: "unicode/1f3a8.png?v8",
            articulated_lorry: "unicode/1f69b.png?v8",
            artificial_satellite: "unicode/1f6f0.png?v8",
            artist: "unicode/1f9d1-1f3a8.png?v8",
            aruba: "unicode/1f1e6-1f1fc.png?v8",
            ascension_island: "unicode/1f1e6-1f1e8.png?v8",
            asterisk: "unicode/002a-20e3.png?v8",
            astonished: "unicode/1f632.png?v8",
            astronaut: "unicode/1f9d1-1f680.png?v8",
            athletic_shoe: "unicode/1f45f.png?v8",
            atm: "unicode/1f3e7.png?v8",
            atom: "atom.png?v8",
            atom_symbol: "unicode/269b.png?v8",
            australia: "unicode/1f1e6-1f1fa.png?v8",
            austria: "unicode/1f1e6-1f1f9.png?v8",
            auto_rickshaw: "unicode/1f6fa.png?v8",
            avocado: "unicode/1f951.png?v8",
            axe: "unicode/1fa93.png?v8",
            azerbaijan: "unicode/1f1e6-1f1ff.png?v8",
            b: "unicode/1f171.png?v8",
            baby: "unicode/1f476.png?v8",
            baby_bottle: "unicode/1f37c.png?v8",
            baby_chick: "unicode/1f424.png?v8",
            baby_symbol: "unicode/1f6bc.png?v8",
            back: "unicode/1f519.png?v8",
            bacon: "unicode/1f953.png?v8",
            badger: "unicode/1f9a1.png?v8",
            badminton: "unicode/1f3f8.png?v8",
            bagel: "unicode/1f96f.png?v8",
            baggage_claim: "unicode/1f6c4.png?v8",
            baguette_bread: "unicode/1f956.png?v8",
            bahamas: "unicode/1f1e7-1f1f8.png?v8",
            bahrain: "unicode/1f1e7-1f1ed.png?v8",
            balance_scale: "unicode/2696.png?v8",
            bald_man: "unicode/1f468-1f9b2.png?v8",
            bald_woman: "unicode/1f469-1f9b2.png?v8",
            ballet_shoes: "unicode/1fa70.png?v8",
            balloon: "unicode/1f388.png?v8",
            ballot_box: "unicode/1f5f3.png?v8",
            ballot_box_with_check: "unicode/2611.png?v8",
            bamboo: "unicode/1f38d.png?v8",
            banana: "unicode/1f34c.png?v8",
            bangbang: "unicode/203c.png?v8",
            bangladesh: "unicode/1f1e7-1f1e9.png?v8",
            banjo: "unicode/1fa95.png?v8",
            bank: "unicode/1f3e6.png?v8",
            bar_chart: "unicode/1f4ca.png?v8",
            barbados: "unicode/1f1e7-1f1e7.png?v8",
            barber: "unicode/1f488.png?v8",
            baseball: "unicode/26be.png?v8",
            basecamp: "basecamp.png?v8",
            basecampy: "basecampy.png?v8",
            basket: "unicode/1f9fa.png?v8",
            basketball: "unicode/1f3c0.png?v8",
            basketball_man: "unicode/26f9-2642.png?v8",
            basketball_woman: "unicode/26f9-2640.png?v8",
            bat: "unicode/1f987.png?v8",
            bath: "unicode/1f6c0.png?v8",
            bathtub: "unicode/1f6c1.png?v8",
            battery: "unicode/1f50b.png?v8",
            beach_umbrella: "unicode/1f3d6.png?v8",
            beans: "unicode/1fad8.png?v8",
            bear: "unicode/1f43b.png?v8",
            bearded_person: "unicode/1f9d4.png?v8",
            beaver: "unicode/1f9ab.png?v8",
            bed: "unicode/1f6cf.png?v8",
            bee: "unicode/1f41d.png?v8",
            beer: "unicode/1f37a.png?v8",
            beers: "unicode/1f37b.png?v8",
            beetle: "unicode/1fab2.png?v8",
            beginner: "unicode/1f530.png?v8",
            belarus: "unicode/1f1e7-1f1fe.png?v8",
            belgium: "unicode/1f1e7-1f1ea.png?v8",
            belize: "unicode/1f1e7-1f1ff.png?v8",
            bell: "unicode/1f514.png?v8",
            bell_pepper: "unicode/1fad1.png?v8",
            bellhop_bell: "unicode/1f6ce.png?v8",
            benin: "unicode/1f1e7-1f1ef.png?v8",
            bento: "unicode/1f371.png?v8",
            bermuda: "unicode/1f1e7-1f1f2.png?v8",
            beverage_box: "unicode/1f9c3.png?v8",
            bhutan: "unicode/1f1e7-1f1f9.png?v8",
            bicyclist: "unicode/1f6b4.png?v8",
            bike: "unicode/1f6b2.png?v8",
            biking_man: "unicode/1f6b4-2642.png?v8",
            biking_woman: "unicode/1f6b4-2640.png?v8",
            bikini: "unicode/1f459.png?v8",
            billed_cap: "unicode/1f9e2.png?v8",
            biohazard: "unicode/2623.png?v8",
            bird: "unicode/1f426.png?v8",
            birthday: "unicode/1f382.png?v8",
            bison: "unicode/1f9ac.png?v8",
            biting_lip: "unicode/1fae6.png?v8",
            black_bird: "unicode/1f426-2b1b.png?v8",
            black_cat: "unicode/1f408-2b1b.png?v8",
            black_circle: "unicode/26ab.png?v8",
            black_flag: "unicode/1f3f4.png?v8",
            black_heart: "unicode/1f5a4.png?v8",
            black_joker: "unicode/1f0cf.png?v8",
            black_large_square: "unicode/2b1b.png?v8",
            black_medium_small_square: "unicode/25fe.png?v8",
            black_medium_square: "unicode/25fc.png?v8",
            black_nib: "unicode/2712.png?v8",
            black_small_square: "unicode/25aa.png?v8",
            black_square_button: "unicode/1f532.png?v8",
            blond_haired_man: "unicode/1f471-2642.png?v8",
            blond_haired_person: "unicode/1f471.png?v8",
            blond_haired_woman: "unicode/1f471-2640.png?v8",
            blonde_woman: "unicode/1f471-2640.png?v8",
            blossom: "unicode/1f33c.png?v8",
            blowfish: "unicode/1f421.png?v8",
            blue_book: "unicode/1f4d8.png?v8",
            blue_car: "unicode/1f699.png?v8",
            blue_heart: "unicode/1f499.png?v8",
            blue_square: "unicode/1f7e6.png?v8",
            blueberries: "unicode/1fad0.png?v8",
            blush: "unicode/1f60a.png?v8",
            boar: "unicode/1f417.png?v8",
            boat: "unicode/26f5.png?v8",
            bolivia: "unicode/1f1e7-1f1f4.png?v8",
            bomb: "unicode/1f4a3.png?v8",
            bone: "unicode/1f9b4.png?v8",
            book: "unicode/1f4d6.png?v8",
            bookmark: "unicode/1f516.png?v8",
            bookmark_tabs: "unicode/1f4d1.png?v8",
            books: "unicode/1f4da.png?v8",
            boom: "unicode/1f4a5.png?v8",
            boomerang: "unicode/1fa83.png?v8",
            boot: "unicode/1f462.png?v8",
            bosnia_herzegovina: "unicode/1f1e7-1f1e6.png?v8",
            botswana: "unicode/1f1e7-1f1fc.png?v8",
            bouncing_ball_man: "unicode/26f9-2642.png?v8",
            bouncing_ball_person: "unicode/26f9.png?v8",
            bouncing_ball_woman: "unicode/26f9-2640.png?v8",
            bouquet: "unicode/1f490.png?v8",
            bouvet_island: "unicode/1f1e7-1f1fb.png?v8",
            bow: "unicode/1f647.png?v8",
            bow_and_arrow: "unicode/1f3f9.png?v8",
            bowing_man: "unicode/1f647-2642.png?v8",
            bowing_woman: "unicode/1f647-2640.png?v8",
            bowl_with_spoon: "unicode/1f963.png?v8",
            bowling: "unicode/1f3b3.png?v8",
            bowtie: "bowtie.png?v8",
            boxing_glove: "unicode/1f94a.png?v8",
            boy: "unicode/1f466.png?v8",
            brain: "unicode/1f9e0.png?v8",
            brazil: "unicode/1f1e7-1f1f7.png?v8",
            bread: "unicode/1f35e.png?v8",
            breast_feeding: "unicode/1f931.png?v8",
            bricks: "unicode/1f9f1.png?v8",
            bride_with_veil: "unicode/1f470-2640.png?v8",
            bridge_at_night: "unicode/1f309.png?v8",
            briefcase: "unicode/1f4bc.png?v8",
            british_indian_ocean_territory: "unicode/1f1ee-1f1f4.png?v8",
            british_virgin_islands: "unicode/1f1fb-1f1ec.png?v8",
            broccoli: "unicode/1f966.png?v8",
            broken_heart: "unicode/1f494.png?v8",
            broom: "unicode/1f9f9.png?v8",
            brown_circle: "unicode/1f7e4.png?v8",
            brown_heart: "unicode/1f90e.png?v8",
            brown_square: "unicode/1f7eb.png?v8",
            brunei: "unicode/1f1e7-1f1f3.png?v8",
            bubble_tea: "unicode/1f9cb.png?v8",
            bubbles: "unicode/1fae7.png?v8",
            bucket: "unicode/1faa3.png?v8",
            bug: "unicode/1f41b.png?v8",
            building_construction: "unicode/1f3d7.png?v8",
            bulb: "unicode/1f4a1.png?v8",
            bulgaria: "unicode/1f1e7-1f1ec.png?v8",
            bullettrain_front: "unicode/1f685.png?v8",
            bullettrain_side: "unicode/1f684.png?v8",
            burkina_faso: "unicode/1f1e7-1f1eb.png?v8",
            burrito: "unicode/1f32f.png?v8",
            burundi: "unicode/1f1e7-1f1ee.png?v8",
            bus: "unicode/1f68c.png?v8",
            business_suit_levitating: "unicode/1f574.png?v8",
            busstop: "unicode/1f68f.png?v8",
            bust_in_silhouette: "unicode/1f464.png?v8",
            busts_in_silhouette: "unicode/1f465.png?v8",
            butter: "unicode/1f9c8.png?v8",
            butterfly: "unicode/1f98b.png?v8",
            cactus: "unicode/1f335.png?v8",
            cake: "unicode/1f370.png?v8",
            calendar: "unicode/1f4c6.png?v8",
            call_me_hand: "unicode/1f919.png?v8",
            calling: "unicode/1f4f2.png?v8",
            cambodia: "unicode/1f1f0-1f1ed.png?v8",
            camel: "unicode/1f42b.png?v8",
            camera: "unicode/1f4f7.png?v8",
            camera_flash: "unicode/1f4f8.png?v8",
            cameroon: "unicode/1f1e8-1f1f2.png?v8",
            camping: "unicode/1f3d5.png?v8",
            canada: "unicode/1f1e8-1f1e6.png?v8",
            canary_islands: "unicode/1f1ee-1f1e8.png?v8",
            cancer: "unicode/264b.png?v8",
            candle: "unicode/1f56f.png?v8",
            candy: "unicode/1f36c.png?v8",
            canned_food: "unicode/1f96b.png?v8",
            canoe: "unicode/1f6f6.png?v8",
            cape_verde: "unicode/1f1e8-1f1fb.png?v8",
            capital_abcd: "unicode/1f520.png?v8",
            capricorn: "unicode/2651.png?v8",
            car: "unicode/1f697.png?v8",
            card_file_box: "unicode/1f5c3.png?v8",
            card_index: "unicode/1f4c7.png?v8",
            card_index_dividers: "unicode/1f5c2.png?v8",
            caribbean_netherlands: "unicode/1f1e7-1f1f6.png?v8",
            carousel_horse: "unicode/1f3a0.png?v8",
            carpentry_saw: "unicode/1fa9a.png?v8",
            carrot: "unicode/1f955.png?v8",
            cartwheeling: "unicode/1f938.png?v8",
            cat: "unicode/1f431.png?v8",
            cat2: "unicode/1f408.png?v8",
            cayman_islands: "unicode/1f1f0-1f1fe.png?v8",
            cd: "unicode/1f4bf.png?v8",
            central_african_republic: "unicode/1f1e8-1f1eb.png?v8",
            ceuta_melilla: "unicode/1f1ea-1f1e6.png?v8",
            chad: "unicode/1f1f9-1f1e9.png?v8",
            chains: "unicode/26d3.png?v8",
            chair: "unicode/1fa91.png?v8",
            champagne: "unicode/1f37e.png?v8",
            chart: "unicode/1f4b9.png?v8",
            chart_with_downwards_trend: "unicode/1f4c9.png?v8",
            chart_with_upwards_trend: "unicode/1f4c8.png?v8",
            checkered_flag: "unicode/1f3c1.png?v8",
            cheese: "unicode/1f9c0.png?v8",
            cherries: "unicode/1f352.png?v8",
            cherry_blossom: "unicode/1f338.png?v8",
            chess_pawn: "unicode/265f.png?v8",
            chestnut: "unicode/1f330.png?v8",
            chicken: "unicode/1f414.png?v8",
            child: "unicode/1f9d2.png?v8",
            children_crossing: "unicode/1f6b8.png?v8",
            chile: "unicode/1f1e8-1f1f1.png?v8",
            chipmunk: "unicode/1f43f.png?v8",
            chocolate_bar: "unicode/1f36b.png?v8",
            chopsticks: "unicode/1f962.png?v8",
            christmas_island: "unicode/1f1e8-1f1fd.png?v8",
            christmas_tree: "unicode/1f384.png?v8",
            church: "unicode/26ea.png?v8",
            cinema: "unicode/1f3a6.png?v8",
            circus_tent: "unicode/1f3aa.png?v8",
            city_sunrise: "unicode/1f307.png?v8",
            city_sunset: "unicode/1f306.png?v8",
            cityscape: "unicode/1f3d9.png?v8",
            cl: "unicode/1f191.png?v8",
            clamp: "unicode/1f5dc.png?v8",
            clap: "unicode/1f44f.png?v8",
            clapper: "unicode/1f3ac.png?v8",
            classical_building: "unicode/1f3db.png?v8",
            climbing: "unicode/1f9d7.png?v8",
            climbing_man: "unicode/1f9d7-2642.png?v8",
            climbing_woman: "unicode/1f9d7-2640.png?v8",
            clinking_glasses: "unicode/1f942.png?v8",
            clipboard: "unicode/1f4cb.png?v8",
            clipperton_island: "unicode/1f1e8-1f1f5.png?v8",
            clock1: "unicode/1f550.png?v8",
            clock10: "unicode/1f559.png?v8",
            clock1030: "unicode/1f565.png?v8",
            clock11: "unicode/1f55a.png?v8",
            clock1130: "unicode/1f566.png?v8",
            clock12: "unicode/1f55b.png?v8",
            clock1230: "unicode/1f567.png?v8",
            clock130: "unicode/1f55c.png?v8",
            clock2: "unicode/1f551.png?v8",
            clock230: "unicode/1f55d.png?v8",
            clock3: "unicode/1f552.png?v8",
            clock330: "unicode/1f55e.png?v8",
            clock4: "unicode/1f553.png?v8",
            clock430: "unicode/1f55f.png?v8",
            clock5: "unicode/1f554.png?v8",
            clock530: "unicode/1f560.png?v8",
            clock6: "unicode/1f555.png?v8",
            clock630: "unicode/1f561.png?v8",
            clock7: "unicode/1f556.png?v8",
            clock730: "unicode/1f562.png?v8",
            clock8: "unicode/1f557.png?v8",
            clock830: "unicode/1f563.png?v8",
            clock9: "unicode/1f558.png?v8",
            clock930: "unicode/1f564.png?v8",
            closed_book: "unicode/1f4d5.png?v8",
            closed_lock_with_key: "unicode/1f510.png?v8",
            closed_umbrella: "unicode/1f302.png?v8",
            cloud: "unicode/2601.png?v8",
            cloud_with_lightning: "unicode/1f329.png?v8",
            cloud_with_lightning_and_rain: "unicode/26c8.png?v8",
            cloud_with_rain: "unicode/1f327.png?v8",
            cloud_with_snow: "unicode/1f328.png?v8",
            clown_face: "unicode/1f921.png?v8",
            clubs: "unicode/2663.png?v8",
            cn: "unicode/1f1e8-1f1f3.png?v8",
            coat: "unicode/1f9e5.png?v8",
            cockroach: "unicode/1fab3.png?v8",
            cocktail: "unicode/1f378.png?v8",
            coconut: "unicode/1f965.png?v8",
            cocos_islands: "unicode/1f1e8-1f1e8.png?v8",
            coffee: "unicode/2615.png?v8",
            coffin: "unicode/26b0.png?v8",
            coin: "unicode/1fa99.png?v8",
            cold_face: "unicode/1f976.png?v8",
            cold_sweat: "unicode/1f630.png?v8",
            collision: "unicode/1f4a5.png?v8",
            colombia: "unicode/1f1e8-1f1f4.png?v8",
            comet: "unicode/2604.png?v8",
            comoros: "unicode/1f1f0-1f1f2.png?v8",
            compass: "unicode/1f9ed.png?v8",
            computer: "unicode/1f4bb.png?v8",
            computer_mouse: "unicode/1f5b1.png?v8",
            confetti_ball: "unicode/1f38a.png?v8",
            confounded: "unicode/1f616.png?v8",
            confused: "unicode/1f615.png?v8",
            congo_brazzaville: "unicode/1f1e8-1f1ec.png?v8",
            congo_kinshasa: "unicode/1f1e8-1f1e9.png?v8",
            congratulations: "unicode/3297.png?v8",
            construction: "unicode/1f6a7.png?v8",
            construction_worker: "unicode/1f477.png?v8",
            construction_worker_man: "unicode/1f477-2642.png?v8",
            construction_worker_woman: "unicode/1f477-2640.png?v8",
            control_knobs: "unicode/1f39b.png?v8",
            convenience_store: "unicode/1f3ea.png?v8",
            cook: "unicode/1f9d1-1f373.png?v8",
            cook_islands: "unicode/1f1e8-1f1f0.png?v8",
            cookie: "unicode/1f36a.png?v8",
            cool: "unicode/1f192.png?v8",
            cop: "unicode/1f46e.png?v8",
            copyright: "unicode/00a9.png?v8",
            coral: "unicode/1fab8.png?v8",
            corn: "unicode/1f33d.png?v8",
            costa_rica: "unicode/1f1e8-1f1f7.png?v8",
            cote_divoire: "unicode/1f1e8-1f1ee.png?v8",
            couch_and_lamp: "unicode/1f6cb.png?v8",
            couple: "unicode/1f46b.png?v8",
            couple_with_heart: "unicode/1f491.png?v8",
            couple_with_heart_man_man: "unicode/1f468-2764-1f468.png?v8",
            couple_with_heart_woman_man: "unicode/1f469-2764-1f468.png?v8",
            couple_with_heart_woman_woman: "unicode/1f469-2764-1f469.png?v8",
            couplekiss: "unicode/1f48f.png?v8",
            couplekiss_man_man: "unicode/1f468-2764-1f48b-1f468.png?v8",
            couplekiss_man_woman: "unicode/1f469-2764-1f48b-1f468.png?v8",
            couplekiss_woman_woman: "unicode/1f469-2764-1f48b-1f469.png?v8",
            cow: "unicode/1f42e.png?v8",
            cow2: "unicode/1f404.png?v8",
            cowboy_hat_face: "unicode/1f920.png?v8",
            crab: "unicode/1f980.png?v8",
            crayon: "unicode/1f58d.png?v8",
            credit_card: "unicode/1f4b3.png?v8",
            crescent_moon: "unicode/1f319.png?v8",
            cricket: "unicode/1f997.png?v8",
            cricket_game: "unicode/1f3cf.png?v8",
            croatia: "unicode/1f1ed-1f1f7.png?v8",
            crocodile: "unicode/1f40a.png?v8",
            croissant: "unicode/1f950.png?v8",
            crossed_fingers: "unicode/1f91e.png?v8",
            crossed_flags: "unicode/1f38c.png?v8",
            crossed_swords: "unicode/2694.png?v8",
            crown: "unicode/1f451.png?v8",
            crutch: "unicode/1fa7c.png?v8",
            cry: "unicode/1f622.png?v8",
            crying_cat_face: "unicode/1f63f.png?v8",
            crystal_ball: "unicode/1f52e.png?v8",
            cuba: "unicode/1f1e8-1f1fa.png?v8",
            cucumber: "unicode/1f952.png?v8",
            cup_with_straw: "unicode/1f964.png?v8",
            cupcake: "unicode/1f9c1.png?v8",
            cupid: "unicode/1f498.png?v8",
            curacao: "unicode/1f1e8-1f1fc.png?v8",
            curling_stone: "unicode/1f94c.png?v8",
            curly_haired_man: "unicode/1f468-1f9b1.png?v8",
            curly_haired_woman: "unicode/1f469-1f9b1.png?v8",
            curly_loop: "unicode/27b0.png?v8",
            currency_exchange: "unicode/1f4b1.png?v8",
            curry: "unicode/1f35b.png?v8",
            cursing_face: "unicode/1f92c.png?v8",
            custard: "unicode/1f36e.png?v8",
            customs: "unicode/1f6c3.png?v8",
            cut_of_meat: "unicode/1f969.png?v8",
            cyclone: "unicode/1f300.png?v8",
            cyprus: "unicode/1f1e8-1f1fe.png?v8",
            czech_republic: "unicode/1f1e8-1f1ff.png?v8",
            dagger: "unicode/1f5e1.png?v8",
            dancer: "unicode/1f483.png?v8",
            dancers: "unicode/1f46f.png?v8",
            dancing_men: "unicode/1f46f-2642.png?v8",
            dancing_women: "unicode/1f46f-2640.png?v8",
            dango: "unicode/1f361.png?v8",
            dark_sunglasses: "unicode/1f576.png?v8",
            dart: "unicode/1f3af.png?v8",
            dash: "unicode/1f4a8.png?v8",
            date: "unicode/1f4c5.png?v8",
            de: "unicode/1f1e9-1f1ea.png?v8",
            deaf_man: "unicode/1f9cf-2642.png?v8",
            deaf_person: "unicode/1f9cf.png?v8",
            deaf_woman: "unicode/1f9cf-2640.png?v8",
            deciduous_tree: "unicode/1f333.png?v8",
            deer: "unicode/1f98c.png?v8",
            denmark: "unicode/1f1e9-1f1f0.png?v8",
            department_store: "unicode/1f3ec.png?v8",
            dependabot: "dependabot.png?v8",
            derelict_house: "unicode/1f3da.png?v8",
            desert: "unicode/1f3dc.png?v8",
            desert_island: "unicode/1f3dd.png?v8",
            desktop_computer: "unicode/1f5a5.png?v8",
            detective: "unicode/1f575.png?v8",
            diamond_shape_with_a_dot_inside: "unicode/1f4a0.png?v8",
            diamonds: "unicode/2666.png?v8",
            diego_garcia: "unicode/1f1e9-1f1ec.png?v8",
            disappointed: "unicode/1f61e.png?v8",
            disappointed_relieved: "unicode/1f625.png?v8",
            disguised_face: "unicode/1f978.png?v8",
            diving_mask: "unicode/1f93f.png?v8",
            diya_lamp: "unicode/1fa94.png?v8",
            dizzy: "unicode/1f4ab.png?v8",
            dizzy_face: "unicode/1f635.png?v8",
            djibouti: "unicode/1f1e9-1f1ef.png?v8",
            dna: "unicode/1f9ec.png?v8",
            do_not_litter: "unicode/1f6af.png?v8",
            dodo: "unicode/1f9a4.png?v8",
            dog: "unicode/1f436.png?v8",
            dog2: "unicode/1f415.png?v8",
            dollar: "unicode/1f4b5.png?v8",
            dolls: "unicode/1f38e.png?v8",
            dolphin: "unicode/1f42c.png?v8",
            dominica: "unicode/1f1e9-1f1f2.png?v8",
            dominican_republic: "unicode/1f1e9-1f1f4.png?v8",
            donkey: "unicode/1facf.png?v8",
            door: "unicode/1f6aa.png?v8",
            dotted_line_face: "unicode/1fae5.png?v8",
            doughnut: "unicode/1f369.png?v8",
            dove: "unicode/1f54a.png?v8",
            dragon: "unicode/1f409.png?v8",
            dragon_face: "unicode/1f432.png?v8",
            dress: "unicode/1f457.png?v8",
            dromedary_camel: "unicode/1f42a.png?v8",
            drooling_face: "unicode/1f924.png?v8",
            drop_of_blood: "unicode/1fa78.png?v8",
            droplet: "unicode/1f4a7.png?v8",
            drum: "unicode/1f941.png?v8",
            duck: "unicode/1f986.png?v8",
            dumpling: "unicode/1f95f.png?v8",
            dvd: "unicode/1f4c0.png?v8",
            "e-mail": "unicode/1f4e7.png?v8",
            eagle: "unicode/1f985.png?v8",
            ear: "unicode/1f442.png?v8",
            ear_of_rice: "unicode/1f33e.png?v8",
            ear_with_hearing_aid: "unicode/1f9bb.png?v8",
            earth_africa: "unicode/1f30d.png?v8",
            earth_americas: "unicode/1f30e.png?v8",
            earth_asia: "unicode/1f30f.png?v8",
            ecuador: "unicode/1f1ea-1f1e8.png?v8",
            egg: "unicode/1f95a.png?v8",
            eggplant: "unicode/1f346.png?v8",
            egypt: "unicode/1f1ea-1f1ec.png?v8",
            eight: "unicode/0038-20e3.png?v8",
            eight_pointed_black_star: "unicode/2734.png?v8",
            eight_spoked_asterisk: "unicode/2733.png?v8",
            eject_button: "unicode/23cf.png?v8",
            el_salvador: "unicode/1f1f8-1f1fb.png?v8",
            electric_plug: "unicode/1f50c.png?v8",
            electron: "electron.png?v8",
            elephant: "unicode/1f418.png?v8",
            elevator: "unicode/1f6d7.png?v8",
            elf: "unicode/1f9dd.png?v8",
            elf_man: "unicode/1f9dd-2642.png?v8",
            elf_woman: "unicode/1f9dd-2640.png?v8",
            email: "unicode/1f4e7.png?v8",
            empty_nest: "unicode/1fab9.png?v8",
            end: "unicode/1f51a.png?v8",
            england: "unicode/1f3f4-e0067-e0062-e0065-e006e-e0067-e007f.png?v8",
            envelope: "unicode/2709.png?v8",
            envelope_with_arrow: "unicode/1f4e9.png?v8",
            equatorial_guinea: "unicode/1f1ec-1f1f6.png?v8",
            eritrea: "unicode/1f1ea-1f1f7.png?v8",
            es: "unicode/1f1ea-1f1f8.png?v8",
            estonia: "unicode/1f1ea-1f1ea.png?v8",
            ethiopia: "unicode/1f1ea-1f1f9.png?v8",
            eu: "unicode/1f1ea-1f1fa.png?v8",
            euro: "unicode/1f4b6.png?v8",
            european_castle: "unicode/1f3f0.png?v8",
            european_post_office: "unicode/1f3e4.png?v8",
            european_union: "unicode/1f1ea-1f1fa.png?v8",
            evergreen_tree: "unicode/1f332.png?v8",
            exclamation: "unicode/2757.png?v8",
            exploding_head: "unicode/1f92f.png?v8",
            expressionless: "unicode/1f611.png?v8",
            eye: "unicode/1f441.png?v8",
            eye_speech_bubble: "unicode/1f441-1f5e8.png?v8",
            eyeglasses: "unicode/1f453.png?v8",
            eyes: "unicode/1f440.png?v8",
            face_exhaling: "unicode/1f62e-1f4a8.png?v8",
            face_holding_back_tears: "unicode/1f979.png?v8",
            face_in_clouds: "unicode/1f636-1f32b.png?v8",
            face_with_diagonal_mouth: "unicode/1fae4.png?v8",
            face_with_head_bandage: "unicode/1f915.png?v8",
            face_with_open_eyes_and_hand_over_mouth: "unicode/1fae2.png?v8",
            face_with_peeking_eye: "unicode/1fae3.png?v8",
            face_with_spiral_eyes: "unicode/1f635-1f4ab.png?v8",
            face_with_thermometer: "unicode/1f912.png?v8",
            facepalm: "unicode/1f926.png?v8",
            facepunch: "unicode/1f44a.png?v8",
            factory: "unicode/1f3ed.png?v8",
            factory_worker: "unicode/1f9d1-1f3ed.png?v8",
            fairy: "unicode/1f9da.png?v8",
            fairy_man: "unicode/1f9da-2642.png?v8",
            fairy_woman: "unicode/1f9da-2640.png?v8",
            falafel: "unicode/1f9c6.png?v8",
            falkland_islands: "unicode/1f1eb-1f1f0.png?v8",
            fallen_leaf: "unicode/1f342.png?v8",
            family: "unicode/1f46a.png?v8",
            family_man_boy: "unicode/1f468-1f466.png?v8",
            family_man_boy_boy: "unicode/1f468-1f466-1f466.png?v8",
            family_man_girl: "unicode/1f468-1f467.png?v8",
            family_man_girl_boy: "unicode/1f468-1f467-1f466.png?v8",
            family_man_girl_girl: "unicode/1f468-1f467-1f467.png?v8",
            family_man_man_boy: "unicode/1f468-1f468-1f466.png?v8",
            family_man_man_boy_boy: "unicode/1f468-1f468-1f466-1f466.png?v8",
            family_man_man_girl: "unicode/1f468-1f468-1f467.png?v8",
            family_man_man_girl_boy: "unicode/1f468-1f468-1f467-1f466.png?v8",
            family_man_man_girl_girl: "unicode/1f468-1f468-1f467-1f467.png?v8",
            family_man_woman_boy: "unicode/1f468-1f469-1f466.png?v8",
            family_man_woman_boy_boy: "unicode/1f468-1f469-1f466-1f466.png?v8",
            family_man_woman_girl: "unicode/1f468-1f469-1f467.png?v8",
            family_man_woman_girl_boy: "unicode/1f468-1f469-1f467-1f466.png?v8",
            family_man_woman_girl_girl: "unicode/1f468-1f469-1f467-1f467.png?v8",
            family_woman_boy: "unicode/1f469-1f466.png?v8",
            family_woman_boy_boy: "unicode/1f469-1f466-1f466.png?v8",
            family_woman_girl: "unicode/1f469-1f467.png?v8",
            family_woman_girl_boy: "unicode/1f469-1f467-1f466.png?v8",
            family_woman_girl_girl: "unicode/1f469-1f467-1f467.png?v8",
            family_woman_woman_boy: "unicode/1f469-1f469-1f466.png?v8",
            family_woman_woman_boy_boy: "unicode/1f469-1f469-1f466-1f466.png?v8",
            family_woman_woman_girl: "unicode/1f469-1f469-1f467.png?v8",
            family_woman_woman_girl_boy: "unicode/1f469-1f469-1f467-1f466.png?v8",
            family_woman_woman_girl_girl: "unicode/1f469-1f469-1f467-1f467.png?v8",
            farmer: "unicode/1f9d1-1f33e.png?v8",
            faroe_islands: "unicode/1f1eb-1f1f4.png?v8",
            fast_forward: "unicode/23e9.png?v8",
            fax: "unicode/1f4e0.png?v8",
            fearful: "unicode/1f628.png?v8",
            feather: "unicode/1fab6.png?v8",
            feelsgood: "feelsgood.png?v8",
            feet: "unicode/1f43e.png?v8",
            female_detective: "unicode/1f575-2640.png?v8",
            female_sign: "unicode/2640.png?v8",
            ferris_wheel: "unicode/1f3a1.png?v8",
            ferry: "unicode/26f4.png?v8",
            field_hockey: "unicode/1f3d1.png?v8",
            fiji: "unicode/1f1eb-1f1ef.png?v8",
            file_cabinet: "unicode/1f5c4.png?v8",
            file_folder: "unicode/1f4c1.png?v8",
            film_projector: "unicode/1f4fd.png?v8",
            film_strip: "unicode/1f39e.png?v8",
            finland: "unicode/1f1eb-1f1ee.png?v8",
            finnadie: "finnadie.png?v8",
            fire: "unicode/1f525.png?v8",
            fire_engine: "unicode/1f692.png?v8",
            fire_extinguisher: "unicode/1f9ef.png?v8",
            firecracker: "unicode/1f9e8.png?v8",
            firefighter: "unicode/1f9d1-1f692.png?v8",
            fireworks: "unicode/1f386.png?v8",
            first_quarter_moon: "unicode/1f313.png?v8",
            first_quarter_moon_with_face: "unicode/1f31b.png?v8",
            fish: "unicode/1f41f.png?v8",
            fish_cake: "unicode/1f365.png?v8",
            fishing_pole_and_fish: "unicode/1f3a3.png?v8",
            fishsticks: "fishsticks.png?v8",
            fist: "unicode/270a.png?v8",
            fist_left: "unicode/1f91b.png?v8",
            fist_oncoming: "unicode/1f44a.png?v8",
            fist_raised: "unicode/270a.png?v8",
            fist_right: "unicode/1f91c.png?v8",
            five: "unicode/0035-20e3.png?v8",
            flags: "unicode/1f38f.png?v8",
            flamingo: "unicode/1f9a9.png?v8",
            flashlight: "unicode/1f526.png?v8",
            flat_shoe: "unicode/1f97f.png?v8",
            flatbread: "unicode/1fad3.png?v8",
            fleur_de_lis: "unicode/269c.png?v8",
            flight_arrival: "unicode/1f6ec.png?v8",
            flight_departure: "unicode/1f6eb.png?v8",
            flipper: "unicode/1f42c.png?v8",
            floppy_disk: "unicode/1f4be.png?v8",
            flower_playing_cards: "unicode/1f3b4.png?v8",
            flushed: "unicode/1f633.png?v8",
            flute: "unicode/1fa88.png?v8",
            fly: "unicode/1fab0.png?v8",
            flying_disc: "unicode/1f94f.png?v8",
            flying_saucer: "unicode/1f6f8.png?v8",
            fog: "unicode/1f32b.png?v8",
            foggy: "unicode/1f301.png?v8",
            folding_hand_fan: "unicode/1faad.png?v8",
            fondue: "unicode/1fad5.png?v8",
            foot: "unicode/1f9b6.png?v8",
            football: "unicode/1f3c8.png?v8",
            footprints: "unicode/1f463.png?v8",
            fork_and_knife: "unicode/1f374.png?v8",
            fortune_cookie: "unicode/1f960.png?v8",
            fountain: "unicode/26f2.png?v8",
            fountain_pen: "unicode/1f58b.png?v8",
            four: "unicode/0034-20e3.png?v8",
            four_leaf_clover: "unicode/1f340.png?v8",
            fox_face: "unicode/1f98a.png?v8",
            fr: "unicode/1f1eb-1f1f7.png?v8",
            framed_picture: "unicode/1f5bc.png?v8",
            free: "unicode/1f193.png?v8",
            french_guiana: "unicode/1f1ec-1f1eb.png?v8",
            french_polynesia: "unicode/1f1f5-1f1eb.png?v8",
            french_southern_territories: "unicode/1f1f9-1f1eb.png?v8",
            fried_egg: "unicode/1f373.png?v8",
            fried_shrimp: "unicode/1f364.png?v8",
            fries: "unicode/1f35f.png?v8",
            frog: "unicode/1f438.png?v8",
            frowning: "unicode/1f626.png?v8",
            frowning_face: "unicode/2639.png?v8",
            frowning_man: "unicode/1f64d-2642.png?v8",
            frowning_person: "unicode/1f64d.png?v8",
            frowning_woman: "unicode/1f64d-2640.png?v8",
            fu: "unicode/1f595.png?v8",
            fuelpump: "unicode/26fd.png?v8",
            full_moon: "unicode/1f315.png?v8",
            full_moon_with_face: "unicode/1f31d.png?v8",
            funeral_urn: "unicode/26b1.png?v8",
            gabon: "unicode/1f1ec-1f1e6.png?v8",
            gambia: "unicode/1f1ec-1f1f2.png?v8",
            game_die: "unicode/1f3b2.png?v8",
            garlic: "unicode/1f9c4.png?v8",
            gb: "unicode/1f1ec-1f1e7.png?v8",
            gear: "unicode/2699.png?v8",
            gem: "unicode/1f48e.png?v8",
            gemini: "unicode/264a.png?v8",
            genie: "unicode/1f9de.png?v8",
            genie_man: "unicode/1f9de-2642.png?v8",
            genie_woman: "unicode/1f9de-2640.png?v8",
            georgia: "unicode/1f1ec-1f1ea.png?v8",
            ghana: "unicode/1f1ec-1f1ed.png?v8",
            ghost: "unicode/1f47b.png?v8",
            gibraltar: "unicode/1f1ec-1f1ee.png?v8",
            gift: "unicode/1f381.png?v8",
            gift_heart: "unicode/1f49d.png?v8",
            ginger_root: "unicode/1fada.png?v8",
            giraffe: "unicode/1f992.png?v8",
            girl: "unicode/1f467.png?v8",
            globe_with_meridians: "unicode/1f310.png?v8",
            gloves: "unicode/1f9e4.png?v8",
            goal_net: "unicode/1f945.png?v8",
            goat: "unicode/1f410.png?v8",
            goberserk: "goberserk.png?v8",
            godmode: "godmode.png?v8",
            goggles: "unicode/1f97d.png?v8",
            golf: "unicode/26f3.png?v8",
            golfing: "unicode/1f3cc.png?v8",
            golfing_man: "unicode/1f3cc-2642.png?v8",
            golfing_woman: "unicode/1f3cc-2640.png?v8",
            goose: "unicode/1fabf.png?v8",
            gorilla: "unicode/1f98d.png?v8",
            grapes: "unicode/1f347.png?v8",
            greece: "unicode/1f1ec-1f1f7.png?v8",
            green_apple: "unicode/1f34f.png?v8",
            green_book: "unicode/1f4d7.png?v8",
            green_circle: "unicode/1f7e2.png?v8",
            green_heart: "unicode/1f49a.png?v8",
            green_salad: "unicode/1f957.png?v8",
            green_square: "unicode/1f7e9.png?v8",
            greenland: "unicode/1f1ec-1f1f1.png?v8",
            grenada: "unicode/1f1ec-1f1e9.png?v8",
            grey_exclamation: "unicode/2755.png?v8",
            grey_heart: "unicode/1fa76.png?v8",
            grey_question: "unicode/2754.png?v8",
            grimacing: "unicode/1f62c.png?v8",
            grin: "unicode/1f601.png?v8",
            grinning: "unicode/1f600.png?v8",
            guadeloupe: "unicode/1f1ec-1f1f5.png?v8",
            guam: "unicode/1f1ec-1f1fa.png?v8",
            guard: "unicode/1f482.png?v8",
            guardsman: "unicode/1f482-2642.png?v8",
            guardswoman: "unicode/1f482-2640.png?v8",
            guatemala: "unicode/1f1ec-1f1f9.png?v8",
            guernsey: "unicode/1f1ec-1f1ec.png?v8",
            guide_dog: "unicode/1f9ae.png?v8",
            guinea: "unicode/1f1ec-1f1f3.png?v8",
            guinea_bissau: "unicode/1f1ec-1f1fc.png?v8",
            guitar: "unicode/1f3b8.png?v8",
            gun: "unicode/1f52b.png?v8",
            guyana: "unicode/1f1ec-1f1fe.png?v8",
            hair_pick: "unicode/1faae.png?v8",
            haircut: "unicode/1f487.png?v8",
            haircut_man: "unicode/1f487-2642.png?v8",
            haircut_woman: "unicode/1f487-2640.png?v8",
            haiti: "unicode/1f1ed-1f1f9.png?v8",
            hamburger: "unicode/1f354.png?v8",
            hammer: "unicode/1f528.png?v8",
            hammer_and_pick: "unicode/2692.png?v8",
            hammer_and_wrench: "unicode/1f6e0.png?v8",
            hamsa: "unicode/1faac.png?v8",
            hamster: "unicode/1f439.png?v8",
            hand: "unicode/270b.png?v8",
            hand_over_mouth: "unicode/1f92d.png?v8",
            hand_with_index_finger_and_thumb_crossed: "unicode/1faf0.png?v8",
            handbag: "unicode/1f45c.png?v8",
            handball_person: "unicode/1f93e.png?v8",
            handshake: "unicode/1f91d.png?v8",
            hankey: "unicode/1f4a9.png?v8",
            hash: "unicode/0023-20e3.png?v8",
            hatched_chick: "unicode/1f425.png?v8",
            hatching_chick: "unicode/1f423.png?v8",
            headphones: "unicode/1f3a7.png?v8",
            headstone: "unicode/1faa6.png?v8",
            health_worker: "unicode/1f9d1-2695.png?v8",
            hear_no_evil: "unicode/1f649.png?v8",
            heard_mcdonald_islands: "unicode/1f1ed-1f1f2.png?v8",
            heart: "unicode/2764.png?v8",
            heart_decoration: "unicode/1f49f.png?v8",
            heart_eyes: "unicode/1f60d.png?v8",
            heart_eyes_cat: "unicode/1f63b.png?v8",
            heart_hands: "unicode/1faf6.png?v8",
            heart_on_fire: "unicode/2764-1f525.png?v8",
            heartbeat: "unicode/1f493.png?v8",
            heartpulse: "unicode/1f497.png?v8",
            hearts: "unicode/2665.png?v8",
            heavy_check_mark: "unicode/2714.png?v8",
            heavy_division_sign: "unicode/2797.png?v8",
            heavy_dollar_sign: "unicode/1f4b2.png?v8",
            heavy_equals_sign: "unicode/1f7f0.png?v8",
            heavy_exclamation_mark: "unicode/2757.png?v8",
            heavy_heart_exclamation: "unicode/2763.png?v8",
            heavy_minus_sign: "unicode/2796.png?v8",
            heavy_multiplication_x: "unicode/2716.png?v8",
            heavy_plus_sign: "unicode/2795.png?v8",
            hedgehog: "unicode/1f994.png?v8",
            helicopter: "unicode/1f681.png?v8",
            herb: "unicode/1f33f.png?v8",
            hibiscus: "unicode/1f33a.png?v8",
            high_brightness: "unicode/1f506.png?v8",
            high_heel: "unicode/1f460.png?v8",
            hiking_boot: "unicode/1f97e.png?v8",
            hindu_temple: "unicode/1f6d5.png?v8",
            hippopotamus: "unicode/1f99b.png?v8",
            hocho: "unicode/1f52a.png?v8",
            hole: "unicode/1f573.png?v8",
            honduras: "unicode/1f1ed-1f1f3.png?v8",
            honey_pot: "unicode/1f36f.png?v8",
            honeybee: "unicode/1f41d.png?v8",
            hong_kong: "unicode/1f1ed-1f1f0.png?v8",
            hook: "unicode/1fa9d.png?v8",
            horse: "unicode/1f434.png?v8",
            horse_racing: "unicode/1f3c7.png?v8",
            hospital: "unicode/1f3e5.png?v8",
            hot_face: "unicode/1f975.png?v8",
            hot_pepper: "unicode/1f336.png?v8",
            hotdog: "unicode/1f32d.png?v8",
            hotel: "unicode/1f3e8.png?v8",
            hotsprings: "unicode/2668.png?v8",
            hourglass: "unicode/231b.png?v8",
            hourglass_flowing_sand: "unicode/23f3.png?v8",
            house: "unicode/1f3e0.png?v8",
            house_with_garden: "unicode/1f3e1.png?v8",
            houses: "unicode/1f3d8.png?v8",
            hugs: "unicode/1f917.png?v8",
            hungary: "unicode/1f1ed-1f1fa.png?v8",
            hurtrealbad: "hurtrealbad.png?v8",
            hushed: "unicode/1f62f.png?v8",
            hut: "unicode/1f6d6.png?v8",
            hyacinth: "unicode/1fabb.png?v8",
            ice_cream: "unicode/1f368.png?v8",
            ice_cube: "unicode/1f9ca.png?v8",
            ice_hockey: "unicode/1f3d2.png?v8",
            ice_skate: "unicode/26f8.png?v8",
            icecream: "unicode/1f366.png?v8",
            iceland: "unicode/1f1ee-1f1f8.png?v8",
            id: "unicode/1f194.png?v8",
            identification_card: "unicode/1faaa.png?v8",
            ideograph_advantage: "unicode/1f250.png?v8",
            imp: "unicode/1f47f.png?v8",
            inbox_tray: "unicode/1f4e5.png?v8",
            incoming_envelope: "unicode/1f4e8.png?v8",
            index_pointing_at_the_viewer: "unicode/1faf5.png?v8",
            india: "unicode/1f1ee-1f1f3.png?v8",
            indonesia: "unicode/1f1ee-1f1e9.png?v8",
            infinity: "unicode/267e.png?v8",
            information_desk_person: "unicode/1f481.png?v8",
            information_source: "unicode/2139.png?v8",
            innocent: "unicode/1f607.png?v8",
            interrobang: "unicode/2049.png?v8",
            iphone: "unicode/1f4f1.png?v8",
            iran: "unicode/1f1ee-1f1f7.png?v8",
            iraq: "unicode/1f1ee-1f1f6.png?v8",
            ireland: "unicode/1f1ee-1f1ea.png?v8",
            isle_of_man: "unicode/1f1ee-1f1f2.png?v8",
            israel: "unicode/1f1ee-1f1f1.png?v8",
            it: "unicode/1f1ee-1f1f9.png?v8",
            izakaya_lantern: "unicode/1f3ee.png?v8",
            jack_o_lantern: "unicode/1f383.png?v8",
            jamaica: "unicode/1f1ef-1f1f2.png?v8",
            japan: "unicode/1f5fe.png?v8",
            japanese_castle: "unicode/1f3ef.png?v8",
            japanese_goblin: "unicode/1f47a.png?v8",
            japanese_ogre: "unicode/1f479.png?v8",
            jar: "unicode/1fad9.png?v8",
            jeans: "unicode/1f456.png?v8",
            jellyfish: "unicode/1fabc.png?v8",
            jersey: "unicode/1f1ef-1f1ea.png?v8",
            jigsaw: "unicode/1f9e9.png?v8",
            jordan: "unicode/1f1ef-1f1f4.png?v8",
            joy: "unicode/1f602.png?v8",
            joy_cat: "unicode/1f639.png?v8",
            joystick: "unicode/1f579.png?v8",
            jp: "unicode/1f1ef-1f1f5.png?v8",
            judge: "unicode/1f9d1-2696.png?v8",
            juggling_person: "unicode/1f939.png?v8",
            kaaba: "unicode/1f54b.png?v8",
            kangaroo: "unicode/1f998.png?v8",
            kazakhstan: "unicode/1f1f0-1f1ff.png?v8",
            kenya: "unicode/1f1f0-1f1ea.png?v8",
            key: "unicode/1f511.png?v8",
            keyboard: "unicode/2328.png?v8",
            keycap_ten: "unicode/1f51f.png?v8",
            khanda: "unicode/1faaf.png?v8",
            kick_scooter: "unicode/1f6f4.png?v8",
            kimono: "unicode/1f458.png?v8",
            kiribati: "unicode/1f1f0-1f1ee.png?v8",
            kiss: "unicode/1f48b.png?v8",
            kissing: "unicode/1f617.png?v8",
            kissing_cat: "unicode/1f63d.png?v8",
            kissing_closed_eyes: "unicode/1f61a.png?v8",
            kissing_heart: "unicode/1f618.png?v8",
            kissing_smiling_eyes: "unicode/1f619.png?v8",
            kite: "unicode/1fa81.png?v8",
            kiwi_fruit: "unicode/1f95d.png?v8",
            kneeling_man: "unicode/1f9ce-2642.png?v8",
            kneeling_person: "unicode/1f9ce.png?v8",
            kneeling_woman: "unicode/1f9ce-2640.png?v8",
            knife: "unicode/1f52a.png?v8",
            knot: "unicode/1faa2.png?v8",
            koala: "unicode/1f428.png?v8",
            koko: "unicode/1f201.png?v8",
            kosovo: "unicode/1f1fd-1f1f0.png?v8",
            kr: "unicode/1f1f0-1f1f7.png?v8",
            kuwait: "unicode/1f1f0-1f1fc.png?v8",
            kyrgyzstan: "unicode/1f1f0-1f1ec.png?v8",
            lab_coat: "unicode/1f97c.png?v8",
            label: "unicode/1f3f7.png?v8",
            lacrosse: "unicode/1f94d.png?v8",
            ladder: "unicode/1fa9c.png?v8",
            lady_beetle: "unicode/1f41e.png?v8",
            lantern: "unicode/1f3ee.png?v8",
            laos: "unicode/1f1f1-1f1e6.png?v8",
            large_blue_circle: "unicode/1f535.png?v8",
            large_blue_diamond: "unicode/1f537.png?v8",
            large_orange_diamond: "unicode/1f536.png?v8",
            last_quarter_moon: "unicode/1f317.png?v8",
            last_quarter_moon_with_face: "unicode/1f31c.png?v8",
            latin_cross: "unicode/271d.png?v8",
            latvia: "unicode/1f1f1-1f1fb.png?v8",
            laughing: "unicode/1f606.png?v8",
            leafy_green: "unicode/1f96c.png?v8",
            leaves: "unicode/1f343.png?v8",
            lebanon: "unicode/1f1f1-1f1e7.png?v8",
            ledger: "unicode/1f4d2.png?v8",
            left_luggage: "unicode/1f6c5.png?v8",
            left_right_arrow: "unicode/2194.png?v8",
            left_speech_bubble: "unicode/1f5e8.png?v8",
            leftwards_arrow_with_hook: "unicode/21a9.png?v8",
            leftwards_hand: "unicode/1faf2.png?v8",
            leftwards_pushing_hand: "unicode/1faf7.png?v8",
            leg: "unicode/1f9b5.png?v8",
            lemon: "unicode/1f34b.png?v8",
            leo: "unicode/264c.png?v8",
            leopard: "unicode/1f406.png?v8",
            lesotho: "unicode/1f1f1-1f1f8.png?v8",
            level_slider: "unicode/1f39a.png?v8",
            liberia: "unicode/1f1f1-1f1f7.png?v8",
            libra: "unicode/264e.png?v8",
            libya: "unicode/1f1f1-1f1fe.png?v8",
            liechtenstein: "unicode/1f1f1-1f1ee.png?v8",
            light_blue_heart: "unicode/1fa75.png?v8",
            light_rail: "unicode/1f688.png?v8",
            link: "unicode/1f517.png?v8",
            lion: "unicode/1f981.png?v8",
            lips: "unicode/1f444.png?v8",
            lipstick: "unicode/1f484.png?v8",
            lithuania: "unicode/1f1f1-1f1f9.png?v8",
            lizard: "unicode/1f98e.png?v8",
            llama: "unicode/1f999.png?v8",
            lobster: "unicode/1f99e.png?v8",
            lock: "unicode/1f512.png?v8",
            lock_with_ink_pen: "unicode/1f50f.png?v8",
            lollipop: "unicode/1f36d.png?v8",
            long_drum: "unicode/1fa98.png?v8",
            loop: "unicode/27bf.png?v8",
            lotion_bottle: "unicode/1f9f4.png?v8",
            lotus: "unicode/1fab7.png?v8",
            lotus_position: "unicode/1f9d8.png?v8",
            lotus_position_man: "unicode/1f9d8-2642.png?v8",
            lotus_position_woman: "unicode/1f9d8-2640.png?v8",
            loud_sound: "unicode/1f50a.png?v8",
            loudspeaker: "unicode/1f4e2.png?v8",
            love_hotel: "unicode/1f3e9.png?v8",
            love_letter: "unicode/1f48c.png?v8",
            love_you_gesture: "unicode/1f91f.png?v8",
            low_battery: "unicode/1faab.png?v8",
            low_brightness: "unicode/1f505.png?v8",
            luggage: "unicode/1f9f3.png?v8",
            lungs: "unicode/1fac1.png?v8",
            luxembourg: "unicode/1f1f1-1f1fa.png?v8",
            lying_face: "unicode/1f925.png?v8",
            m: "unicode/24c2.png?v8",
            macau: "unicode/1f1f2-1f1f4.png?v8",
            macedonia: "unicode/1f1f2-1f1f0.png?v8",
            madagascar: "unicode/1f1f2-1f1ec.png?v8",
            mag: "unicode/1f50d.png?v8",
            mag_right: "unicode/1f50e.png?v8",
            mage: "unicode/1f9d9.png?v8",
            mage_man: "unicode/1f9d9-2642.png?v8",
            mage_woman: "unicode/1f9d9-2640.png?v8",
            magic_wand: "unicode/1fa84.png?v8",
            magnet: "unicode/1f9f2.png?v8",
            mahjong: "unicode/1f004.png?v8",
            mailbox: "unicode/1f4eb.png?v8",
            mailbox_closed: "unicode/1f4ea.png?v8",
            mailbox_with_mail: "unicode/1f4ec.png?v8",
            mailbox_with_no_mail: "unicode/1f4ed.png?v8",
            malawi: "unicode/1f1f2-1f1fc.png?v8",
            malaysia: "unicode/1f1f2-1f1fe.png?v8",
            maldives: "unicode/1f1f2-1f1fb.png?v8",
            male_detective: "unicode/1f575-2642.png?v8",
            male_sign: "unicode/2642.png?v8",
            mali: "unicode/1f1f2-1f1f1.png?v8",
            malta: "unicode/1f1f2-1f1f9.png?v8",
            mammoth: "unicode/1f9a3.png?v8",
            man: "unicode/1f468.png?v8",
            man_artist: "unicode/1f468-1f3a8.png?v8",
            man_astronaut: "unicode/1f468-1f680.png?v8",
            man_beard: "unicode/1f9d4-2642.png?v8",
            man_cartwheeling: "unicode/1f938-2642.png?v8",
            man_cook: "unicode/1f468-1f373.png?v8",
            man_dancing: "unicode/1f57a.png?v8",
            man_facepalming: "unicode/1f926-2642.png?v8",
            man_factory_worker: "unicode/1f468-1f3ed.png?v8",
            man_farmer: "unicode/1f468-1f33e.png?v8",
            man_feeding_baby: "unicode/1f468-1f37c.png?v8",
            man_firefighter: "unicode/1f468-1f692.png?v8",
            man_health_worker: "unicode/1f468-2695.png?v8",
            man_in_manual_wheelchair: "unicode/1f468-1f9bd.png?v8",
            man_in_motorized_wheelchair: "unicode/1f468-1f9bc.png?v8",
            man_in_tuxedo: "unicode/1f935-2642.png?v8",
            man_judge: "unicode/1f468-2696.png?v8",
            man_juggling: "unicode/1f939-2642.png?v8",
            man_mechanic: "unicode/1f468-1f527.png?v8",
            man_office_worker: "unicode/1f468-1f4bc.png?v8",
            man_pilot: "unicode/1f468-2708.png?v8",
            man_playing_handball: "unicode/1f93e-2642.png?v8",
            man_playing_water_polo: "unicode/1f93d-2642.png?v8",
            man_scientist: "unicode/1f468-1f52c.png?v8",
            man_shrugging: "unicode/1f937-2642.png?v8",
            man_singer: "unicode/1f468-1f3a4.png?v8",
            man_student: "unicode/1f468-1f393.png?v8",
            man_teacher: "unicode/1f468-1f3eb.png?v8",
            man_technologist: "unicode/1f468-1f4bb.png?v8",
            man_with_gua_pi_mao: "unicode/1f472.png?v8",
            man_with_probing_cane: "unicode/1f468-1f9af.png?v8",
            man_with_turban: "unicode/1f473-2642.png?v8",
            man_with_veil: "unicode/1f470-2642.png?v8",
            mandarin: "unicode/1f34a.png?v8",
            mango: "unicode/1f96d.png?v8",
            mans_shoe: "unicode/1f45e.png?v8",
            mantelpiece_clock: "unicode/1f570.png?v8",
            manual_wheelchair: "unicode/1f9bd.png?v8",
            maple_leaf: "unicode/1f341.png?v8",
            maracas: "unicode/1fa87.png?v8",
            marshall_islands: "unicode/1f1f2-1f1ed.png?v8",
            martial_arts_uniform: "unicode/1f94b.png?v8",
            martinique: "unicode/1f1f2-1f1f6.png?v8",
            mask: "unicode/1f637.png?v8",
            massage: "unicode/1f486.png?v8",
            massage_man: "unicode/1f486-2642.png?v8",
            massage_woman: "unicode/1f486-2640.png?v8",
            mate: "unicode/1f9c9.png?v8",
            mauritania: "unicode/1f1f2-1f1f7.png?v8",
            mauritius: "unicode/1f1f2-1f1fa.png?v8",
            mayotte: "unicode/1f1fe-1f1f9.png?v8",
            meat_on_bone: "unicode/1f356.png?v8",
            mechanic: "unicode/1f9d1-1f527.png?v8",
            mechanical_arm: "unicode/1f9be.png?v8",
            mechanical_leg: "unicode/1f9bf.png?v8",
            medal_military: "unicode/1f396.png?v8",
            medal_sports: "unicode/1f3c5.png?v8",
            medical_symbol: "unicode/2695.png?v8",
            mega: "unicode/1f4e3.png?v8",
            melon: "unicode/1f348.png?v8",
            melting_face: "unicode/1fae0.png?v8",
            memo: "unicode/1f4dd.png?v8",
            men_wrestling: "unicode/1f93c-2642.png?v8",
            mending_heart: "unicode/2764-1fa79.png?v8",
            menorah: "unicode/1f54e.png?v8",
            mens: "unicode/1f6b9.png?v8",
            mermaid: "unicode/1f9dc-2640.png?v8",
            merman: "unicode/1f9dc-2642.png?v8",
            merperson: "unicode/1f9dc.png?v8",
            metal: "unicode/1f918.png?v8",
            metro: "unicode/1f687.png?v8",
            mexico: "unicode/1f1f2-1f1fd.png?v8",
            microbe: "unicode/1f9a0.png?v8",
            micronesia: "unicode/1f1eb-1f1f2.png?v8",
            microphone: "unicode/1f3a4.png?v8",
            microscope: "unicode/1f52c.png?v8",
            middle_finger: "unicode/1f595.png?v8",
            military_helmet: "unicode/1fa96.png?v8",
            milk_glass: "unicode/1f95b.png?v8",
            milky_way: "unicode/1f30c.png?v8",
            minibus: "unicode/1f690.png?v8",
            minidisc: "unicode/1f4bd.png?v8",
            mirror: "unicode/1fa9e.png?v8",
            mirror_ball: "unicode/1faa9.png?v8",
            mobile_phone_off: "unicode/1f4f4.png?v8",
            moldova: "unicode/1f1f2-1f1e9.png?v8",
            monaco: "unicode/1f1f2-1f1e8.png?v8",
            money_mouth_face: "unicode/1f911.png?v8",
            money_with_wings: "unicode/1f4b8.png?v8",
            moneybag: "unicode/1f4b0.png?v8",
            mongolia: "unicode/1f1f2-1f1f3.png?v8",
            monkey: "unicode/1f412.png?v8",
            monkey_face: "unicode/1f435.png?v8",
            monocle_face: "unicode/1f9d0.png?v8",
            monorail: "unicode/1f69d.png?v8",
            montenegro: "unicode/1f1f2-1f1ea.png?v8",
            montserrat: "unicode/1f1f2-1f1f8.png?v8",
            moon: "unicode/1f314.png?v8",
            moon_cake: "unicode/1f96e.png?v8",
            moose: "unicode/1face.png?v8",
            morocco: "unicode/1f1f2-1f1e6.png?v8",
            mortar_board: "unicode/1f393.png?v8",
            mosque: "unicode/1f54c.png?v8",
            mosquito: "unicode/1f99f.png?v8",
            motor_boat: "unicode/1f6e5.png?v8",
            motor_scooter: "unicode/1f6f5.png?v8",
            motorcycle: "unicode/1f3cd.png?v8",
            motorized_wheelchair: "unicode/1f9bc.png?v8",
            motorway: "unicode/1f6e3.png?v8",
            mount_fuji: "unicode/1f5fb.png?v8",
            mountain: "unicode/26f0.png?v8",
            mountain_bicyclist: "unicode/1f6b5.png?v8",
            mountain_biking_man: "unicode/1f6b5-2642.png?v8",
            mountain_biking_woman: "unicode/1f6b5-2640.png?v8",
            mountain_cableway: "unicode/1f6a0.png?v8",
            mountain_railway: "unicode/1f69e.png?v8",
            mountain_snow: "unicode/1f3d4.png?v8",
            mouse: "unicode/1f42d.png?v8",
            mouse2: "unicode/1f401.png?v8",
            mouse_trap: "unicode/1faa4.png?v8",
            movie_camera: "unicode/1f3a5.png?v8",
            moyai: "unicode/1f5ff.png?v8",
            mozambique: "unicode/1f1f2-1f1ff.png?v8",
            mrs_claus: "unicode/1f936.png?v8",
            muscle: "unicode/1f4aa.png?v8",
            mushroom: "unicode/1f344.png?v8",
            musical_keyboard: "unicode/1f3b9.png?v8",
            musical_note: "unicode/1f3b5.png?v8",
            musical_score: "unicode/1f3bc.png?v8",
            mute: "unicode/1f507.png?v8",
            mx_claus: "unicode/1f9d1-1f384.png?v8",
            myanmar: "unicode/1f1f2-1f1f2.png?v8",
            nail_care: "unicode/1f485.png?v8",
            name_badge: "unicode/1f4db.png?v8",
            namibia: "unicode/1f1f3-1f1e6.png?v8",
            national_park: "unicode/1f3de.png?v8",
            nauru: "unicode/1f1f3-1f1f7.png?v8",
            nauseated_face: "unicode/1f922.png?v8",
            nazar_amulet: "unicode/1f9ff.png?v8",
            neckbeard: "neckbeard.png?v8",
            necktie: "unicode/1f454.png?v8",
            negative_squared_cross_mark: "unicode/274e.png?v8",
            nepal: "unicode/1f1f3-1f1f5.png?v8",
            nerd_face: "unicode/1f913.png?v8",
            nest_with_eggs: "unicode/1faba.png?v8",
            nesting_dolls: "unicode/1fa86.png?v8",
            netherlands: "unicode/1f1f3-1f1f1.png?v8",
            neutral_face: "unicode/1f610.png?v8",
            new: "unicode/1f195.png?v8",
            new_caledonia: "unicode/1f1f3-1f1e8.png?v8",
            new_moon: "unicode/1f311.png?v8",
            new_moon_with_face: "unicode/1f31a.png?v8",
            new_zealand: "unicode/1f1f3-1f1ff.png?v8",
            newspaper: "unicode/1f4f0.png?v8",
            newspaper_roll: "unicode/1f5de.png?v8",
            next_track_button: "unicode/23ed.png?v8",
            ng: "unicode/1f196.png?v8",
            ng_man: "unicode/1f645-2642.png?v8",
            ng_woman: "unicode/1f645-2640.png?v8",
            nicaragua: "unicode/1f1f3-1f1ee.png?v8",
            niger: "unicode/1f1f3-1f1ea.png?v8",
            nigeria: "unicode/1f1f3-1f1ec.png?v8",
            night_with_stars: "unicode/1f303.png?v8",
            nine: "unicode/0039-20e3.png?v8",
            ninja: "unicode/1f977.png?v8",
            niue: "unicode/1f1f3-1f1fa.png?v8",
            no_bell: "unicode/1f515.png?v8",
            no_bicycles: "unicode/1f6b3.png?v8",
            no_entry: "unicode/26d4.png?v8",
            no_entry_sign: "unicode/1f6ab.png?v8",
            no_good: "unicode/1f645.png?v8",
            no_good_man: "unicode/1f645-2642.png?v8",
            no_good_woman: "unicode/1f645-2640.png?v8",
            no_mobile_phones: "unicode/1f4f5.png?v8",
            no_mouth: "unicode/1f636.png?v8",
            no_pedestrians: "unicode/1f6b7.png?v8",
            no_smoking: "unicode/1f6ad.png?v8",
            "non-potable_water": "unicode/1f6b1.png?v8",
            norfolk_island: "unicode/1f1f3-1f1eb.png?v8",
            north_korea: "unicode/1f1f0-1f1f5.png?v8",
            northern_mariana_islands: "unicode/1f1f2-1f1f5.png?v8",
            norway: "unicode/1f1f3-1f1f4.png?v8",
            nose: "unicode/1f443.png?v8",
            notebook: "unicode/1f4d3.png?v8",
            notebook_with_decorative_cover: "unicode/1f4d4.png?v8",
            notes: "unicode/1f3b6.png?v8",
            nut_and_bolt: "unicode/1f529.png?v8",
            o: "unicode/2b55.png?v8",
            o2: "unicode/1f17e.png?v8",
            ocean: "unicode/1f30a.png?v8",
            octocat: "octocat.png?v8",
            octopus: "unicode/1f419.png?v8",
            oden: "unicode/1f362.png?v8",
            office: "unicode/1f3e2.png?v8",
            office_worker: "unicode/1f9d1-1f4bc.png?v8",
            oil_drum: "unicode/1f6e2.png?v8",
            ok: "unicode/1f197.png?v8",
            ok_hand: "unicode/1f44c.png?v8",
            ok_man: "unicode/1f646-2642.png?v8",
            ok_person: "unicode/1f646.png?v8",
            ok_woman: "unicode/1f646-2640.png?v8",
            old_key: "unicode/1f5dd.png?v8",
            older_adult: "unicode/1f9d3.png?v8",
            older_man: "unicode/1f474.png?v8",
            older_woman: "unicode/1f475.png?v8",
            olive: "unicode/1fad2.png?v8",
            om: "unicode/1f549.png?v8",
            oman: "unicode/1f1f4-1f1f2.png?v8",
            on: "unicode/1f51b.png?v8",
            oncoming_automobile: "unicode/1f698.png?v8",
            oncoming_bus: "unicode/1f68d.png?v8",
            oncoming_police_car: "unicode/1f694.png?v8",
            oncoming_taxi: "unicode/1f696.png?v8",
            one: "unicode/0031-20e3.png?v8",
            one_piece_swimsuit: "unicode/1fa71.png?v8",
            onion: "unicode/1f9c5.png?v8",
            open_book: "unicode/1f4d6.png?v8",
            open_file_folder: "unicode/1f4c2.png?v8",
            open_hands: "unicode/1f450.png?v8",
            open_mouth: "unicode/1f62e.png?v8",
            open_umbrella: "unicode/2602.png?v8",
            ophiuchus: "unicode/26ce.png?v8",
            orange: "unicode/1f34a.png?v8",
            orange_book: "unicode/1f4d9.png?v8",
            orange_circle: "unicode/1f7e0.png?v8",
            orange_heart: "unicode/1f9e1.png?v8",
            orange_square: "unicode/1f7e7.png?v8",
            orangutan: "unicode/1f9a7.png?v8",
            orthodox_cross: "unicode/2626.png?v8",
            otter: "unicode/1f9a6.png?v8",
            outbox_tray: "unicode/1f4e4.png?v8",
            owl: "unicode/1f989.png?v8",
            ox: "unicode/1f402.png?v8",
            oyster: "unicode/1f9aa.png?v8",
            package: "unicode/1f4e6.png?v8",
            page_facing_up: "unicode/1f4c4.png?v8",
            page_with_curl: "unicode/1f4c3.png?v8",
            pager: "unicode/1f4df.png?v8",
            paintbrush: "unicode/1f58c.png?v8",
            pakistan: "unicode/1f1f5-1f1f0.png?v8",
            palau: "unicode/1f1f5-1f1fc.png?v8",
            palestinian_territories: "unicode/1f1f5-1f1f8.png?v8",
            palm_down_hand: "unicode/1faf3.png?v8",
            palm_tree: "unicode/1f334.png?v8",
            palm_up_hand: "unicode/1faf4.png?v8",
            palms_up_together: "unicode/1f932.png?v8",
            panama: "unicode/1f1f5-1f1e6.png?v8",
            pancakes: "unicode/1f95e.png?v8",
            panda_face: "unicode/1f43c.png?v8",
            paperclip: "unicode/1f4ce.png?v8",
            paperclips: "unicode/1f587.png?v8",
            papua_new_guinea: "unicode/1f1f5-1f1ec.png?v8",
            parachute: "unicode/1fa82.png?v8",
            paraguay: "unicode/1f1f5-1f1fe.png?v8",
            parasol_on_ground: "unicode/26f1.png?v8",
            parking: "unicode/1f17f.png?v8",
            parrot: "unicode/1f99c.png?v8",
            part_alternation_mark: "unicode/303d.png?v8",
            partly_sunny: "unicode/26c5.png?v8",
            partying_face: "unicode/1f973.png?v8",
            passenger_ship: "unicode/1f6f3.png?v8",
            passport_control: "unicode/1f6c2.png?v8",
            pause_button: "unicode/23f8.png?v8",
            paw_prints: "unicode/1f43e.png?v8",
            pea_pod: "unicode/1fadb.png?v8",
            peace_symbol: "unicode/262e.png?v8",
            peach: "unicode/1f351.png?v8",
            peacock: "unicode/1f99a.png?v8",
            peanuts: "unicode/1f95c.png?v8",
            pear: "unicode/1f350.png?v8",
            pen: "unicode/1f58a.png?v8",
            pencil: "unicode/1f4dd.png?v8",
            pencil2: "unicode/270f.png?v8",
            penguin: "unicode/1f427.png?v8",
            pensive: "unicode/1f614.png?v8",
            people_holding_hands: "unicode/1f9d1-1f91d-1f9d1.png?v8",
            people_hugging: "unicode/1fac2.png?v8",
            performing_arts: "unicode/1f3ad.png?v8",
            persevere: "unicode/1f623.png?v8",
            person_bald: "unicode/1f9d1-1f9b2.png?v8",
            person_curly_hair: "unicode/1f9d1-1f9b1.png?v8",
            person_feeding_baby: "unicode/1f9d1-1f37c.png?v8",
            person_fencing: "unicode/1f93a.png?v8",
            person_in_manual_wheelchair: "unicode/1f9d1-1f9bd.png?v8",
            person_in_motorized_wheelchair: "unicode/1f9d1-1f9bc.png?v8",
            person_in_tuxedo: "unicode/1f935.png?v8",
            person_red_hair: "unicode/1f9d1-1f9b0.png?v8",
            person_white_hair: "unicode/1f9d1-1f9b3.png?v8",
            person_with_crown: "unicode/1fac5.png?v8",
            person_with_probing_cane: "unicode/1f9d1-1f9af.png?v8",
            person_with_turban: "unicode/1f473.png?v8",
            person_with_veil: "unicode/1f470.png?v8",
            peru: "unicode/1f1f5-1f1ea.png?v8",
            petri_dish: "unicode/1f9eb.png?v8",
            philippines: "unicode/1f1f5-1f1ed.png?v8",
            phone: "unicode/260e.png?v8",
            pick: "unicode/26cf.png?v8",
            pickup_truck: "unicode/1f6fb.png?v8",
            pie: "unicode/1f967.png?v8",
            pig: "unicode/1f437.png?v8",
            pig2: "unicode/1f416.png?v8",
            pig_nose: "unicode/1f43d.png?v8",
            pill: "unicode/1f48a.png?v8",
            pilot: "unicode/1f9d1-2708.png?v8",
            pinata: "unicode/1fa85.png?v8",
            pinched_fingers: "unicode/1f90c.png?v8",
            pinching_hand: "unicode/1f90f.png?v8",
            pineapple: "unicode/1f34d.png?v8",
            ping_pong: "unicode/1f3d3.png?v8",
            pink_heart: "unicode/1fa77.png?v8",
            pirate_flag: "unicode/1f3f4-2620.png?v8",
            pisces: "unicode/2653.png?v8",
            pitcairn_islands: "unicode/1f1f5-1f1f3.png?v8",
            pizza: "unicode/1f355.png?v8",
            placard: "unicode/1faa7.png?v8",
            place_of_worship: "unicode/1f6d0.png?v8",
            plate_with_cutlery: "unicode/1f37d.png?v8",
            play_or_pause_button: "unicode/23ef.png?v8",
            playground_slide: "unicode/1f6dd.png?v8",
            pleading_face: "unicode/1f97a.png?v8",
            plunger: "unicode/1faa0.png?v8",
            point_down: "unicode/1f447.png?v8",
            point_left: "unicode/1f448.png?v8",
            point_right: "unicode/1f449.png?v8",
            point_up: "unicode/261d.png?v8",
            point_up_2: "unicode/1f446.png?v8",
            poland: "unicode/1f1f5-1f1f1.png?v8",
            polar_bear: "unicode/1f43b-2744.png?v8",
            police_car: "unicode/1f693.png?v8",
            police_officer: "unicode/1f46e.png?v8",
            policeman: "unicode/1f46e-2642.png?v8",
            policewoman: "unicode/1f46e-2640.png?v8",
            poodle: "unicode/1f429.png?v8",
            poop: "unicode/1f4a9.png?v8",
            popcorn: "unicode/1f37f.png?v8",
            portugal: "unicode/1f1f5-1f1f9.png?v8",
            post_office: "unicode/1f3e3.png?v8",
            postal_horn: "unicode/1f4ef.png?v8",
            postbox: "unicode/1f4ee.png?v8",
            potable_water: "unicode/1f6b0.png?v8",
            potato: "unicode/1f954.png?v8",
            potted_plant: "unicode/1fab4.png?v8",
            pouch: "unicode/1f45d.png?v8",
            poultry_leg: "unicode/1f357.png?v8",
            pound: "unicode/1f4b7.png?v8",
            pouring_liquid: "unicode/1fad7.png?v8",
            pout: "unicode/1f621.png?v8",
            pouting_cat: "unicode/1f63e.png?v8",
            pouting_face: "unicode/1f64e.png?v8",
            pouting_man: "unicode/1f64e-2642.png?v8",
            pouting_woman: "unicode/1f64e-2640.png?v8",
            pray: "unicode/1f64f.png?v8",
            prayer_beads: "unicode/1f4ff.png?v8",
            pregnant_man: "unicode/1fac3.png?v8",
            pregnant_person: "unicode/1fac4.png?v8",
            pregnant_woman: "unicode/1f930.png?v8",
            pretzel: "unicode/1f968.png?v8",
            previous_track_button: "unicode/23ee.png?v8",
            prince: "unicode/1f934.png?v8",
            princess: "unicode/1f478.png?v8",
            printer: "unicode/1f5a8.png?v8",
            probing_cane: "unicode/1f9af.png?v8",
            puerto_rico: "unicode/1f1f5-1f1f7.png?v8",
            punch: "unicode/1f44a.png?v8",
            purple_circle: "unicode/1f7e3.png?v8",
            purple_heart: "unicode/1f49c.png?v8",
            purple_square: "unicode/1f7ea.png?v8",
            purse: "unicode/1f45b.png?v8",
            pushpin: "unicode/1f4cc.png?v8",
            put_litter_in_its_place: "unicode/1f6ae.png?v8",
            qatar: "unicode/1f1f6-1f1e6.png?v8",
            question: "unicode/2753.png?v8",
            rabbit: "unicode/1f430.png?v8",
            rabbit2: "unicode/1f407.png?v8",
            raccoon: "unicode/1f99d.png?v8",
            racehorse: "unicode/1f40e.png?v8",
            racing_car: "unicode/1f3ce.png?v8",
            radio: "unicode/1f4fb.png?v8",
            radio_button: "unicode/1f518.png?v8",
            radioactive: "unicode/2622.png?v8",
            rage: "unicode/1f621.png?v8",
            rage1: "rage1.png?v8",
            rage2: "rage2.png?v8",
            rage3: "rage3.png?v8",
            rage4: "rage4.png?v8",
            railway_car: "unicode/1f683.png?v8",
            railway_track: "unicode/1f6e4.png?v8",
            rainbow: "unicode/1f308.png?v8",
            rainbow_flag: "unicode/1f3f3-1f308.png?v8",
            raised_back_of_hand: "unicode/1f91a.png?v8",
            raised_eyebrow: "unicode/1f928.png?v8",
            raised_hand: "unicode/270b.png?v8",
            raised_hand_with_fingers_splayed: "unicode/1f590.png?v8",
            raised_hands: "unicode/1f64c.png?v8",
            raising_hand: "unicode/1f64b.png?v8",
            raising_hand_man: "unicode/1f64b-2642.png?v8",
            raising_hand_woman: "unicode/1f64b-2640.png?v8",
            ram: "unicode/1f40f.png?v8",
            ramen: "unicode/1f35c.png?v8",
            rat: "unicode/1f400.png?v8",
            razor: "unicode/1fa92.png?v8",
            receipt: "unicode/1f9fe.png?v8",
            record_button: "unicode/23fa.png?v8",
            recycle: "unicode/267b.png?v8",
            red_car: "unicode/1f697.png?v8",
            red_circle: "unicode/1f534.png?v8",
            red_envelope: "unicode/1f9e7.png?v8",
            red_haired_man: "unicode/1f468-1f9b0.png?v8",
            red_haired_woman: "unicode/1f469-1f9b0.png?v8",
            red_square: "unicode/1f7e5.png?v8",
            registered: "unicode/00ae.png?v8",
            relaxed: "unicode/263a.png?v8",
            relieved: "unicode/1f60c.png?v8",
            reminder_ribbon: "unicode/1f397.png?v8",
            repeat: "unicode/1f501.png?v8",
            repeat_one: "unicode/1f502.png?v8",
            rescue_worker_helmet: "unicode/26d1.png?v8",
            restroom: "unicode/1f6bb.png?v8",
            reunion: "unicode/1f1f7-1f1ea.png?v8",
            revolving_hearts: "unicode/1f49e.png?v8",
            rewind: "unicode/23ea.png?v8",
            rhinoceros: "unicode/1f98f.png?v8",
            ribbon: "unicode/1f380.png?v8",
            rice: "unicode/1f35a.png?v8",
            rice_ball: "unicode/1f359.png?v8",
            rice_cracker: "unicode/1f358.png?v8",
            rice_scene: "unicode/1f391.png?v8",
            right_anger_bubble: "unicode/1f5ef.png?v8",
            rightwards_hand: "unicode/1faf1.png?v8",
            rightwards_pushing_hand: "unicode/1faf8.png?v8",
            ring: "unicode/1f48d.png?v8",
            ring_buoy: "unicode/1f6df.png?v8",
            ringed_planet: "unicode/1fa90.png?v8",
            robot: "unicode/1f916.png?v8",
            rock: "unicode/1faa8.png?v8",
            rocket: "unicode/1f680.png?v8",
            rofl: "unicode/1f923.png?v8",
            roll_eyes: "unicode/1f644.png?v8",
            roll_of_paper: "unicode/1f9fb.png?v8",
            roller_coaster: "unicode/1f3a2.png?v8",
            roller_skate: "unicode/1f6fc.png?v8",
            romania: "unicode/1f1f7-1f1f4.png?v8",
            rooster: "unicode/1f413.png?v8",
            rose: "unicode/1f339.png?v8",
            rosette: "unicode/1f3f5.png?v8",
            rotating_light: "unicode/1f6a8.png?v8",
            round_pushpin: "unicode/1f4cd.png?v8",
            rowboat: "unicode/1f6a3.png?v8",
            rowing_man: "unicode/1f6a3-2642.png?v8",
            rowing_woman: "unicode/1f6a3-2640.png?v8",
            ru: "unicode/1f1f7-1f1fa.png?v8",
            rugby_football: "unicode/1f3c9.png?v8",
            runner: "unicode/1f3c3.png?v8",
            running: "unicode/1f3c3.png?v8",
            running_man: "unicode/1f3c3-2642.png?v8",
            running_shirt_with_sash: "unicode/1f3bd.png?v8",
            running_woman: "unicode/1f3c3-2640.png?v8",
            rwanda: "unicode/1f1f7-1f1fc.png?v8",
            sa: "unicode/1f202.png?v8",
            safety_pin: "unicode/1f9f7.png?v8",
            safety_vest: "unicode/1f9ba.png?v8",
            sagittarius: "unicode/2650.png?v8",
            sailboat: "unicode/26f5.png?v8",
            sake: "unicode/1f376.png?v8",
            salt: "unicode/1f9c2.png?v8",
            saluting_face: "unicode/1fae1.png?v8",
            samoa: "unicode/1f1fc-1f1f8.png?v8",
            san_marino: "unicode/1f1f8-1f1f2.png?v8",
            sandal: "unicode/1f461.png?v8",
            sandwich: "unicode/1f96a.png?v8",
            santa: "unicode/1f385.png?v8",
            sao_tome_principe: "unicode/1f1f8-1f1f9.png?v8",
            sari: "unicode/1f97b.png?v8",
            sassy_man: "unicode/1f481-2642.png?v8",
            sassy_woman: "unicode/1f481-2640.png?v8",
            satellite: "unicode/1f4e1.png?v8",
            satisfied: "unicode/1f606.png?v8",
            saudi_arabia: "unicode/1f1f8-1f1e6.png?v8",
            sauna_man: "unicode/1f9d6-2642.png?v8",
            sauna_person: "unicode/1f9d6.png?v8",
            sauna_woman: "unicode/1f9d6-2640.png?v8",
            sauropod: "unicode/1f995.png?v8",
            saxophone: "unicode/1f3b7.png?v8",
            scarf: "unicode/1f9e3.png?v8",
            school: "unicode/1f3eb.png?v8",
            school_satchel: "unicode/1f392.png?v8",
            scientist: "unicode/1f9d1-1f52c.png?v8",
            scissors: "unicode/2702.png?v8",
            scorpion: "unicode/1f982.png?v8",
            scorpius: "unicode/264f.png?v8",
            scotland: "unicode/1f3f4-e0067-e0062-e0073-e0063-e0074-e007f.png?v8",
            scream: "unicode/1f631.png?v8",
            scream_cat: "unicode/1f640.png?v8",
            screwdriver: "unicode/1fa9b.png?v8",
            scroll: "unicode/1f4dc.png?v8",
            seal: "unicode/1f9ad.png?v8",
            seat: "unicode/1f4ba.png?v8",
            secret: "unicode/3299.png?v8",
            see_no_evil: "unicode/1f648.png?v8",
            seedling: "unicode/1f331.png?v8",
            selfie: "unicode/1f933.png?v8",
            senegal: "unicode/1f1f8-1f1f3.png?v8",
            serbia: "unicode/1f1f7-1f1f8.png?v8",
            service_dog: "unicode/1f415-1f9ba.png?v8",
            seven: "unicode/0037-20e3.png?v8",
            sewing_needle: "unicode/1faa1.png?v8",
            seychelles: "unicode/1f1f8-1f1e8.png?v8",
            shaking_face: "unicode/1fae8.png?v8",
            shallow_pan_of_food: "unicode/1f958.png?v8",
            shamrock: "unicode/2618.png?v8",
            shark: "unicode/1f988.png?v8",
            shaved_ice: "unicode/1f367.png?v8",
            sheep: "unicode/1f411.png?v8",
            shell: "unicode/1f41a.png?v8",
            shield: "unicode/1f6e1.png?v8",
            shinto_shrine: "unicode/26e9.png?v8",
            ship: "unicode/1f6a2.png?v8",
            shipit: "shipit.png?v8",
            shirt: "unicode/1f455.png?v8",
            shit: "unicode/1f4a9.png?v8",
            shoe: "unicode/1f45e.png?v8",
            shopping: "unicode/1f6cd.png?v8",
            shopping_cart: "unicode/1f6d2.png?v8",
            shorts: "unicode/1fa73.png?v8",
            shower: "unicode/1f6bf.png?v8",
            shrimp: "unicode/1f990.png?v8",
            shrug: "unicode/1f937.png?v8",
            shushing_face: "unicode/1f92b.png?v8",
            sierra_leone: "unicode/1f1f8-1f1f1.png?v8",
            signal_strength: "unicode/1f4f6.png?v8",
            singapore: "unicode/1f1f8-1f1ec.png?v8",
            singer: "unicode/1f9d1-1f3a4.png?v8",
            sint_maarten: "unicode/1f1f8-1f1fd.png?v8",
            six: "unicode/0036-20e3.png?v8",
            six_pointed_star: "unicode/1f52f.png?v8",
            skateboard: "unicode/1f6f9.png?v8",
            ski: "unicode/1f3bf.png?v8",
            skier: "unicode/26f7.png?v8",
            skull: "unicode/1f480.png?v8",
            skull_and_crossbones: "unicode/2620.png?v8",
            skunk: "unicode/1f9a8.png?v8",
            sled: "unicode/1f6f7.png?v8",
            sleeping: "unicode/1f634.png?v8",
            sleeping_bed: "unicode/1f6cc.png?v8",
            sleepy: "unicode/1f62a.png?v8",
            slightly_frowning_face: "unicode/1f641.png?v8",
            slightly_smiling_face: "unicode/1f642.png?v8",
            slot_machine: "unicode/1f3b0.png?v8",
            sloth: "unicode/1f9a5.png?v8",
            slovakia: "unicode/1f1f8-1f1f0.png?v8",
            slovenia: "unicode/1f1f8-1f1ee.png?v8",
            small_airplane: "unicode/1f6e9.png?v8",
            small_blue_diamond: "unicode/1f539.png?v8",
            small_orange_diamond: "unicode/1f538.png?v8",
            small_red_triangle: "unicode/1f53a.png?v8",
            small_red_triangle_down: "unicode/1f53b.png?v8",
            smile: "unicode/1f604.png?v8",
            smile_cat: "unicode/1f638.png?v8",
            smiley: "unicode/1f603.png?v8",
            smiley_cat: "unicode/1f63a.png?v8",
            smiling_face_with_tear: "unicode/1f972.png?v8",
            smiling_face_with_three_hearts: "unicode/1f970.png?v8",
            smiling_imp: "unicode/1f608.png?v8",
            smirk: "unicode/1f60f.png?v8",
            smirk_cat: "unicode/1f63c.png?v8",
            smoking: "unicode/1f6ac.png?v8",
            snail: "unicode/1f40c.png?v8",
            snake: "unicode/1f40d.png?v8",
            sneezing_face: "unicode/1f927.png?v8",
            snowboarder: "unicode/1f3c2.png?v8",
            snowflake: "unicode/2744.png?v8",
            snowman: "unicode/26c4.png?v8",
            snowman_with_snow: "unicode/2603.png?v8",
            soap: "unicode/1f9fc.png?v8",
            sob: "unicode/1f62d.png?v8",
            soccer: "unicode/26bd.png?v8",
            socks: "unicode/1f9e6.png?v8",
            softball: "unicode/1f94e.png?v8",
            solomon_islands: "unicode/1f1f8-1f1e7.png?v8",
            somalia: "unicode/1f1f8-1f1f4.png?v8",
            soon: "unicode/1f51c.png?v8",
            sos: "unicode/1f198.png?v8",
            sound: "unicode/1f509.png?v8",
            south_africa: "unicode/1f1ff-1f1e6.png?v8",
            south_georgia_south_sandwich_islands: "unicode/1f1ec-1f1f8.png?v8",
            south_sudan: "unicode/1f1f8-1f1f8.png?v8",
            space_invader: "unicode/1f47e.png?v8",
            spades: "unicode/2660.png?v8",
            spaghetti: "unicode/1f35d.png?v8",
            sparkle: "unicode/2747.png?v8",
            sparkler: "unicode/1f387.png?v8",
            sparkles: "unicode/2728.png?v8",
            sparkling_heart: "unicode/1f496.png?v8",
            speak_no_evil: "unicode/1f64a.png?v8",
            speaker: "unicode/1f508.png?v8",
            speaking_head: "unicode/1f5e3.png?v8",
            speech_balloon: "unicode/1f4ac.png?v8",
            speedboat: "unicode/1f6a4.png?v8",
            spider: "unicode/1f577.png?v8",
            spider_web: "unicode/1f578.png?v8",
            spiral_calendar: "unicode/1f5d3.png?v8",
            spiral_notepad: "unicode/1f5d2.png?v8",
            sponge: "unicode/1f9fd.png?v8",
            spoon: "unicode/1f944.png?v8",
            squid: "unicode/1f991.png?v8",
            sri_lanka: "unicode/1f1f1-1f1f0.png?v8",
            st_barthelemy: "unicode/1f1e7-1f1f1.png?v8",
            st_helena: "unicode/1f1f8-1f1ed.png?v8",
            st_kitts_nevis: "unicode/1f1f0-1f1f3.png?v8",
            st_lucia: "unicode/1f1f1-1f1e8.png?v8",
            st_martin: "unicode/1f1f2-1f1eb.png?v8",
            st_pierre_miquelon: "unicode/1f1f5-1f1f2.png?v8",
            st_vincent_grenadines: "unicode/1f1fb-1f1e8.png?v8",
            stadium: "unicode/1f3df.png?v8",
            standing_man: "unicode/1f9cd-2642.png?v8",
            standing_person: "unicode/1f9cd.png?v8",
            standing_woman: "unicode/1f9cd-2640.png?v8",
            star: "unicode/2b50.png?v8",
            star2: "unicode/1f31f.png?v8",
            star_and_crescent: "unicode/262a.png?v8",
            star_of_david: "unicode/2721.png?v8",
            star_struck: "unicode/1f929.png?v8",
            stars: "unicode/1f320.png?v8",
            station: "unicode/1f689.png?v8",
            statue_of_liberty: "unicode/1f5fd.png?v8",
            steam_locomotive: "unicode/1f682.png?v8",
            stethoscope: "unicode/1fa7a.png?v8",
            stew: "unicode/1f372.png?v8",
            stop_button: "unicode/23f9.png?v8",
            stop_sign: "unicode/1f6d1.png?v8",
            stopwatch: "unicode/23f1.png?v8",
            straight_ruler: "unicode/1f4cf.png?v8",
            strawberry: "unicode/1f353.png?v8",
            stuck_out_tongue: "unicode/1f61b.png?v8",
            stuck_out_tongue_closed_eyes: "unicode/1f61d.png?v8",
            stuck_out_tongue_winking_eye: "unicode/1f61c.png?v8",
            student: "unicode/1f9d1-1f393.png?v8",
            studio_microphone: "unicode/1f399.png?v8",
            stuffed_flatbread: "unicode/1f959.png?v8",
            sudan: "unicode/1f1f8-1f1e9.png?v8",
            sun_behind_large_cloud: "unicode/1f325.png?v8",
            sun_behind_rain_cloud: "unicode/1f326.png?v8",
            sun_behind_small_cloud: "unicode/1f324.png?v8",
            sun_with_face: "unicode/1f31e.png?v8",
            sunflower: "unicode/1f33b.png?v8",
            sunglasses: "unicode/1f60e.png?v8",
            sunny: "unicode/2600.png?v8",
            sunrise: "unicode/1f305.png?v8",
            sunrise_over_mountains: "unicode/1f304.png?v8",
            superhero: "unicode/1f9b8.png?v8",
            superhero_man: "unicode/1f9b8-2642.png?v8",
            superhero_woman: "unicode/1f9b8-2640.png?v8",
            supervillain: "unicode/1f9b9.png?v8",
            supervillain_man: "unicode/1f9b9-2642.png?v8",
            supervillain_woman: "unicode/1f9b9-2640.png?v8",
            surfer: "unicode/1f3c4.png?v8",
            surfing_man: "unicode/1f3c4-2642.png?v8",
            surfing_woman: "unicode/1f3c4-2640.png?v8",
            suriname: "unicode/1f1f8-1f1f7.png?v8",
            sushi: "unicode/1f363.png?v8",
            suspect: "suspect.png?v8",
            suspension_railway: "unicode/1f69f.png?v8",
            svalbard_jan_mayen: "unicode/1f1f8-1f1ef.png?v8",
            swan: "unicode/1f9a2.png?v8",
            swaziland: "unicode/1f1f8-1f1ff.png?v8",
            sweat: "unicode/1f613.png?v8",
            sweat_drops: "unicode/1f4a6.png?v8",
            sweat_smile: "unicode/1f605.png?v8",
            sweden: "unicode/1f1f8-1f1ea.png?v8",
            sweet_potato: "unicode/1f360.png?v8",
            swim_brief: "unicode/1fa72.png?v8",
            swimmer: "unicode/1f3ca.png?v8",
            swimming_man: "unicode/1f3ca-2642.png?v8",
            swimming_woman: "unicode/1f3ca-2640.png?v8",
            switzerland: "unicode/1f1e8-1f1ed.png?v8",
            symbols: "unicode/1f523.png?v8",
            synagogue: "unicode/1f54d.png?v8",
            syria: "unicode/1f1f8-1f1fe.png?v8",
            syringe: "unicode/1f489.png?v8",
            "t-rex": "unicode/1f996.png?v8",
            taco: "unicode/1f32e.png?v8",
            tada: "unicode/1f389.png?v8",
            taiwan: "unicode/1f1f9-1f1fc.png?v8",
            tajikistan: "unicode/1f1f9-1f1ef.png?v8",
            takeout_box: "unicode/1f961.png?v8",
            tamale: "unicode/1fad4.png?v8",
            tanabata_tree: "unicode/1f38b.png?v8",
            tangerine: "unicode/1f34a.png?v8",
            tanzania: "unicode/1f1f9-1f1ff.png?v8",
            taurus: "unicode/2649.png?v8",
            taxi: "unicode/1f695.png?v8",
            tea: "unicode/1f375.png?v8",
            teacher: "unicode/1f9d1-1f3eb.png?v8",
            teapot: "unicode/1fad6.png?v8",
            technologist: "unicode/1f9d1-1f4bb.png?v8",
            teddy_bear: "unicode/1f9f8.png?v8",
            telephone: "unicode/260e.png?v8",
            telephone_receiver: "unicode/1f4de.png?v8",
            telescope: "unicode/1f52d.png?v8",
            tennis: "unicode/1f3be.png?v8",
            tent: "unicode/26fa.png?v8",
            test_tube: "unicode/1f9ea.png?v8",
            thailand: "unicode/1f1f9-1f1ed.png?v8",
            thermometer: "unicode/1f321.png?v8",
            thinking: "unicode/1f914.png?v8",
            thong_sandal: "unicode/1fa74.png?v8",
            thought_balloon: "unicode/1f4ad.png?v8",
            thread: "unicode/1f9f5.png?v8",
            three: "unicode/0033-20e3.png?v8",
            thumbsdown: "unicode/1f44e.png?v8",
            thumbsup: "unicode/1f44d.png?v8",
            ticket: "unicode/1f3ab.png?v8",
            tickets: "unicode/1f39f.png?v8",
            tiger: "unicode/1f42f.png?v8",
            tiger2: "unicode/1f405.png?v8",
            timer_clock: "unicode/23f2.png?v8",
            timor_leste: "unicode/1f1f9-1f1f1.png?v8",
            tipping_hand_man: "unicode/1f481-2642.png?v8",
            tipping_hand_person: "unicode/1f481.png?v8",
            tipping_hand_woman: "unicode/1f481-2640.png?v8",
            tired_face: "unicode/1f62b.png?v8",
            tm: "unicode/2122.png?v8",
            togo: "unicode/1f1f9-1f1ec.png?v8",
            toilet: "unicode/1f6bd.png?v8",
            tokelau: "unicode/1f1f9-1f1f0.png?v8",
            tokyo_tower: "unicode/1f5fc.png?v8",
            tomato: "unicode/1f345.png?v8",
            tonga: "unicode/1f1f9-1f1f4.png?v8",
            tongue: "unicode/1f445.png?v8",
            toolbox: "unicode/1f9f0.png?v8",
            tooth: "unicode/1f9b7.png?v8",
            toothbrush: "unicode/1faa5.png?v8",
            top: "unicode/1f51d.png?v8",
            tophat: "unicode/1f3a9.png?v8",
            tornado: "unicode/1f32a.png?v8",
            tr: "unicode/1f1f9-1f1f7.png?v8",
            trackball: "unicode/1f5b2.png?v8",
            tractor: "unicode/1f69c.png?v8",
            traffic_light: "unicode/1f6a5.png?v8",
            train: "unicode/1f68b.png?v8",
            train2: "unicode/1f686.png?v8",
            tram: "unicode/1f68a.png?v8",
            transgender_flag: "unicode/1f3f3-26a7.png?v8",
            transgender_symbol: "unicode/26a7.png?v8",
            triangular_flag_on_post: "unicode/1f6a9.png?v8",
            triangular_ruler: "unicode/1f4d0.png?v8",
            trident: "unicode/1f531.png?v8",
            trinidad_tobago: "unicode/1f1f9-1f1f9.png?v8",
            tristan_da_cunha: "unicode/1f1f9-1f1e6.png?v8",
            triumph: "unicode/1f624.png?v8",
            troll: "unicode/1f9cc.png?v8",
            trolleybus: "unicode/1f68e.png?v8",
            trollface: "trollface.png?v8",
            trophy: "unicode/1f3c6.png?v8",
            tropical_drink: "unicode/1f379.png?v8",
            tropical_fish: "unicode/1f420.png?v8",
            truck: "unicode/1f69a.png?v8",
            trumpet: "unicode/1f3ba.png?v8",
            tshirt: "unicode/1f455.png?v8",
            tulip: "unicode/1f337.png?v8",
            tumbler_glass: "unicode/1f943.png?v8",
            tunisia: "unicode/1f1f9-1f1f3.png?v8",
            turkey: "unicode/1f983.png?v8",
            turkmenistan: "unicode/1f1f9-1f1f2.png?v8",
            turks_caicos_islands: "unicode/1f1f9-1f1e8.png?v8",
            turtle: "unicode/1f422.png?v8",
            tuvalu: "unicode/1f1f9-1f1fb.png?v8",
            tv: "unicode/1f4fa.png?v8",
            twisted_rightwards_arrows: "unicode/1f500.png?v8",
            two: "unicode/0032-20e3.png?v8",
            two_hearts: "unicode/1f495.png?v8",
            two_men_holding_hands: "unicode/1f46c.png?v8",
            two_women_holding_hands: "unicode/1f46d.png?v8",
            u5272: "unicode/1f239.png?v8",
            u5408: "unicode/1f234.png?v8",
            u55b6: "unicode/1f23a.png?v8",
            u6307: "unicode/1f22f.png?v8",
            u6708: "unicode/1f237.png?v8",
            u6709: "unicode/1f236.png?v8",
            u6e80: "unicode/1f235.png?v8",
            u7121: "unicode/1f21a.png?v8",
            u7533: "unicode/1f238.png?v8",
            u7981: "unicode/1f232.png?v8",
            u7a7a: "unicode/1f233.png?v8",
            uganda: "unicode/1f1fa-1f1ec.png?v8",
            uk: "unicode/1f1ec-1f1e7.png?v8",
            ukraine: "unicode/1f1fa-1f1e6.png?v8",
            umbrella: "unicode/2614.png?v8",
            unamused: "unicode/1f612.png?v8",
            underage: "unicode/1f51e.png?v8",
            unicorn: "unicode/1f984.png?v8",
            united_arab_emirates: "unicode/1f1e6-1f1ea.png?v8",
            united_nations: "unicode/1f1fa-1f1f3.png?v8",
            unlock: "unicode/1f513.png?v8",
            up: "unicode/1f199.png?v8",
            upside_down_face: "unicode/1f643.png?v8",
            uruguay: "unicode/1f1fa-1f1fe.png?v8",
            us: "unicode/1f1fa-1f1f8.png?v8",
            us_outlying_islands: "unicode/1f1fa-1f1f2.png?v8",
            us_virgin_islands: "unicode/1f1fb-1f1ee.png?v8",
            uzbekistan: "unicode/1f1fa-1f1ff.png?v8",
            v: "unicode/270c.png?v8",
            vampire: "unicode/1f9db.png?v8",
            vampire_man: "unicode/1f9db-2642.png?v8",
            vampire_woman: "unicode/1f9db-2640.png?v8",
            vanuatu: "unicode/1f1fb-1f1fa.png?v8",
            vatican_city: "unicode/1f1fb-1f1e6.png?v8",
            venezuela: "unicode/1f1fb-1f1ea.png?v8",
            vertical_traffic_light: "unicode/1f6a6.png?v8",
            vhs: "unicode/1f4fc.png?v8",
            vibration_mode: "unicode/1f4f3.png?v8",
            video_camera: "unicode/1f4f9.png?v8",
            video_game: "unicode/1f3ae.png?v8",
            vietnam: "unicode/1f1fb-1f1f3.png?v8",
            violin: "unicode/1f3bb.png?v8",
            virgo: "unicode/264d.png?v8",
            volcano: "unicode/1f30b.png?v8",
            volleyball: "unicode/1f3d0.png?v8",
            vomiting_face: "unicode/1f92e.png?v8",
            vs: "unicode/1f19a.png?v8",
            vulcan_salute: "unicode/1f596.png?v8",
            waffle: "unicode/1f9c7.png?v8",
            wales: "unicode/1f3f4-e0067-e0062-e0077-e006c-e0073-e007f.png?v8",
            walking: "unicode/1f6b6.png?v8",
            walking_man: "unicode/1f6b6-2642.png?v8",
            walking_woman: "unicode/1f6b6-2640.png?v8",
            wallis_futuna: "unicode/1f1fc-1f1eb.png?v8",
            waning_crescent_moon: "unicode/1f318.png?v8",
            waning_gibbous_moon: "unicode/1f316.png?v8",
            warning: "unicode/26a0.png?v8",
            wastebasket: "unicode/1f5d1.png?v8",
            watch: "unicode/231a.png?v8",
            water_buffalo: "unicode/1f403.png?v8",
            water_polo: "unicode/1f93d.png?v8",
            watermelon: "unicode/1f349.png?v8",
            wave: "unicode/1f44b.png?v8",
            wavy_dash: "unicode/3030.png?v8",
            waxing_crescent_moon: "unicode/1f312.png?v8",
            waxing_gibbous_moon: "unicode/1f314.png?v8",
            wc: "unicode/1f6be.png?v8",
            weary: "unicode/1f629.png?v8",
            wedding: "unicode/1f492.png?v8",
            weight_lifting: "unicode/1f3cb.png?v8",
            weight_lifting_man: "unicode/1f3cb-2642.png?v8",
            weight_lifting_woman: "unicode/1f3cb-2640.png?v8",
            western_sahara: "unicode/1f1ea-1f1ed.png?v8",
            whale: "unicode/1f433.png?v8",
            whale2: "unicode/1f40b.png?v8",
            wheel: "unicode/1f6de.png?v8",
            wheel_of_dharma: "unicode/2638.png?v8",
            wheelchair: "unicode/267f.png?v8",
            white_check_mark: "unicode/2705.png?v8",
            white_circle: "unicode/26aa.png?v8",
            white_flag: "unicode/1f3f3.png?v8",
            white_flower: "unicode/1f4ae.png?v8",
            white_haired_man: "unicode/1f468-1f9b3.png?v8",
            white_haired_woman: "unicode/1f469-1f9b3.png?v8",
            white_heart: "unicode/1f90d.png?v8",
            white_large_square: "unicode/2b1c.png?v8",
            white_medium_small_square: "unicode/25fd.png?v8",
            white_medium_square: "unicode/25fb.png?v8",
            white_small_square: "unicode/25ab.png?v8",
            white_square_button: "unicode/1f533.png?v8",
            wilted_flower: "unicode/1f940.png?v8",
            wind_chime: "unicode/1f390.png?v8",
            wind_face: "unicode/1f32c.png?v8",
            window: "unicode/1fa9f.png?v8",
            wine_glass: "unicode/1f377.png?v8",
            wing: "unicode/1fabd.png?v8",
            wink: "unicode/1f609.png?v8",
            wireless: "unicode/1f6dc.png?v8",
            wolf: "unicode/1f43a.png?v8",
            woman: "unicode/1f469.png?v8",
            woman_artist: "unicode/1f469-1f3a8.png?v8",
            woman_astronaut: "unicode/1f469-1f680.png?v8",
            woman_beard: "unicode/1f9d4-2640.png?v8",
            woman_cartwheeling: "unicode/1f938-2640.png?v8",
            woman_cook: "unicode/1f469-1f373.png?v8",
            woman_dancing: "unicode/1f483.png?v8",
            woman_facepalming: "unicode/1f926-2640.png?v8",
            woman_factory_worker: "unicode/1f469-1f3ed.png?v8",
            woman_farmer: "unicode/1f469-1f33e.png?v8",
            woman_feeding_baby: "unicode/1f469-1f37c.png?v8",
            woman_firefighter: "unicode/1f469-1f692.png?v8",
            woman_health_worker: "unicode/1f469-2695.png?v8",
            woman_in_manual_wheelchair: "unicode/1f469-1f9bd.png?v8",
            woman_in_motorized_wheelchair: "unicode/1f469-1f9bc.png?v8",
            woman_in_tuxedo: "unicode/1f935-2640.png?v8",
            woman_judge: "unicode/1f469-2696.png?v8",
            woman_juggling: "unicode/1f939-2640.png?v8",
            woman_mechanic: "unicode/1f469-1f527.png?v8",
            woman_office_worker: "unicode/1f469-1f4bc.png?v8",
            woman_pilot: "unicode/1f469-2708.png?v8",
            woman_playing_handball: "unicode/1f93e-2640.png?v8",
            woman_playing_water_polo: "unicode/1f93d-2640.png?v8",
            woman_scientist: "unicode/1f469-1f52c.png?v8",
            woman_shrugging: "unicode/1f937-2640.png?v8",
            woman_singer: "unicode/1f469-1f3a4.png?v8",
            woman_student: "unicode/1f469-1f393.png?v8",
            woman_teacher: "unicode/1f469-1f3eb.png?v8",
            woman_technologist: "unicode/1f469-1f4bb.png?v8",
            woman_with_headscarf: "unicode/1f9d5.png?v8",
            woman_with_probing_cane: "unicode/1f469-1f9af.png?v8",
            woman_with_turban: "unicode/1f473-2640.png?v8",
            woman_with_veil: "unicode/1f470-2640.png?v8",
            womans_clothes: "unicode/1f45a.png?v8",
            womans_hat: "unicode/1f452.png?v8",
            women_wrestling: "unicode/1f93c-2640.png?v8",
            womens: "unicode/1f6ba.png?v8",
            wood: "unicode/1fab5.png?v8",
            woozy_face: "unicode/1f974.png?v8",
            world_map: "unicode/1f5fa.png?v8",
            worm: "unicode/1fab1.png?v8",
            worried: "unicode/1f61f.png?v8",
            wrench: "unicode/1f527.png?v8",
            wrestling: "unicode/1f93c.png?v8",
            writing_hand: "unicode/270d.png?v8",
            x: "unicode/274c.png?v8",
            x_ray: "unicode/1fa7b.png?v8",
            yarn: "unicode/1f9f6.png?v8",
            yawning_face: "unicode/1f971.png?v8",
            yellow_circle: "unicode/1f7e1.png?v8",
            yellow_heart: "unicode/1f49b.png?v8",
            yellow_square: "unicode/1f7e8.png?v8",
            yemen: "unicode/1f1fe-1f1ea.png?v8",
            yen: "unicode/1f4b4.png?v8",
            yin_yang: "unicode/262f.png?v8",
            yo_yo: "unicode/1fa80.png?v8",
            yum: "unicode/1f60b.png?v8",
            zambia: "unicode/1f1ff-1f1f2.png?v8",
            zany_face: "unicode/1f92a.png?v8",
            zap: "unicode/26a1.png?v8",
            zebra: "unicode/1f993.png?v8",
            zero: "unicode/0030-20e3.png?v8",
            zimbabwe: "unicode/1f1ff-1f1fc.png?v8",
            zipper_mouth_face: "unicode/1f910.png?v8",
            zombie: "unicode/1f9df.png?v8",
            zombie_man: "unicode/1f9df-2642.png?v8",
            zombie_woman: "unicode/1f9df-2640.png?v8",
            zzz: "unicode/1f4a4.png?v8"
        }
    };
    function replaceEmojiShorthand(m, $1, useNativeEmoji) {
        const emojiMatch = emojiData.data[$1];
        let result = m;
        if (emojiMatch) {
            if (useNativeEmoji && /unicode/.test(emojiMatch)) {
                const emojiUnicode = emojiMatch.replace("unicode/", "").replace(/\.png.*/, "").split("-").map((u => `&#x${u};`)).join("&zwj;").concat("&#xFE0E;");
                result = `<span class="emoji">${emojiUnicode}</span>`;
            } else {
                result = `<img src="${emojiData.baseURL}${emojiMatch}.png" alt="${$1}" class="emoji" loading="lazy">`;
            }
        }
        return result;
    }
    function emojify(text, useNativeEmoji) {
        return text.replace(/<(code|pre|script|template)[^>]*?>[\s\S]+?<\/(code|pre|script|template)>/g, (m => m.replace(/:/g, "__colon__"))).replace(/<!--[\s\S]+?-->/g, (m => m.replace(/:/g, "__colon__"))).replace(/([a-z]{2,}:)?\/\/[^\s'">)]+/gi, (m => m.replace(/:/g, "__colon__"))).replace(/:([a-z0-9_\-+]+?):/g, ((m, $1) => replaceEmojiShorthand(m, $1, useNativeEmoji))).replace(/__colon__/g, ":");
    }
    function getAndRemoveConfig() {
        let str = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : "";
        const config = {};
        if (str) {
            str = str.replace(/^('|")/, "").replace(/('|")$/, "").replace(/(?:^|\s):([\w-]+:?)=?([\w-%]+)?/g, ((m, key, value) => {
                if (key.indexOf(":") === -1) {
                    config[key] = value && value.replace(/&quot;/g, "") || true;
                    return "";
                }
                return m;
            })).trim();
        }
        return {
            str: str,
            config: config
        };
    }
    function removeAtag() {
        let str = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : "";
        return str.replace(/(<\/?a.*?>)/gi, "");
    }
    function getAndRemoveDocisfyIgnoreConfig() {
        let content = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : "";
        let ignoreAllSubs, ignoreSubHeading;
        if (/<!-- {docsify-ignore} -->/g.test(content)) {
            content = content.replace("\x3c!-- {docsify-ignore} --\x3e", "");
            ignoreSubHeading = true;
        }
        if (/{docsify-ignore}/g.test(content)) {
            content = content.replace("{docsify-ignore}", "");
            ignoreSubHeading = true;
        }
        if (/<!-- {docsify-ignore-all} -->/g.test(content)) {
            content = content.replace("\x3c!-- {docsify-ignore-all} --\x3e", "");
            ignoreAllSubs = true;
        }
        if (/{docsify-ignore-all}/g.test(content)) {
            content = content.replace("{docsify-ignore-all}", "");
            ignoreAllSubs = true;
        }
        return {
            content: content,
            ignoreAllSubs: ignoreAllSubs,
            ignoreSubHeading: ignoreSubHeading
        };
    }
    const imageCompiler = _ref => {
        let {renderer: renderer, contentBase: contentBase, router: router} = _ref;
        return renderer.image = (href, title, text) => {
            let url = href;
            const attrs = [];
            const {str: str, config: config} = getAndRemoveConfig(title);
            title = str;
            if (config["no-zoom"]) {
                attrs.push("data-no-zoom");
            }
            if (title) {
                attrs.push(`title="${title}"`);
            }
            if (config.size) {
                const [width, height] = config.size.split("x");
                if (height) {
                    attrs.push(`width="${width}" height="${height}"`);
                } else {
                    attrs.push(`width="${width}"`);
                }
            }
            if (config.class) {
                attrs.push(`class="${config.class}"`);
            }
            if (config.id) {
                attrs.push(`id="${config.id}"`);
            }
            if (!isAbsolutePath(href)) {
                url = getPath(contentBase, getParentPath(router.getCurrentPath()), href);
            }
            return `<img src="${url}" data-origin="${href}" alt="${text}" ${attrs.join(" ")} />`;
        };
    };
    var commonjsGlobal = typeof globalThis !== "undefined" ? globalThis : typeof window !== "undefined" ? window : typeof global !== "undefined" ? global : typeof self !== "undefined" ? self : {};
    function getDefaultExportFromCjs(x) {
        return x && x.__esModule && Object.prototype.hasOwnProperty.call(x, "default") ? x["default"] : x;
    }
    var prism$1 = {
        exports: {}
    };
    (function(module) {
        var _self = typeof window !== "undefined" ? window : typeof WorkerGlobalScope !== "undefined" && self instanceof WorkerGlobalScope ? self : {};
        var Prism = function(_self) {
            var lang = /(?:^|\s)lang(?:uage)?-([\w-]+)(?=\s|$)/i;
            var uniqueId = 0;
            var plainTextGrammar = {};
            var _ = {
                manual: _self.Prism && _self.Prism.manual,
                disableWorkerMessageHandler: _self.Prism && _self.Prism.disableWorkerMessageHandler,
                util: {
                    encode: function encode(tokens) {
                        if (tokens instanceof Token) {
                            return new Token(tokens.type, encode(tokens.content), tokens.alias);
                        } else if (Array.isArray(tokens)) {
                            return tokens.map(encode);
                        } else {
                            return tokens.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/\u00a0/g, " ");
                        }
                    },
                    type: function(o) {
                        return Object.prototype.toString.call(o).slice(8, -1);
                    },
                    objId: function(obj) {
                        if (!obj["__id"]) {
                            Object.defineProperty(obj, "__id", {
                                value: ++uniqueId
                            });
                        }
                        return obj["__id"];
                    },
                    clone: function deepClone(o, visited) {
                        visited = visited || {};
                        var clone;
                        var id;
                        switch (_.util.type(o)) {
                            case "Object":
                                id = _.util.objId(o);
                                if (visited[id]) {
                                    return visited[id];
                                }
                                clone = {};
                                visited[id] = clone;
                                for (var key in o) {
                                    if (o.hasOwnProperty(key)) {
                                        clone[key] = deepClone(o[key], visited);
                                    }
                                }
                                return clone;

                            case "Array":
                                id = _.util.objId(o);
                                if (visited[id]) {
                                    return visited[id];
                                }
                                clone = [];
                                visited[id] = clone;
                                o.forEach((function(v, i) {
                                    clone[i] = deepClone(v, visited);
                                }));
                                return clone;

                            default:
                                return o;
                        }
                    },
                    getLanguage: function(element) {
                        while (element) {
                            var m = lang.exec(element.className);
                            if (m) {
                                return m[1].toLowerCase();
                            }
                            element = element.parentElement;
                        }
                        return "none";
                    },
                    setLanguage: function(element, language) {
                        element.className = element.className.replace(RegExp(lang, "gi"), "");
                        element.classList.add("language-" + language);
                    },
                    currentScript: function() {
                        if (typeof document === "undefined") {
                            return null;
                        }
                        if ("currentScript" in document && 1 < 2) {
                            return document.currentScript;
                        }
                        try {
                            throw new Error;
                        } catch (err) {
                            var src = (/at [^(\r\n]*\((.*):[^:]+:[^:]+\)$/i.exec(err.stack) || [])[1];
                            if (src) {
                                var scripts = document.getElementsByTagName("script");
                                for (var i in scripts) {
                                    if (scripts[i].src == src) {
                                        return scripts[i];
                                    }
                                }
                            }
                            return null;
                        }
                    },
                    isActive: function(element, className, defaultActivation) {
                        var no = "no-" + className;
                        while (element) {
                            var classList = element.classList;
                            if (classList.contains(className)) {
                                return true;
                            }
                            if (classList.contains(no)) {
                                return false;
                            }
                            element = element.parentElement;
                        }
                        return !!defaultActivation;
                    }
                },
                languages: {
                    plain: plainTextGrammar,
                    plaintext: plainTextGrammar,
                    text: plainTextGrammar,
                    txt: plainTextGrammar,
                    extend: function(id, redef) {
                        var lang = _.util.clone(_.languages[id]);
                        for (var key in redef) {
                            lang[key] = redef[key];
                        }
                        return lang;
                    },
                    insertBefore: function(inside, before, insert, root) {
                        root = root || _.languages;
                        var grammar = root[inside];
                        var ret = {};
                        for (var token in grammar) {
                            if (grammar.hasOwnProperty(token)) {
                                if (token == before) {
                                    for (var newToken in insert) {
                                        if (insert.hasOwnProperty(newToken)) {
                                            ret[newToken] = insert[newToken];
                                        }
                                    }
                                }
                                if (!insert.hasOwnProperty(token)) {
                                    ret[token] = grammar[token];
                                }
                            }
                        }
                        var old = root[inside];
                        root[inside] = ret;
                        _.languages.DFS(_.languages, (function(key, value) {
                            if (value === old && key != inside) {
                                this[key] = ret;
                            }
                        }));
                        return ret;
                    },
                    DFS: function DFS(o, callback, type, visited) {
                        visited = visited || {};
                        var objId = _.util.objId;
                        for (var i in o) {
                            if (o.hasOwnProperty(i)) {
                                callback.call(o, i, o[i], type || i);
                                var property = o[i];
                                var propertyType = _.util.type(property);
                                if (propertyType === "Object" && !visited[objId(property)]) {
                                    visited[objId(property)] = true;
                                    DFS(property, callback, null, visited);
                                } else if (propertyType === "Array" && !visited[objId(property)]) {
                                    visited[objId(property)] = true;
                                    DFS(property, callback, i, visited);
                                }
                            }
                        }
                    }
                },
                plugins: {},
                highlightAll: function(async, callback) {
                    _.highlightAllUnder(document, async, callback);
                },
                highlightAllUnder: function(container, async, callback) {
                    var env = {
                        callback: callback,
                        container: container,
                        selector: 'code[class*="language-"], [class*="language-"] code, code[class*="lang-"], [class*="lang-"] code'
                    };
                    _.hooks.run("before-highlightall", env);
                    env.elements = Array.prototype.slice.apply(env.container.querySelectorAll(env.selector));
                    _.hooks.run("before-all-elements-highlight", env);
                    for (var i = 0, element; element = env.elements[i++]; ) {
                        _.highlightElement(element, async === true, env.callback);
                    }
                },
                highlightElement: function(element, async, callback) {
                    var language = _.util.getLanguage(element);
                    var grammar = _.languages[language];
                    _.util.setLanguage(element, language);
                    var parent = element.parentElement;
                    if (parent && parent.nodeName.toLowerCase() === "pre") {
                        _.util.setLanguage(parent, language);
                    }
                    var code = element.textContent;
                    var env = {
                        element: element,
                        language: language,
                        grammar: grammar,
                        code: code
                    };
                    function insertHighlightedCode(highlightedCode) {
                        env.highlightedCode = highlightedCode;
                        _.hooks.run("before-insert", env);
                        env.element.innerHTML = env.highlightedCode;
                        _.hooks.run("after-highlight", env);
                        _.hooks.run("complete", env);
                        callback && callback.call(env.element);
                    }
                    _.hooks.run("before-sanity-check", env);
                    parent = env.element.parentElement;
                    if (parent && parent.nodeName.toLowerCase() === "pre" && !parent.hasAttribute("tabindex")) {
                        parent.setAttribute("tabindex", "0");
                    }
                    if (!env.code) {
                        _.hooks.run("complete", env);
                        callback && callback.call(env.element);
                        return;
                    }
                    _.hooks.run("before-highlight", env);
                    if (!env.grammar) {
                        insertHighlightedCode(_.util.encode(env.code));
                        return;
                    }
                    if (async && _self.Worker) {
                        var worker = new Worker(_.filename);
                        worker.onmessage = function(evt) {
                            insertHighlightedCode(evt.data);
                        };
                        worker.postMessage(JSON.stringify({
                            language: env.language,
                            code: env.code,
                            immediateClose: true
                        }));
                    } else {
                        insertHighlightedCode(_.highlight(env.code, env.grammar, env.language));
                    }
                },
                highlight: function(text, grammar, language) {
                    var env = {
                        code: text,
                        grammar: grammar,
                        language: language
                    };
                    _.hooks.run("before-tokenize", env);
                    if (!env.grammar) {
                        throw new Error('The language "' + env.language + '" has no grammar.');
                    }
                    env.tokens = _.tokenize(env.code, env.grammar);
                    _.hooks.run("after-tokenize", env);
                    return Token.stringify(_.util.encode(env.tokens), env.language);
                },
                tokenize: function(text, grammar) {
                    var rest = grammar.rest;
                    if (rest) {
                        for (var token in rest) {
                            grammar[token] = rest[token];
                        }
                        delete grammar.rest;
                    }
                    var tokenList = new LinkedList;
                    addAfter(tokenList, tokenList.head, text);
                    matchGrammar(text, tokenList, grammar, tokenList.head, 0);
                    return toArray(tokenList);
                },
                hooks: {
                    all: {},
                    add: function(name, callback) {
                        var hooks = _.hooks.all;
                        hooks[name] = hooks[name] || [];
                        hooks[name].push(callback);
                    },
                    run: function(name, env) {
                        var callbacks = _.hooks.all[name];
                        if (!callbacks || !callbacks.length) {
                            return;
                        }
                        for (var i = 0, callback; callback = callbacks[i++]; ) {
                            callback(env);
                        }
                    }
                },
                Token: Token
            };
            _self.Prism = _;
            function Token(type, content, alias, matchedStr) {
                this.type = type;
                this.content = content;
                this.alias = alias;
                this.length = (matchedStr || "").length | 0;
            }
            Token.stringify = function stringify(o, language) {
                if (typeof o == "string") {
                    return o;
                }
                if (Array.isArray(o)) {
                    var s = "";
                    o.forEach((function(e) {
                        s += stringify(e, language);
                    }));
                    return s;
                }
                var env = {
                    type: o.type,
                    content: stringify(o.content, language),
                    tag: "span",
                    classes: [ "token", o.type ],
                    attributes: {},
                    language: language
                };
                var aliases = o.alias;
                if (aliases) {
                    if (Array.isArray(aliases)) {
                        Array.prototype.push.apply(env.classes, aliases);
                    } else {
                        env.classes.push(aliases);
                    }
                }
                _.hooks.run("wrap", env);
                var attributes = "";
                for (var name in env.attributes) {
                    attributes += " " + name + '="' + (env.attributes[name] || "").replace(/"/g, "&quot;") + '"';
                }
                return "<" + env.tag + ' class="' + env.classes.join(" ") + '"' + attributes + ">" + env.content + "</" + env.tag + ">";
            };
            function matchPattern(pattern, pos, text, lookbehind) {
                pattern.lastIndex = pos;
                var match = pattern.exec(text);
                if (match && lookbehind && match[1]) {
                    var lookbehindLength = match[1].length;
                    match.index += lookbehindLength;
                    match[0] = match[0].slice(lookbehindLength);
                }
                return match;
            }
            function matchGrammar(text, tokenList, grammar, startNode, startPos, rematch) {
                for (var token in grammar) {
                    if (!grammar.hasOwnProperty(token) || !grammar[token]) {
                        continue;
                    }
                    var patterns = grammar[token];
                    patterns = Array.isArray(patterns) ? patterns : [ patterns ];
                    for (var j = 0; j < patterns.length; ++j) {
                        if (rematch && rematch.cause == token + "," + j) {
                            return;
                        }
                        var patternObj = patterns[j];
                        var inside = patternObj.inside;
                        var lookbehind = !!patternObj.lookbehind;
                        var greedy = !!patternObj.greedy;
                        var alias = patternObj.alias;
                        if (greedy && !patternObj.pattern.global) {
                            var flags = patternObj.pattern.toString().match(/[imsuy]*$/)[0];
                            patternObj.pattern = RegExp(patternObj.pattern.source, flags + "g");
                        }
                        var pattern = patternObj.pattern || patternObj;
                        for (var currentNode = startNode.next, pos = startPos; currentNode !== tokenList.tail; pos += currentNode.value.length,
                            currentNode = currentNode.next) {
                            if (rematch && pos >= rematch.reach) {
                                break;
                            }
                            var str = currentNode.value;
                            if (tokenList.length > text.length) {
                                return;
                            }
                            if (str instanceof Token) {
                                continue;
                            }
                            var removeCount = 1;
                            var match;
                            if (greedy) {
                                match = matchPattern(pattern, pos, text, lookbehind);
                                if (!match || match.index >= text.length) {
                                    break;
                                }
                                var from = match.index;
                                var to = match.index + match[0].length;
                                var p = pos;
                                p += currentNode.value.length;
                                while (from >= p) {
                                    currentNode = currentNode.next;
                                    p += currentNode.value.length;
                                }
                                p -= currentNode.value.length;
                                pos = p;
                                if (currentNode.value instanceof Token) {
                                    continue;
                                }
                                for (var k = currentNode; k !== tokenList.tail && (p < to || typeof k.value === "string"); k = k.next) {
                                    removeCount++;
                                    p += k.value.length;
                                }
                                removeCount--;
                                str = text.slice(pos, p);
                                match.index -= pos;
                            } else {
                                match = matchPattern(pattern, 0, str, lookbehind);
                                if (!match) {
                                    continue;
                                }
                            }
                            var from = match.index;
                            var matchStr = match[0];
                            var before = str.slice(0, from);
                            var after = str.slice(from + matchStr.length);
                            var reach = pos + str.length;
                            if (rematch && reach > rematch.reach) {
                                rematch.reach = reach;
                            }
                            var removeFrom = currentNode.prev;
                            if (before) {
                                removeFrom = addAfter(tokenList, removeFrom, before);
                                pos += before.length;
                            }
                            removeRange(tokenList, removeFrom, removeCount);
                            var wrapped = new Token(token, inside ? _.tokenize(matchStr, inside) : matchStr, alias, matchStr);
                            currentNode = addAfter(tokenList, removeFrom, wrapped);
                            if (after) {
                                addAfter(tokenList, currentNode, after);
                            }
                            if (removeCount > 1) {
                                var nestedRematch = {
                                    cause: token + "," + j,
                                    reach: reach
                                };
                                matchGrammar(text, tokenList, grammar, currentNode.prev, pos, nestedRematch);
                                if (rematch && nestedRematch.reach > rematch.reach) {
                                    rematch.reach = nestedRematch.reach;
                                }
                            }
                        }
                    }
                }
            }
            function LinkedList() {
                var head = {
                    value: null,
                    prev: null,
                    next: null
                };
                var tail = {
                    value: null,
                    prev: head,
                    next: null
                };
                head.next = tail;
                this.head = head;
                this.tail = tail;
                this.length = 0;
            }
            function addAfter(list, node, value) {
                var next = node.next;
                var newNode = {
                    value: value,
                    prev: node,
                    next: next
                };
                node.next = newNode;
                next.prev = newNode;
                list.length++;
                return newNode;
            }
            function removeRange(list, node, count) {
                var next = node.next;
                for (var i = 0; i < count && next !== list.tail; i++) {
                    next = next.next;
                }
                node.next = next;
                next.prev = node;
                list.length -= i;
            }
            function toArray(list) {
                var array = [];
                var node = list.head.next;
                while (node !== list.tail) {
                    array.push(node.value);
                    node = node.next;
                }
                return array;
            }
            if (!_self.document) {
                if (!_self.addEventListener) {
                    return _;
                }
                if (!_.disableWorkerMessageHandler) {
                    _self.addEventListener("message", (function(evt) {
                        var message = JSON.parse(evt.data);
                        var lang = message.language;
                        var code = message.code;
                        var immediateClose = message.immediateClose;
                        _self.postMessage(_.highlight(code, _.languages[lang], lang));
                        if (immediateClose) {
                            _self.close();
                        }
                    }), false);
                }
                return _;
            }
            var script = _.util.currentScript();
            if (script) {
                _.filename = script.src;
                if (script.hasAttribute("data-manual")) {
                    _.manual = true;
                }
            }
            function highlightAutomaticallyCallback() {
                if (!_.manual) {
                    _.highlightAll();
                }
            }
            if (!_.manual) {
                var readyState = document.readyState;
                if (readyState === "loading" || readyState === "interactive" && script && script.defer) {
                    document.addEventListener("DOMContentLoaded", highlightAutomaticallyCallback);
                } else {
                    if (window.requestAnimationFrame) {
                        window.requestAnimationFrame(highlightAutomaticallyCallback);
                    } else {
                        window.setTimeout(highlightAutomaticallyCallback, 16);
                    }
                }
            }
            return _;
        }(_self);
        if (module.exports) {
            module.exports = Prism;
        }
        if (typeof commonjsGlobal !== "undefined") {
            commonjsGlobal.Prism = Prism;
        }
        Prism.languages.markup = {
            comment: {
                pattern: /<!--(?:(?!<!--)[\s\S])*?-->/,
                greedy: true
            },
            prolog: {
                pattern: /<\?[\s\S]+?\?>/,
                greedy: true
            },
            doctype: {
                pattern: /<!DOCTYPE(?:[^>"'[\]]|"[^"]*"|'[^']*')+(?:\[(?:[^<"'\]]|"[^"]*"|'[^']*'|<(?!!--)|<!--(?:[^-]|-(?!->))*-->)*\]\s*)?>/i,
                greedy: true,
                inside: {
                    "internal-subset": {
                        pattern: /(^[^\[]*\[)[\s\S]+(?=\]>$)/,
                        lookbehind: true,
                        greedy: true,
                        inside: null
                    },
                    string: {
                        pattern: /"[^"]*"|'[^']*'/,
                        greedy: true
                    },
                    punctuation: /^<!|>$|[[\]]/,
                    "doctype-tag": /^DOCTYPE/i,
                    name: /[^\s<>'"]+/
                }
            },
            cdata: {
                pattern: /<!\[CDATA\[[\s\S]*?\]\]>/i,
                greedy: true
            },
            tag: {
                pattern: /<\/?(?!\d)[^\s>\/=$<%]+(?:\s(?:\s*[^\s>\/=]+(?:\s*=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+(?=[\s>]))|(?=[\s/>])))+)?\s*\/?>/,
                greedy: true,
                inside: {
                    tag: {
                        pattern: /^<\/?[^\s>\/]+/,
                        inside: {
                            punctuation: /^<\/?/,
                            namespace: /^[^\s>\/:]+:/
                        }
                    },
                    "special-attr": [],
                    "attr-value": {
                        pattern: /=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+)/,
                        inside: {
                            punctuation: [ {
                                pattern: /^=/,
                                alias: "attr-equals"
                            }, {
                                pattern: /^(\s*)["']|["']$/,
                                lookbehind: true
                            } ]
                        }
                    },
                    punctuation: /\/?>/,
                    "attr-name": {
                        pattern: /[^\s>\/]+/,
                        inside: {
                            namespace: /^[^\s>\/:]+:/
                        }
                    }
                }
            },
            entity: [ {
                pattern: /&[\da-z]{1,8};/i,
                alias: "named-entity"
            }, /&#x?[\da-f]{1,8};/i ]
        };
        Prism.languages.markup["tag"].inside["attr-value"].inside["entity"] = Prism.languages.markup["entity"];
        Prism.languages.markup["doctype"].inside["internal-subset"].inside = Prism.languages.markup;
        Prism.hooks.add("wrap", (function(env) {
            if (env.type === "entity") {
                env.attributes["title"] = env.content.replace(/&amp;/, "&");
            }
        }));
        Object.defineProperty(Prism.languages.markup.tag, "addInlined", {
            value: function addInlined(tagName, lang) {
                var includedCdataInside = {};
                includedCdataInside["language-" + lang] = {
                    pattern: /(^<!\[CDATA\[)[\s\S]+?(?=\]\]>$)/i,
                    lookbehind: true,
                    inside: Prism.languages[lang]
                };
                includedCdataInside["cdata"] = /^<!\[CDATA\[|\]\]>$/i;
                var inside = {
                    "included-cdata": {
                        pattern: /<!\[CDATA\[[\s\S]*?\]\]>/i,
                        inside: includedCdataInside
                    }
                };
                inside["language-" + lang] = {
                    pattern: /[\s\S]+/,
                    inside: Prism.languages[lang]
                };
                var def = {};
                def[tagName] = {
                    pattern: RegExp(/(<__[^>]*>)(?:<!\[CDATA\[(?:[^\]]|\](?!\]>))*\]\]>|(?!<!\[CDATA\[)[\s\S])*?(?=<\/__>)/.source.replace(/__/g, (function() {
                        return tagName;
                    })), "i"),
                    lookbehind: true,
                    greedy: true,
                    inside: inside
                };
                Prism.languages.insertBefore("markup", "cdata", def);
            }
        });
        Object.defineProperty(Prism.languages.markup.tag, "addAttribute", {
            value: function(attrName, lang) {
                Prism.languages.markup.tag.inside["special-attr"].push({
                    pattern: RegExp(/(^|["'\s])/.source + "(?:" + attrName + ")" + /\s*=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+(?=[\s>]))/.source, "i"),
                    lookbehind: true,
                    inside: {
                        "attr-name": /^[^\s=]+/,
                        "attr-value": {
                            pattern: /=[\s\S]+/,
                            inside: {
                                value: {
                                    pattern: /(^=\s*(["']|(?!["'])))\S[\s\S]*(?=\2$)/,
                                    lookbehind: true,
                                    alias: [ lang, "language-" + lang ],
                                    inside: Prism.languages[lang]
                                },
                                punctuation: [ {
                                    pattern: /^=/,
                                    alias: "attr-equals"
                                }, /"|'/ ]
                            }
                        }
                    }
                });
            }
        });
        Prism.languages.html = Prism.languages.markup;
        Prism.languages.mathml = Prism.languages.markup;
        Prism.languages.svg = Prism.languages.markup;
        Prism.languages.xml = Prism.languages.extend("markup", {});
        Prism.languages.ssml = Prism.languages.xml;
        Prism.languages.atom = Prism.languages.xml;
        Prism.languages.rss = Prism.languages.xml;
        (function(Prism) {
            var string = /(?:"(?:\\(?:\r\n|[\s\S])|[^"\\\r\n])*"|'(?:\\(?:\r\n|[\s\S])|[^'\\\r\n])*')/;
            Prism.languages.css = {
                comment: /\/\*[\s\S]*?\*\//,
                atrule: {
                    pattern: RegExp("@[\\w-](?:" + /[^;{\s"']|\s+(?!\s)/.source + "|" + string.source + ")*?" + /(?:;|(?=\s*\{))/.source),
                    inside: {
                        rule: /^@[\w-]+/,
                        "selector-function-argument": {
                            pattern: /(\bselector\s*\(\s*(?![\s)]))(?:[^()\s]|\s+(?![\s)])|\((?:[^()]|\([^()]*\))*\))+(?=\s*\))/,
                            lookbehind: true,
                            alias: "selector"
                        },
                        keyword: {
                            pattern: /(^|[^\w-])(?:and|not|only|or)(?![\w-])/,
                            lookbehind: true
                        }
                    }
                },
                url: {
                    pattern: RegExp("\\burl\\((?:" + string.source + "|" + /(?:[^\\\r\n()"']|\\[\s\S])*/.source + ")\\)", "i"),
                    greedy: true,
                    inside: {
                        function: /^url/i,
                        punctuation: /^\(|\)$/,
                        string: {
                            pattern: RegExp("^" + string.source + "$"),
                            alias: "url"
                        }
                    }
                },
                selector: {
                    pattern: RegExp("(^|[{}\\s])[^{}\\s](?:[^{};\"'\\s]|\\s+(?![\\s{])|" + string.source + ")*(?=\\s*\\{)"),
                    lookbehind: true
                },
                string: {
                    pattern: string,
                    greedy: true
                },
                property: {
                    pattern: /(^|[^-\w\xA0-\uFFFF])(?!\s)[-_a-z\xA0-\uFFFF](?:(?!\s)[-\w\xA0-\uFFFF])*(?=\s*:)/i,
                    lookbehind: true
                },
                important: /!important\b/i,
                function: {
                    pattern: /(^|[^-a-z0-9])[-a-z0-9]+(?=\()/i,
                    lookbehind: true
                },
                punctuation: /[(){};:,]/
            };
            Prism.languages.css["atrule"].inside.rest = Prism.languages.css;
            var markup = Prism.languages.markup;
            if (markup) {
                markup.tag.addInlined("style", "css");
                markup.tag.addAttribute("style", "css");
            }
        })(Prism);
        Prism.languages.clike = {
            comment: [ {
                pattern: /(^|[^\\])\/\*[\s\S]*?(?:\*\/|$)/,
                lookbehind: true,
                greedy: true
            }, {
                pattern: /(^|[^\\:])\/\/.*/,
                lookbehind: true,
                greedy: true
            } ],
            string: {
                pattern: /(["'])(?:\\(?:\r\n|[\s\S])|(?!\1)[^\\\r\n])*\1/,
                greedy: true
            },
            "class-name": {
                pattern: /(\b(?:class|extends|implements|instanceof|interface|new|trait)\s+|\bcatch\s+\()[\w.\\]+/i,
                lookbehind: true,
                inside: {
                    punctuation: /[.\\]/
                }
            },
            keyword: /\b(?:break|catch|continue|do|else|finally|for|function|if|in|instanceof|new|null|return|throw|try|while)\b/,
            boolean: /\b(?:false|true)\b/,
            function: /\b\w+(?=\()/,
            number: /\b0x[\da-f]+\b|(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:e[+-]?\d+)?/i,
            operator: /[<>]=?|[!=]=?=?|--?|\+\+?|&&?|\|\|?|[?*/~^%]/,
            punctuation: /[{}[\];(),.:]/
        };
        Prism.languages.javascript = Prism.languages.extend("clike", {
            "class-name": [ Prism.languages.clike["class-name"], {
                pattern: /(^|[^$\w\xA0-\uFFFF])(?!\s)[_$A-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\.(?:constructor|prototype))/,
                lookbehind: true
            } ],
            keyword: [ {
                pattern: /((?:^|\})\s*)catch\b/,
                lookbehind: true
            }, {
                pattern: /(^|[^.]|\.\.\.\s*)\b(?:as|assert(?=\s*\{)|async(?=\s*(?:function\b|\(|[$\w\xA0-\uFFFF]|$))|await|break|case|class|const|continue|debugger|default|delete|do|else|enum|export|extends|finally(?=\s*(?:\{|$))|for|from(?=\s*(?:['"]|$))|function|(?:get|set)(?=\s*(?:[#\[$\w\xA0-\uFFFF]|$))|if|implements|import|in|instanceof|interface|let|new|null|of|package|private|protected|public|return|static|super|switch|this|throw|try|typeof|undefined|var|void|while|with|yield)\b/,
                lookbehind: true
            } ],
            function: /#?(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*(?:\.\s*(?:apply|bind|call)\s*)?\()/,
            number: {
                pattern: RegExp(/(^|[^\w$])/.source + "(?:" + (/NaN|Infinity/.source + "|" + /0[bB][01]+(?:_[01]+)*n?/.source + "|" + /0[oO][0-7]+(?:_[0-7]+)*n?/.source + "|" + /0[xX][\dA-Fa-f]+(?:_[\dA-Fa-f]+)*n?/.source + "|" + /\d+(?:_\d+)*n/.source + "|" + /(?:\d+(?:_\d+)*(?:\.(?:\d+(?:_\d+)*)?)?|\.\d+(?:_\d+)*)(?:[Ee][+-]?\d+(?:_\d+)*)?/.source) + ")" + /(?![\w$])/.source),
                lookbehind: true
            },
            operator: /--|\+\+|\*\*=?|=>|&&=?|\|\|=?|[!=]==|<<=?|>>>?=?|[-+*/%&|^!=<>]=?|\.{3}|\?\?=?|\?\.?|[~:]/
        });
        Prism.languages.javascript["class-name"][0].pattern = /(\b(?:class|extends|implements|instanceof|interface|new)\s+)[\w.\\]+/;
        Prism.languages.insertBefore("javascript", "keyword", {
            regex: {
                pattern: RegExp(/((?:^|[^$\w\xA0-\uFFFF."'\])\s]|\b(?:return|yield))\s*)/.source + /\//.source + "(?:" + /(?:\[(?:[^\]\\\r\n]|\\.)*\]|\\.|[^/\\\[\r\n])+\/[dgimyus]{0,7}/.source + "|" + /(?:\[(?:[^[\]\\\r\n]|\\.|\[(?:[^[\]\\\r\n]|\\.|\[(?:[^[\]\\\r\n]|\\.)*\])*\])*\]|\\.|[^/\\\[\r\n])+\/[dgimyus]{0,7}v[dgimyus]{0,7}/.source + ")" + /(?=(?:\s|\/\*(?:[^*]|\*(?!\/))*\*\/)*(?:$|[\r\n,.;:})\]]|\/\/))/.source),
                lookbehind: true,
                greedy: true,
                inside: {
                    "regex-source": {
                        pattern: /^(\/)[\s\S]+(?=\/[a-z]*$)/,
                        lookbehind: true,
                        alias: "language-regex",
                        inside: Prism.languages.regex
                    },
                    "regex-delimiter": /^\/|\/$/,
                    "regex-flags": /^[a-z]+$/
                }
            },
            "function-variable": {
                pattern: /#?(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*[=:]\s*(?:async\s*)?(?:\bfunction\b|(?:\((?:[^()]|\([^()]*\))*\)|(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*)\s*=>))/,
                alias: "function"
            },
            parameter: [ {
                pattern: /(function(?:\s+(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*)?\s*\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\))/,
                lookbehind: true,
                inside: Prism.languages.javascript
            }, {
                pattern: /(^|[^$\w\xA0-\uFFFF])(?!\s)[_$a-z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*=>)/i,
                lookbehind: true,
                inside: Prism.languages.javascript
            }, {
                pattern: /(\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\)\s*=>)/,
                lookbehind: true,
                inside: Prism.languages.javascript
            }, {
                pattern: /((?:\b|\s|^)(?!(?:as|async|await|break|case|catch|class|const|continue|debugger|default|delete|do|else|enum|export|extends|finally|for|from|function|get|if|implements|import|in|instanceof|interface|let|new|null|of|package|private|protected|public|return|set|static|super|switch|this|throw|try|typeof|undefined|var|void|while|with|yield)(?![$\w\xA0-\uFFFF]))(?:(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*\s*)\(\s*|\]\s*\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\)\s*\{)/,
                lookbehind: true,
                inside: Prism.languages.javascript
            } ],
            constant: /\b[A-Z](?:[A-Z_]|\dx?)*\b/
        });
        Prism.languages.insertBefore("javascript", "string", {
            hashbang: {
                pattern: /^#!.*/,
                greedy: true,
                alias: "comment"
            },
            "template-string": {
                pattern: /`(?:\\[\s\S]|\$\{(?:[^{}]|\{(?:[^{}]|\{[^}]*\})*\})+\}|(?!\$\{)[^\\`])*`/,
                greedy: true,
                inside: {
                    "template-punctuation": {
                        pattern: /^`|`$/,
                        alias: "string"
                    },
                    interpolation: {
                        pattern: /((?:^|[^\\])(?:\\{2})*)\$\{(?:[^{}]|\{(?:[^{}]|\{[^}]*\})*\})+\}/,
                        lookbehind: true,
                        inside: {
                            "interpolation-punctuation": {
                                pattern: /^\$\{|\}$/,
                                alias: "punctuation"
                            },
                            rest: Prism.languages.javascript
                        }
                    },
                    string: /[\s\S]+/
                }
            },
            "string-property": {
                pattern: /((?:^|[,{])[ \t]*)(["'])(?:\\(?:\r\n|[\s\S])|(?!\2)[^\\\r\n])*\2(?=\s*:)/m,
                lookbehind: true,
                greedy: true,
                alias: "property"
            }
        });
        Prism.languages.insertBefore("javascript", "operator", {
            "literal-property": {
                pattern: /((?:^|[,{])[ \t]*)(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*:)/m,
                lookbehind: true,
                alias: "property"
            }
        });
        if (Prism.languages.markup) {
            Prism.languages.markup.tag.addInlined("script", "javascript");
            Prism.languages.markup.tag.addAttribute(/on(?:abort|blur|change|click|composition(?:end|start|update)|dblclick|error|focus(?:in|out)?|key(?:down|up)|load|mouse(?:down|enter|leave|move|out|over|up)|reset|resize|scroll|select|slotchange|submit|unload|wheel)/.source, "javascript");
        }
        Prism.languages.js = Prism.languages.javascript;
        (function() {
            if (typeof Prism === "undefined" || typeof document === "undefined") {
                return;
            }
            if (!Element.prototype.matches) {
                Element.prototype.matches = Element.prototype.msMatchesSelector || Element.prototype.webkitMatchesSelector;
            }
            var LOADING_MESSAGE = "Loading…";
            var FAILURE_MESSAGE = function(status, message) {
                return "✖ Error " + status + " while fetching file: " + message;
            };
            var FAILURE_EMPTY_MESSAGE = "✖ Error: File does not exist or is empty";
            var EXTENSIONS = {
                js: "javascript",
                py: "python",
                rb: "ruby",
                ps1: "powershell",
                psm1: "powershell",
                sh: "bash",
                bat: "batch",
                h: "c",
                tex: "latex"
            };
            var STATUS_ATTR = "data-src-status";
            var STATUS_LOADING = "loading";
            var STATUS_LOADED = "loaded";
            var STATUS_FAILED = "failed";
            var SELECTOR = "pre[data-src]:not([" + STATUS_ATTR + '="' + STATUS_LOADED + '"])' + ":not([" + STATUS_ATTR + '="' + STATUS_LOADING + '"])';
            function loadFile(src, success, error) {
                var xhr = new XMLHttpRequest;
                xhr.open("GET", src, true);
                xhr.onreadystatechange = function() {
                    if (xhr.readyState == 4) {
                        if (xhr.status < 400 && xhr.responseText) {
                            success(xhr.responseText);
                        } else {
                            if (xhr.status >= 400) {
                                error(FAILURE_MESSAGE(xhr.status, xhr.statusText));
                            } else {
                                error(FAILURE_EMPTY_MESSAGE);
                            }
                        }
                    }
                };
                xhr.send(null);
            }
            function parseRange(range) {
                var m = /^\s*(\d+)\s*(?:(,)\s*(?:(\d+)\s*)?)?$/.exec(range || "");
                if (m) {
                    var start = Number(m[1]);
                    var comma = m[2];
                    var end = m[3];
                    if (!comma) {
                        return [ start, start ];
                    }
                    if (!end) {
                        return [ start, undefined ];
                    }
                    return [ start, Number(end) ];
                }
                return undefined;
            }
            Prism.hooks.add("before-highlightall", (function(env) {
                env.selector += ", " + SELECTOR;
            }));
            Prism.hooks.add("before-sanity-check", (function(env) {
                var pre = env.element;
                if (pre.matches(SELECTOR)) {
                    env.code = "";
                    pre.setAttribute(STATUS_ATTR, STATUS_LOADING);
                    var code = pre.appendChild(document.createElement("CODE"));
                    code.textContent = LOADING_MESSAGE;
                    var src = pre.getAttribute("data-src");
                    var language = env.language;
                    if (language === "none") {
                        var extension = (/\.(\w+)$/.exec(src) || [ , "none" ])[1];
                        language = EXTENSIONS[extension] || extension;
                    }
                    Prism.util.setLanguage(code, language);
                    Prism.util.setLanguage(pre, language);
                    var autoloader = Prism.plugins.autoloader;
                    if (autoloader) {
                        autoloader.loadLanguages(language);
                    }
                    loadFile(src, (function(text) {
                        pre.setAttribute(STATUS_ATTR, STATUS_LOADED);
                        var range = parseRange(pre.getAttribute("data-range"));
                        if (range) {
                            var lines = text.split(/\r\n?|\n/g);
                            var start = range[0];
                            var end = range[1] == null ? lines.length : range[1];
                            if (start < 0) {
                                start += lines.length;
                            }
                            start = Math.max(0, Math.min(start - 1, lines.length));
                            if (end < 0) {
                                end += lines.length;
                            }
                            end = Math.max(0, Math.min(end, lines.length));
                            text = lines.slice(start, end).join("\n");
                            if (!pre.hasAttribute("data-start")) {
                                pre.setAttribute("data-start", String(start + 1));
                            }
                        }
                        code.textContent = text;
                        Prism.highlightElement(code);
                    }), (function(error) {
                        pre.setAttribute(STATUS_ATTR, STATUS_FAILED);
                        code.textContent = error;
                    }));
                }
            }));
            Prism.plugins.fileHighlight = {
                highlight: function highlight(container) {
                    var elements = (container || document).querySelectorAll(SELECTOR);
                    for (var i = 0, element; element = elements[i++]; ) {
                        Prism.highlightElement(element);
                    }
                }
            };
            var logged = false;
            Prism.fileHighlight = function() {
                if (!logged) {
                    console.warn("Prism.fileHighlight is deprecated. Use `Prism.plugins.fileHighlight.highlight` instead.");
                    logged = true;
                }
                Prism.plugins.fileHighlight.highlight.apply(this, arguments);
            };
        })();
    })(prism$1);
    var prismExports = prism$1.exports;
    var prism = getDefaultExportFromCjs(prismExports);
    (function(Prism) {
        function getPlaceholder(language, index) {
            return "___" + language.toUpperCase() + index + "___";
        }
        Object.defineProperties(Prism.languages["markup-templating"] = {}, {
            buildPlaceholders: {
                value: function(env, language, placeholderPattern, replaceFilter) {
                    if (env.language !== language) {
                        return;
                    }
                    var tokenStack = env.tokenStack = [];
                    env.code = env.code.replace(placeholderPattern, (function(match) {
                        if (typeof replaceFilter === "function" && !replaceFilter(match)) {
                            return match;
                        }
                        var i = tokenStack.length;
                        var placeholder;
                        while (env.code.indexOf(placeholder = getPlaceholder(language, i)) !== -1) {
                            ++i;
                        }
                        tokenStack[i] = match;
                        return placeholder;
                    }));
                    env.grammar = Prism.languages.markup;
                }
            },
            tokenizePlaceholders: {
                value: function(env, language) {
                    if (env.language !== language || !env.tokenStack) {
                        return;
                    }
                    env.grammar = Prism.languages[language];
                    var j = 0;
                    var keys = Object.keys(env.tokenStack);
                    function walkTokens(tokens) {
                        for (var i = 0; i < tokens.length; i++) {
                            if (j >= keys.length) {
                                break;
                            }
                            var token = tokens[i];
                            if (typeof token === "string" || token.content && typeof token.content === "string") {
                                var k = keys[j];
                                var t = env.tokenStack[k];
                                var s = typeof token === "string" ? token : token.content;
                                var placeholder = getPlaceholder(language, k);
                                var index = s.indexOf(placeholder);
                                if (index > -1) {
                                    ++j;
                                    var before = s.substring(0, index);
                                    var middle = new Prism.Token(language, Prism.tokenize(t, env.grammar), "language-" + language, t);
                                    var after = s.substring(index + placeholder.length);
                                    var replacement = [];
                                    if (before) {
                                        replacement.push.apply(replacement, walkTokens([ before ]));
                                    }
                                    replacement.push(middle);
                                    if (after) {
                                        replacement.push.apply(replacement, walkTokens([ after ]));
                                    }
                                    if (typeof token === "string") {
                                        tokens.splice.apply(tokens, [ i, 1 ].concat(replacement));
                                    } else {
                                        token.content = replacement;
                                    }
                                }
                            } else if (token.content) {
                                walkTokens(token.content);
                            }
                        }
                        return tokens;
                    }
                    walkTokens(env.tokens);
                }
            }
        });
    })(Prism);
    const highlightCodeCompiler = _ref => {
        let {renderer: renderer} = _ref;
        return renderer.code = function(code) {
            let lang = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : "markup";
            const langOrMarkup = prism.languages[lang] || prism.languages.markup;
            const text = prism.highlight(code.replace(/@DOCSIFY_QM@/g, "`"), langOrMarkup, lang);
            return `<pre data-lang="${lang}"><code class="lang-${lang}" tabindex="0">${text}</code></pre>`;
        };
    };
    const paragraphCompiler = _ref => {
        let {renderer: renderer} = _ref;
        return renderer.paragraph = text => {
            let result;
            if (/^!&gt;/.test(text)) {
                result = helper("tip", text);
            } else if (/^\?&gt;/.test(text)) {
                result = helper("warn", text);
            } else {
                result = `<p>${text}</p>`;
            }
            return result;
        };
    };
    const taskListCompiler = _ref => {
        let {renderer: renderer} = _ref;
        return renderer.list = (body, ordered, start) => {
            const isTaskList = /<li class="task-list-item">/.test(body.split('class="task-list"')[0]);
            const isStartReq = start && start > 1;
            const tag = ordered ? "ol" : "ul";
            const tagAttrs = [ isTaskList ? 'class="task-list"' : "", isStartReq ? `start="${start}"` : "" ].join(" ").trim();
            return `<${tag} ${tagAttrs}>${body}</${tag}>`;
        };
    };
    const taskListItemCompiler = _ref => {
        let {renderer: renderer} = _ref;
        return renderer.listitem = text => {
            const isTaskItem = /^(<input.*type="checkbox"[^>]*>)/.test(text);
            const html = isTaskItem ? `<li class="task-list-item"><label>${text}</label></li>` : `<li>${text}</li>`;
            return html;
        };
    };
    const linkCompiler = _ref => {
        let {renderer: renderer, router: router, linkTarget: linkTarget, linkRel: linkRel, compilerClass: compilerClass} = _ref;
        return renderer.link = function(href) {
            let title = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : "";
            let text = arguments.length > 2 ? arguments[2] : undefined;
            const attrs = [];
            const {str: str, config: config} = getAndRemoveConfig(title);
            linkTarget = config.target || linkTarget;
            linkRel = linkTarget === "_blank" ? compilerClass.config.externalLinkRel || "noopener" : "";
            title = str;
            if (!isAbsolutePath(href) && !compilerClass._matchNotCompileLink(href) && !config.ignore) {
                if (href === compilerClass.config.homepage) {
                    href = "README";
                }
                href = router.toURL(href, null, router.getCurrentPath());
                if (config.target) {
                    href.indexOf("mailto:") !== 0 && attrs.push(`target="${linkTarget}"`);
                }
            } else {
                if (!isAbsolutePath(href) && href.slice(0, 2) === "./") {
                    href = document.URL.replace(/\/(?!.*\/).*/, "/").replace("#/./", "") + href;
                }
                attrs.push(href.indexOf("mailto:") === 0 ? "" : `target="${linkTarget}"`);
                attrs.push(href.indexOf("mailto:") === 0 ? "" : linkRel !== "" ? ` rel="${linkRel}"` : "");
            }
            if (config.disabled) {
                attrs.push("disabled");
                href = "javascript:void(0)";
            }
            if (config.class) {
                attrs.push(`class="${config.class}"`);
            }
            if (config.id) {
                attrs.push(`id="${config.id}"`);
            }
            if (title) {
                attrs.push(`title="${title}"`);
            }
            return `<a href="${href}" ${attrs.join(" ")}>${text}</a>`;
        };
    };
    const cachedLinks = {};
    const compileMedia = {
        markdown(url) {
            return {
                url: url
            };
        },
        mermaid(url) {
            return {
                url: url
            };
        },
        iframe(url, title) {
            return {
                html: `<iframe src="${url}" ${title || "width=100% height=400"}></iframe>`
            };
        },
        video(url, title) {
            return {
                html: `<video src="${url}" ${title || "controls"}>Not Support</video>`
            };
        },
        audio(url, title) {
            return {
                html: `<audio src="${url}" ${title || "controls"}>Not Support</audio>`
            };
        },
        code(url, title) {
            let lang = url.match(/\.(\w+)$/);
            lang = title || lang && lang[1];
            if (lang === "md") {
                lang = "markdown";
            }
            return {
                url: url,
                lang: lang
            };
        }
    };
    class Compiler {
        constructor(config, router) {
            this.config = config;
            this.router = router;
            this.cacheTree = {};
            this.toc = [];
            this.cacheTOC = {};
            this.linkTarget = config.externalLinkTarget || "_blank";
            this.linkRel = this.linkTarget === "_blank" ? config.externalLinkRel || "noopener" : "";
            this.contentBase = router.getBasePath();
            const renderer = this._initRenderer();
            this.heading = renderer.heading;
            let compile;
            const mdConf = config.markdown || {};
            if (isFn(mdConf)) {
                compile = mdConf(marked, renderer);
            } else {
                marked.setOptions(Object.assign(mdConf, {
                    renderer: Object.assign(renderer, mdConf.renderer)
                }));
                compile = marked;
            }
            this._marked = compile;
            this.compile = text => {
                let isCached = true;
                const result = cached$1((_ => {
                    isCached = false;
                    let html = "";
                    if (!text) {
                        return text;
                    }
                    if (isPrimitive(text)) {
                        html = compile(text);
                    } else {
                        html = compile.parser(text);
                    }
                    html = config.noEmoji ? html : emojify(html, config.nativeEmoji);
                    slugify.clear();
                    return html;
                }))(text);
                const curFileName = this.router.parse().file;
                if (isCached) {
                    this.toc = this.cacheTOC[curFileName];
                } else {
                    this.cacheTOC[curFileName] = [ ...this.toc ];
                }
                return result;
            };
        }
        compileEmbed(href, title) {
            const {str: str, config: config} = getAndRemoveConfig(title);
            let embed;
            title = str;
            if (config.include) {
                if (!isAbsolutePath(href)) {
                    href = getPath(this.contentBase, getParentPath(this.router.getCurrentPath()), href);
                }
                let media;
                if (config.type && (media = compileMedia[config.type])) {
                    embed = media.call(this, href, title);
                    embed.type = config.type;
                } else {
                    let type = "code";
                    if (/\.(md|markdown)/.test(href)) {
                        type = "markdown";
                    } else if (/\.mmd/.test(href)) {
                        type = "mermaid";
                    } else if (/\.html?/.test(href)) {
                        type = "iframe";
                    } else if (/\.(mp4|ogg)/.test(href)) {
                        type = "video";
                    } else if (/\.mp3/.test(href)) {
                        type = "audio";
                    }
                    embed = compileMedia[type].call(this, href, title);
                    embed.type = type;
                }
                embed.fragment = config.fragment;
                return embed;
            }
        }
        _matchNotCompileLink(link) {
            const links = this.config.noCompileLinks || [];
            for (const n of links) {
                const re = cachedLinks[n] || (cachedLinks[n] = new RegExp(`^${n}$`));
                if (re.test(link)) {
                    return link;
                }
            }
        }
        _initRenderer() {
            const renderer = new marked.Renderer;
            const {linkTarget: linkTarget, linkRel: linkRel, router: router, contentBase: contentBase} = this;
            const _self = this;
            const origin = {};
            origin.heading = renderer.heading = function(text, level) {
                let {str: str, config: config} = getAndRemoveConfig(text);
                const nextToc = {
                    level: level,
                    title: str
                };
                const {content: content, ignoreAllSubs: ignoreAllSubs, ignoreSubHeading: ignoreSubHeading} = getAndRemoveDocisfyIgnoreConfig(str);
                str = content.trim();
                nextToc.title = removeAtag(str);
                nextToc.ignoreAllSubs = ignoreAllSubs;
                nextToc.ignoreSubHeading = ignoreSubHeading;
                const slug = slugify(config.id || str);
                const url = router.toURL(router.getCurrentPath(), {
                    id: slug
                });
                nextToc.slug = url;
                _self.toc.push(nextToc);
                return `<h${level} id="${slug}" tabindex="-1"><a href="${url}" data-id="${slug}" class="anchor"><span>${str}</span></a></h${level}>`;
            };
            origin.code = highlightCodeCompiler({
                renderer: renderer
            });
            origin.link = linkCompiler({
                renderer: renderer,
                router: router,
                linkTarget: linkTarget,
                linkRel: linkRel,
                compilerClass: _self
            });
            origin.paragraph = paragraphCompiler({
                renderer: renderer
            });
            origin.image = imageCompiler({
                renderer: renderer,
                contentBase: contentBase,
                router: router
            });
            origin.list = taskListCompiler({
                renderer: renderer
            });
            origin.listitem = taskListItemCompiler({
                renderer: renderer
            });
            renderer.origin = origin;
            return renderer;
        }
        sidebar(text, level) {
            const {toc: toc} = this;
            const currentPath = this.router.getCurrentPath();
            let html = "";
            if (text) {
                html = this.compile(text);
            } else {
                for (let i = 0; i < toc.length; i++) {
                    if (toc[i].ignoreSubHeading) {
                        const deletedHeaderLevel = toc[i].level;
                        toc.splice(i, 1);
                        for (let j = i; j < toc.length && deletedHeaderLevel < toc[j].level; j++) {
                            toc.splice(j, 1) && j-- && i++;
                        }
                        i--;
                    }
                }
                const tree$1 = this.cacheTree[currentPath] || genTree(toc, level);
                html = tree(tree$1, "<ul>{inner}</ul>");
                this.cacheTree[currentPath] = tree$1;
            }
            return html;
        }
        subSidebar(level) {
            if (!level) {
                this.toc = [];
                return;
            }
            const currentPath = this.router.getCurrentPath();
            const {cacheTree: cacheTree, toc: toc} = this;
            toc[0] && toc[0].ignoreAllSubs && toc.splice(0);
            toc[0] && toc[0].level === 1 && toc.shift();
            for (let i = 0; i < toc.length; i++) {
                toc[i].ignoreSubHeading && toc.splice(i, 1) && i--;
            }
            const tree$1 = cacheTree[currentPath] || genTree(toc, level);
            cacheTree[currentPath] = tree$1;
            this.toc = [];
            return tree(tree$1);
        }
        header(text, level) {
            return this.heading(text, level);
        }
        article(text) {
            return this.compile(text);
        }
        cover(text) {
            const cacheToc = this.toc.slice();
            const html = this.compile(text);
            this.toc = cacheToc.slice();
            return html;
        }
    }
    var _createClass = function() {
        function defineProperties(target, props) {
            for (var i = 0; i < props.length; i++) {
                var descriptor = props[i];
                descriptor.enumerable = descriptor.enumerable || false;
                descriptor.configurable = true;
                if ("value" in descriptor) descriptor.writable = true;
                Object.defineProperty(target, descriptor.key, descriptor);
            }
        }
        return function(Constructor, protoProps, staticProps) {
            if (protoProps) defineProperties(Constructor.prototype, protoProps);
            if (staticProps) defineProperties(Constructor, staticProps);
            return Constructor;
        };
    }();
    var _templateObject = _taggedTemplateLiteral([ "", "" ], [ "", "" ]);
    function _taggedTemplateLiteral(strings, raw) {
        return Object.freeze(Object.defineProperties(strings, {
            raw: {
                value: Object.freeze(raw)
            }
        }));
    }
    function _classCallCheck(instance, Constructor) {
        if (!(instance instanceof Constructor)) {
            throw new TypeError("Cannot call a class as a function");
        }
    }
    var TemplateTag = function() {
        function TemplateTag() {
            var _this = this;
            for (var _len = arguments.length, transformers = Array(_len), _key = 0; _key < _len; _key++) {
                transformers[_key] = arguments[_key];
            }
            _classCallCheck(this, TemplateTag);
            this.tag = function(strings) {
                for (var _len2 = arguments.length, expressions = Array(_len2 > 1 ? _len2 - 1 : 0), _key2 = 1; _key2 < _len2; _key2++) {
                    expressions[_key2 - 1] = arguments[_key2];
                }
                if (typeof strings === "function") {
                    return _this.interimTag.bind(_this, strings);
                }
                if (typeof strings === "string") {
                    return _this.transformEndResult(strings);
                }
                strings = strings.map(_this.transformString.bind(_this));
                return _this.transformEndResult(strings.reduce(_this.processSubstitutions.bind(_this, expressions)));
            };
            if (transformers.length > 0 && Array.isArray(transformers[0])) {
                transformers = transformers[0];
            }
            this.transformers = transformers.map((function(transformer) {
                return typeof transformer === "function" ? transformer() : transformer;
            }));
            return this.tag;
        }
        _createClass(TemplateTag, [ {
            key: "interimTag",
            value: function interimTag(previousTag, template) {
                for (var _len3 = arguments.length, substitutions = Array(_len3 > 2 ? _len3 - 2 : 0), _key3 = 2; _key3 < _len3; _key3++) {
                    substitutions[_key3 - 2] = arguments[_key3];
                }
                return this.tag(_templateObject, previousTag.apply(undefined, [ template ].concat(substitutions)));
            }
        }, {
            key: "processSubstitutions",
            value: function processSubstitutions(substitutions, resultSoFar, remainingPart) {
                var substitution = this.transformSubstitution(substitutions.shift(), resultSoFar);
                return "".concat(resultSoFar, substitution, remainingPart);
            }
        }, {
            key: "transformString",
            value: function transformString(str) {
                var cb = function cb(res, transform) {
                    return transform.onString ? transform.onString(res) : res;
                };
                return this.transformers.reduce(cb, str);
            }
        }, {
            key: "transformSubstitution",
            value: function transformSubstitution(substitution, resultSoFar) {
                var cb = function cb(res, transform) {
                    return transform.onSubstitution ? transform.onSubstitution(res, resultSoFar) : res;
                };
                return this.transformers.reduce(cb, substitution);
            }
        }, {
            key: "transformEndResult",
            value: function transformEndResult(endResult) {
                var cb = function cb(res, transform) {
                    return transform.onEndResult ? transform.onEndResult(res) : res;
                };
                return this.transformers.reduce(cb, endResult);
            }
        } ]);
        return TemplateTag;
    }();
    var trimResultTransformer = function trimResultTransformer() {
        var side = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : "";
        return {
            onEndResult: function onEndResult(endResult) {
                if (side === "") {
                    return endResult.trim();
                }
                side = side.toLowerCase();
                if (side === "start" || side === "left") {
                    return endResult.replace(/^\s*/, "");
                }
                if (side === "end" || side === "right") {
                    return endResult.replace(/\s*$/, "");
                }
                throw new Error("Side not supported: " + side);
            }
        };
    };
    function _toConsumableArray(arr) {
        if (Array.isArray(arr)) {
            for (var i = 0, arr2 = Array(arr.length); i < arr.length; i++) {
                arr2[i] = arr[i];
            }
            return arr2;
        } else {
            return Array.from(arr);
        }
    }
    var stripIndentTransformer = function stripIndentTransformer() {
        var type = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : "initial";
        return {
            onEndResult: function onEndResult(endResult) {
                if (type === "initial") {
                    var match = endResult.match(/^[^\S\n]*(?=\S)/gm);
                    var indent = match && Math.min.apply(Math, _toConsumableArray(match.map((function(el) {
                        return el.length;
                    }))));
                    if (indent) {
                        var regexp = new RegExp("^.{" + indent + "}", "gm");
                        return endResult.replace(regexp, "");
                    }
                    return endResult;
                }
                if (type === "all") {
                    return endResult.replace(/^[^\S\n]+/gm, "");
                }
                throw new Error("Unknown type: " + type);
            }
        };
    };
    var replaceResultTransformer = function replaceResultTransformer(replaceWhat, replaceWith) {
        return {
            onEndResult: function onEndResult(endResult) {
                if (replaceWhat == null || replaceWith == null) {
                    throw new Error("replaceResultTransformer requires at least 2 arguments.");
                }
                return endResult.replace(replaceWhat, replaceWith);
            }
        };
    };
    var replaceSubstitutionTransformer = function replaceSubstitutionTransformer(replaceWhat, replaceWith) {
        return {
            onSubstitution: function onSubstitution(substitution, resultSoFar) {
                if (replaceWhat == null || replaceWith == null) {
                    throw new Error("replaceSubstitutionTransformer requires at least 2 arguments.");
                }
                if (substitution == null) {
                    return substitution;
                } else {
                    return substitution.toString().replace(replaceWhat, replaceWith);
                }
            }
        };
    };
    var defaults = {
        separator: "",
        conjunction: "",
        serial: false
    };
    var inlineArrayTransformer = function inlineArrayTransformer() {
        var opts = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : defaults;
        return {
            onSubstitution: function onSubstitution(substitution, resultSoFar) {
                if (Array.isArray(substitution)) {
                    var arrayLength = substitution.length;
                    var separator = opts.separator;
                    var conjunction = opts.conjunction;
                    var serial = opts.serial;
                    var indent = resultSoFar.match(/(\n?[^\S\n]+)$/);
                    if (indent) {
                        substitution = substitution.join(separator + indent[1]);
                    } else {
                        substitution = substitution.join(separator + " ");
                    }
                    if (conjunction && arrayLength > 1) {
                        var separatorIndex = substitution.lastIndexOf(separator);
                        substitution = substitution.slice(0, separatorIndex) + (serial ? separator : "") + " " + conjunction + substitution.slice(separatorIndex + 1);
                    }
                }
                return substitution;
            }
        };
    };
    var splitStringTransformer = function splitStringTransformer(splitBy) {
        return {
            onSubstitution: function onSubstitution(substitution, resultSoFar) {
                {
                    if (typeof substitution === "string" && substitution.includes(splitBy)) {
                        substitution = substitution.split(splitBy);
                    }
                }
                return substitution;
            }
        };
    };
    var isValidValue = function isValidValue(x) {
        return x != null && !Number.isNaN(x) && typeof x !== "boolean";
    };
    var removeNonPrintingValuesTransformer = function removeNonPrintingValuesTransformer() {
        return {
            onSubstitution: function onSubstitution(substitution) {
                if (Array.isArray(substitution)) {
                    return substitution.filter(isValidValue);
                }
                if (isValidValue(substitution)) {
                    return substitution;
                }
                return "";
            }
        };
    };
    new TemplateTag(inlineArrayTransformer({
        separator: ","
    }), stripIndentTransformer, trimResultTransformer);
    new TemplateTag(inlineArrayTransformer({
        separator: ",",
        conjunction: "and"
    }), stripIndentTransformer, trimResultTransformer);
    new TemplateTag(inlineArrayTransformer({
        separator: ",",
        conjunction: "or"
    }), stripIndentTransformer, trimResultTransformer);
    new TemplateTag(splitStringTransformer("\n"), removeNonPrintingValuesTransformer, inlineArrayTransformer, stripIndentTransformer, trimResultTransformer);
    new TemplateTag(splitStringTransformer("\n"), inlineArrayTransformer, stripIndentTransformer, trimResultTransformer, replaceSubstitutionTransformer(/&/g, "&amp;"), replaceSubstitutionTransformer(/</g, "&lt;"), replaceSubstitutionTransformer(/>/g, "&gt;"), replaceSubstitutionTransformer(/"/g, "&quot;"), replaceSubstitutionTransformer(/'/g, "&#x27;"), replaceSubstitutionTransformer(/`/g, "&#x60;"));
    new TemplateTag(replaceResultTransformer(/(?:\n(?:\s*))+/g, " "), trimResultTransformer);
    new TemplateTag(replaceResultTransformer(/(?:\n\s*)/g, ""), trimResultTransformer);
    new TemplateTag(inlineArrayTransformer({
        separator: ","
    }), replaceResultTransformer(/(?:\s+)/g, " "), trimResultTransformer);
    new TemplateTag(inlineArrayTransformer({
        separator: ",",
        conjunction: "or"
    }), replaceResultTransformer(/(?:\s+)/g, " "), trimResultTransformer);
    new TemplateTag(inlineArrayTransformer({
        separator: ",",
        conjunction: "and"
    }), replaceResultTransformer(/(?:\s+)/g, " "), trimResultTransformer);
    new TemplateTag(inlineArrayTransformer, stripIndentTransformer, trimResultTransformer);
    new TemplateTag(inlineArrayTransformer, replaceResultTransformer(/(?:\s+)/g, " "), trimResultTransformer);
    var stripIndent = new TemplateTag(stripIndentTransformer, trimResultTransformer);
    new TemplateTag(stripIndentTransformer("all"), trimResultTransformer);
    let barEl;
    let timeId;
    function init() {
        const div = create("div");
        div.classList.add("progress");
        div.setAttribute("role", "progressbar");
        div.setAttribute("aria-valuemin", "0");
        div.setAttribute("aria-valuemax", "100");
        div.setAttribute("aria-label", "Loading...");
        appendTo(body, div);
        barEl = div;
    }
    function progressbar(info) {
        const {loaded: loaded, total: total, step: step} = info;
        let num;
        !barEl && init();
        if (typeof step !== "undefined") {
            num = parseInt(barEl.style.width || 0, 10) + step;
            num = num > 80 ? 80 : num;
        } else {
            num = Math.floor(loaded / total * 100);
        }
        barEl.style.opacity = 1;
        barEl.style.width = num >= 95 ? "100%" : num + "%";
        barEl.setAttribute("aria-valuenow", num >= 95 ? 100 : num);
        if (num >= 95) {
            clearTimeout(timeId);
            timeId = setTimeout((_ => {
                barEl.style.opacity = 0;
                barEl.style.width = "0%";
                barEl.removeAttribute("aria-valuenow");
            }), 200);
        }
    }
    const cache = {};
    function get(url) {
        let hasBar = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : false;
        let headers = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : {};
        const xhr = new XMLHttpRequest;
        const cached = cache[url];
        if (cached) {
            return {
                then: cb => cb(cached.content, cached.opt),
                abort: noop
            };
        }
        xhr.open("GET", url);
        for (const i of Object.keys(headers)) {
            xhr.setRequestHeader(i, headers[i]);
        }
        xhr.send();
        return {
            then(success) {
                let error = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : noop;
                const getResponseStatus = event => ({
                    ok: event.target.status >= 200 && event.target.status < 300,
                    status: event.target.status,
                    statusText: event.target.statusText
                });
                if (hasBar) {
                    const id = setInterval((_ => progressbar({
                        step: Math.floor(Math.random() * 5 + 1)
                    })), 500);
                    xhr.addEventListener("progress", progressbar);
                    xhr.addEventListener("loadend", (evt => {
                        progressbar(evt);
                        clearInterval(id);
                    }));
                }
                xhr.addEventListener("error", (event => {
                    error(event, getResponseStatus(event));
                }));
                xhr.addEventListener("load", (event => {
                    const target = event.target;
                    if (target.status >= 400) {
                        error(event, getResponseStatus(event));
                    } else {
                        if (typeof target.response !== "string") {
                            throw new TypeError("Unsupported content type.");
                        }
                        const result = cache[url] = {
                            content: target.response,
                            opt: {
                                updatedAt: xhr.getResponseHeader("last-modified") ?? ""
                            }
                        };
                        success(result.content, result.opt, getResponseStatus(event));
                    }
                }));
            },
            abort: _ => xhr.readyState !== 4 && xhr.abort()
        };
    }
    const cached = {};
    function walkFetchEmbed(_ref, cb) {
        let {embedTokens: embedTokens, compile: compile, fetch: fetch} = _ref;
        let token;
        let step = 0;
        let count = 0;
        if (!embedTokens.length) {
            return cb({});
        }
        while (token = embedTokens[step++]) {
            const currentToken = token;
            const next = text => {
                let embedToken;
                if (text) {
                    if (currentToken.embed.type === "markdown") {
                        let path = currentToken.embed.url.split("/");
                        path.pop();
                        path = path.join("/");
                        text = text.replace(/\[([^[\]]+)\]\(([^)]+)\)/g, (x => {
                            const linkBeginIndex = x.indexOf("(");
                            if (x.slice(linkBeginIndex, linkBeginIndex + 2) === "(.") {
                                return x.substring(0, linkBeginIndex) + `(${window.location.protocol}//${window.location.host}${path}/` + x.substring(linkBeginIndex + 1, x.length - 1) + ")";
                            }
                            return x;
                        }));
                        const frontMatterInstalled = ($docsify.frontMatter || {}).installed || false;
                        if (frontMatterInstalled === true) {
                            text = $docsify.frontMatter.parseMarkdown(text);
                        }
                        embedToken = compile.lexer(text);
                    } else if (currentToken.embed.type === "code") {
                        if (currentToken.embed.fragment) {
                            const fragment = currentToken.embed.fragment;
                            const pattern = new RegExp(`(?:###|\\/\\/\\/)\\s*\\[${fragment}\\]([\\s\\S]*)(?:###|\\/\\/\\/)\\s*\\[${fragment}\\]`);
                            text = stripIndent((text.match(pattern) || [])[1] || "").trim();
                        }
                        embedToken = compile.lexer("```" + currentToken.embed.lang + "\n" + text.replace(/`/g, "@DOCSIFY_QM@") + "\n```\n");
                    } else if (currentToken.embed.type === "mermaid") {
                        embedToken = [ {
                            type: "html",
                            text: `<div class="mermaid">\n${text}\n</div>`
                        } ];
                        embedToken.links = {};
                    } else {
                        embedToken = [ {
                            type: "html",
                            text: text
                        } ];
                        embedToken.links = {};
                    }
                }
                cb({
                    token: currentToken,
                    embedToken: embedToken
                });
                if (++count >= embedTokens.length) {
                    cb({});
                }
            };
            if (token.embed.url) {
                get(token.embed.url).then(next);
            } else {
                next(token.embed.html);
            }
        }
    }
    function prerenderEmbed(_ref2, done) {
        let {compiler: compiler, raw: raw = "", fetch: fetch} = _ref2;
        const hit = cached[raw];
        if (hit) {
            const copy = hit.slice();
            copy.links = hit.links;
            return done(copy);
        }
        const compile = compiler._marked;
        let tokens = compile.lexer(raw);
        const embedTokens = [];
        const linkRE = compile.Lexer.rules.inline.normal.link;
        const links = tokens.links;
        tokens.forEach(((token, index) => {
            if (token.type === "paragraph") {
                token.text = token.text.replace(new RegExp(linkRE.source, "g"), ((src, filename, href, title) => {
                    const embed = compiler.compileEmbed(href, title);
                    if (embed) {
                        embedTokens.push({
                            index: index,
                            embed: embed
                        });
                    }
                    return src;
                }));
            }
        }));
        const moves = [];
        walkFetchEmbed({
            compile: compile,
            embedTokens: embedTokens,
            fetch: fetch
        }, (_ref3 => {
            let {embedToken: embedToken, token: token} = _ref3;
            if (token) {
                let index = token.index;
                moves.forEach((pos => {
                    if (index > pos.start) {
                        index += pos.length;
                    }
                }));
                Object.assign(links, embedToken.links);
                tokens = tokens.slice(0, index).concat(embedToken, tokens.slice(index + 1));
                moves.push({
                    start: index,
                    length: embedToken.length - 1
                });
            } else {
                cached[raw] = tokens.concat();
                tokens.links = cached[raw].links = links;
                done(tokens);
            }
        }));
    }
    function Render(Base) {
        return class Render extends Base {
            #vueGlobalData;
            #addTextAsTitleAttribute(cssSelector) {
                findAll(cssSelector).forEach((elm => {
                    if (!elm.title && elm.innerText) {
                        elm.title = elm.innerText;
                    }
                }));
            }
            #executeScript() {
                const script = findAll(".markdown-section>script").filter((s => !/template/.test(s.type)))[0];
                if (!script) {
                    return false;
                }
                const code = script.innerText.trim();
                if (!code) {
                    return false;
                }
                new Function(code)();
            }
            #formatUpdated(html, updated, fn) {
                updated = typeof fn === "function" ? fn(updated) : typeof fn === "string" ? tinydate(fn)(new Date(updated)) : updated;
                return html.replace(/{docsify-updated}/g, updated);
            }
            #renderMain(html) {
                const docsifyConfig = this.config;
                const markdownElm = find(".markdown-section");
                const vueVersion = "Vue" in window && window.Vue.version && Number(window.Vue.version.charAt(0));
                const isMountedVue = elm => {
                    const isVue2 = Boolean(elm.__vue__ && elm.__vue__._isVue);
                    const isVue3 = Boolean(elm._vnode && elm._vnode.__v_skip);
                    return isVue2 || isVue3;
                };
                if ("Vue" in window) {
                    const mountedElms = findAll(".markdown-section > *").filter((elm => isMountedVue(elm)));
                    for (const mountedElm of mountedElms) {
                        if (vueVersion === 2) {
                            mountedElm.__vue__.$destroy();
                        } else if (vueVersion === 3) {
                            mountedElm.__vue_app__.unmount();
                        }
                    }
                }
                this._renderTo(markdownElm, html);
                !docsifyConfig.loadSidebar && this._renderSidebar();
                if (docsifyConfig.executeScript || "Vue" in window && docsifyConfig.executeScript !== false) {
                    this.#executeScript();
                }
                if ("Vue" in window) {
                    const vueGlobalOptions = docsifyConfig.vueGlobalOptions || {};
                    const vueMountData = [];
                    const vueComponentNames = Object.keys(docsifyConfig.vueComponents || {});
                    if (vueVersion === 2 && vueComponentNames.length) {
                        vueComponentNames.forEach((name => {
                            const isNotRegistered = !window.Vue.options.components[name];
                            if (isNotRegistered) {
                                window.Vue.component(name, docsifyConfig.vueComponents[name]);
                            }
                        }));
                    }
                    if (!this.#vueGlobalData && vueGlobalOptions.data && typeof vueGlobalOptions.data === "function") {
                        this.#vueGlobalData = vueGlobalOptions.data();
                    }
                    vueMountData.push(...Object.keys(docsifyConfig.vueMounts || {}).map((cssSelector => [ find(markdownElm, cssSelector), docsifyConfig.vueMounts[cssSelector] ])).filter((_ref => {
                        let [elm, vueConfig] = _ref;
                        return elm;
                    })));
                    const reHasBraces = /{{2}[^{}]*}{2}/;
                    const reHasDirective = /<[^>/]+\s([@:]|v-)[\w-:.[\]]+[=>\s]/;
                    vueMountData.push(...findAll(".markdown-section > *").filter((elm => !vueMountData.some((_ref2 => {
                        let [e, c] = _ref2;
                        return e === elm;
                    })))).filter((elm => {
                        const isVueMount = elm.tagName.toLowerCase() in (docsifyConfig.vueComponents || {}) || elm.querySelector(vueComponentNames.join(",") || null) || reHasBraces.test(elm.outerHTML) || reHasDirective.test(elm.outerHTML);
                        return isVueMount;
                    })).map((elm => {
                        const vueConfig = {
                            ...vueGlobalOptions
                        };
                        if (this.#vueGlobalData) {
                            vueConfig.data = () => this.#vueGlobalData;
                        }
                        return [ elm, vueConfig ];
                    })));
                    if (vueMountData.length === 0) {
                        return;
                    }
                    for (const [mountElm, vueConfig] of vueMountData) {
                        const isVueAttr = "data-isvue";
                        const isSkipElm = mountElm.matches("pre, :not([v-template]):has(pre), script") || isMountedVue(mountElm) || mountElm.querySelector(`[${isVueAttr}]`);
                        if (!isSkipElm) {
                            mountElm.setAttribute(isVueAttr, "");
                            if (vueVersion === 2) {
                                vueConfig.el = undefined;
                                new window.Vue(vueConfig).$mount(mountElm);
                            } else if (vueVersion === 3) {
                                const app = window.Vue.createApp(vueConfig);
                                vueComponentNames.forEach((name => {
                                    const config = docsifyConfig.vueComponents[name];
                                    app.component(name, config);
                                }));
                                app.mount(mountElm);
                            }
                        }
                    }
                }
            }
            #renderNameLink(vm) {
                const el = getNode(".app-name-link");
                const nameLink = vm.config.nameLink;
                const path = vm.route.path;
                if (!el) {
                    return;
                }
                if (isPrimitive(vm.config.nameLink)) {
                    el.setAttribute("href", nameLink);
                } else if (typeof nameLink === "object") {
                    const match = Object.keys(nameLink).filter((key => path.indexOf(key) > -1))[0];
                    el.setAttribute("href", nameLink[match]);
                }
            }
            #renderSkipLink(vm) {
                const {skipLink: skipLink} = vm.config;
                if (skipLink !== false) {
                    const el = getNode("#skip-to-content");
                    let skipLinkText = typeof skipLink === "string" ? skipLink : "Skip to main content";
                    if (skipLink?.constructor === Object) {
                        const matchingPath = Object.keys(skipLink).find((path => vm.route.path.startsWith(path.startsWith("/") ? path : `/${path}`)));
                        const matchingText = matchingPath && skipLink[matchingPath];
                        skipLinkText = matchingText || skipLinkText;
                    }
                    if (el) {
                        el.innerHTML = skipLinkText;
                    } else {
                        const html = `<button id="skip-to-content">${skipLinkText}</button>`;
                        body.insertAdjacentHTML("afterbegin", html);
                    }
                }
            }
            _renderTo(el, content, replace) {
                const node = getNode(el);
                if (node) {
                    node[replace ? "outerHTML" : "innerHTML"] = content;
                }
            }
            _renderSidebar(text) {
                const {maxLevel: maxLevel, subMaxLevel: subMaxLevel, loadSidebar: loadSidebar, hideSidebar: hideSidebar} = this.config;
                const sidebarEl = getNode("aside.sidebar");
                const sidebarToggleEl = getNode("button.sidebar-toggle");
                if (hideSidebar) {
                    body.classList.add("hidesidebar");
                    sidebarEl?.remove(sidebarEl);
                    sidebarToggleEl?.remove(sidebarToggleEl);
                    return null;
                }
                this._renderTo(".sidebar-nav", this.compiler.sidebar(text, maxLevel));
                sidebarToggleEl.setAttribute("aria-expanded", !isMobile);
                const activeElmHref = this.router.toURL(this.route.path);
                const activeEl = find(`.sidebar-nav a[href="${activeElmHref}"]`);
                this.#addTextAsTitleAttribute(".sidebar-nav a");
                if (loadSidebar && activeEl) {
                    activeEl.parentNode.innerHTML += this.compiler.subSidebar(subMaxLevel) || "";
                } else {
                    this.compiler.subSidebar();
                }
                this._bindEventOnRendered(activeEl);
            }
            _bindEventOnRendered(activeEl) {
                const {autoHeader: autoHeader} = this.config;
                this.onRender();
                if (autoHeader && activeEl) {
                    const main = getNode("#main");
                    const firstNode = main.children[0];
                    if (firstNode && firstNode.tagName !== "H1") {
                        const h1 = this.compiler.header(activeEl.innerText, 1);
                        const wrapper = create("div", h1);
                        before(main, wrapper.children[0]);
                    }
                }
            }
            _renderNav(text) {
                text && this._renderTo("nav", this.compiler.compile(text));
                this.#addTextAsTitleAttribute("nav a");
            }
            _renderMain(text) {
                let opt = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
                let next = arguments.length > 2 ? arguments[2] : undefined;
                const {response: response} = this.route;
                if (response && !response.ok && (!text || response.status !== 404)) {
                    text = `# ${response.status} - ${response.statusText}`;
                }
                this.callHook("beforeEach", text, (result => {
                    let html;
                    const callback = () => {
                        if (opt.updatedAt) {
                            html = this.#formatUpdated(html, opt.updatedAt, this.config.formatUpdated);
                        }
                        this.callHook("afterEach", html, (hookData => {
                            this.#renderMain(hookData);
                            next();
                        }));
                    };
                    if (this.isHTML) {
                        html = this.result = text;
                        callback();
                    } else {
                        prerenderEmbed({
                            compiler: this.compiler,
                            raw: result
                        }, (tokens => {
                            html = this.compiler.compile(tokens);
                            callback();
                        }));
                    }
                }));
            }
            _renderCover(text, coverOnly) {
                const el = getNode(".cover");
                toggleClass(getNode("main"), coverOnly ? "add" : "remove", "hidden");
                if (!text) {
                    toggleClass(el, "remove", "show");
                    return;
                }
                toggleClass(el, "add", "show");
                let html = this.coverIsHTML ? text : this.compiler.cover(text);
                const m = html.trim().match('<p><img.*?data-origin="(.*?)"[^a]+alt="(.*?)">([^<]*?)</p>$');
                if (m) {
                    if (m[2] === "color") {
                        el.style.background = m[1] + (m[3] || "");
                    } else {
                        let path = m[1];
                        toggleClass(el, "add", "has-mask");
                        if (!isAbsolutePath(m[1])) {
                            path = getPath(this.router.getBasePath(), m[1]);
                        }
                        el.style.backgroundImage = `url(${path})`;
                        el.style.backgroundSize = "cover";
                        el.style.backgroundPosition = "center center";
                    }
                    html = html.replace(m[0], "");
                }
                this._renderTo(".cover-main", html);
            }
            _updateRender() {
                this.#renderNameLink(this);
                this.#renderSkipLink(this);
            }
            initRender() {
                const config = this.config;
                this.compiler = new Compiler(config, this.router);
                {
                    window.__current_docsify_compiler__ = this.compiler;
                }
                const id = config.el || "#app";
                const el = find(id);
                if (el) {
                    let html = "";
                    if (config.repo) {
                        html += corner(config.repo, config.cornerExternalLinkTarget);
                    }
                    if (config.coverpage) {
                        html += cover();
                    }
                    if (config.logo) {
                        const isBase64 = /^data:image/.test(config.logo);
                        const isExternal = /(?:http[s]?:)?\/\//.test(config.logo);
                        const isRelative = /^\./.test(config.logo);
                        if (!isBase64 && !isExternal && !isRelative) {
                            config.logo = getPath(this.router.getBasePath(), config.logo);
                        }
                    }
                    html += main(config);
                    this._renderTo(el, html, true);
                } else {
                    this.rendered = true;
                }
                if (config.loadNavbar) {
                    const navEl = find("nav") || create("nav");
                    const isMergedSidebar = config.mergeNavbar && isMobile;
                    navEl.setAttribute("aria-label", "secondary");
                    if (isMergedSidebar) {
                        find(".sidebar").prepend(navEl);
                    } else {
                        body.prepend(navEl);
                        navEl.classList.add("app-nav");
                        navEl.classList.toggle("no-badge", !config.repo);
                    }
                }
                if (config.themeColor) {
                    $.head.appendChild(create("div", theme(config.themeColor)).firstElementChild);
                }
                this._updateRender();
                toggleClass(body, "ready");
            }
        };
    }
    function Fetch(Base) {
        return class Fetch extends Base {
            #loadNested(path, qs, file, next, vm, first) {
                path = first ? path : path.replace(/\/$/, "");
                path = getParentPath(path);
                if (!path) {
                    return;
                }
                get(vm.router.getFile(path + file) + qs, false, vm.config.requestHeaders).then(next, (_error => this.#loadNested(path, qs, file, next, vm)));
            }
            #last;
            #abort=() => this.#last && this.#last.abort && this.#last.abort();
            #request=(url, requestHeaders) => {
                this.#abort();
                this.#last = get(url, true, requestHeaders);
                return this.#last;
            };
            #get404Path=(path, config) => {
                const {notFoundPage: notFoundPage, ext: ext} = config;
                const defaultPath = "_404" + (ext || ".md");
                let key;
                let path404;
                switch (typeof notFoundPage) {
                    case "boolean":
                        path404 = defaultPath;
                        break;

                    case "string":
                        path404 = notFoundPage;
                        break;

                    case "object":
                        key = Object.keys(notFoundPage).sort(((a, b) => b.length - a.length)).filter((k => path.match(new RegExp("^" + k))))[0];
                        path404 = key && notFoundPage[key] || defaultPath;
                        break;
                }
                return path404;
            };
            _loadSideAndNav(path, qs, loadSidebar, cb) {
                return () => {
                    if (!loadSidebar) {
                        return cb();
                    }
                    const fn = result => {
                        this._renderSidebar(result);
                        cb();
                    };
                    this.#loadNested(path, qs, loadSidebar, fn, this, true);
                };
            }
            _fetch() {
                let cb = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : noop;
                const {query: query} = this.route;
                const {path: path} = this.route;
                if (isExternal(path)) {
                    history.replaceState(null, "", "#");
                    this.router.normalize();
                } else {
                    const qs = stringifyQuery(query, [ "id" ]);
                    const {loadNavbar: loadNavbar, requestHeaders: requestHeaders, loadSidebar: loadSidebar} = this.config;
                    const file = this.router.getFile(path);
                    this.isRemoteUrl = isExternal(file);
                    this.isHTML = /\.html$/g.test(file);
                    const contentFetched = (text, opt, response) => {
                        this.route.response = response;
                        this._renderMain(text, opt, this._loadSideAndNav(path, qs, loadSidebar, cb));
                    };
                    const contentFailedToFetch = (_error, response) => {
                        this.route.response = response;
                        this._fetchFallbackPage(path, qs, cb) || this._fetch404(file, qs, cb);
                    };
                    if (!this.isRemoteUrl) {
                        this.matchVirtualRoute(path).then((contents => {
                            if (typeof contents === "string") {
                                contentFetched(contents);
                            } else {
                                this.#request(file + qs, requestHeaders).then(contentFetched, contentFailedToFetch);
                            }
                        }));
                    } else {
                        this.#request(file + qs, requestHeaders).then(contentFetched, contentFailedToFetch);
                    }
                    loadNavbar && this.#loadNested(path, qs, loadNavbar, (text => this._renderNav(text)), this, true);
                }
            }
            _fetchCover() {
                const {coverpage: coverpage, requestHeaders: requestHeaders} = this.config;
                const query = this.route.query;
                const root = getParentPath(this.route.path);
                if (coverpage) {
                    let path = null;
                    const routePath = this.route.path;
                    if (typeof coverpage === "string") {
                        if (routePath === "/") {
                            path = coverpage;
                        }
                    } else if (Array.isArray(coverpage)) {
                        path = coverpage.indexOf(routePath) > -1 && "_coverpage";
                    } else {
                        const cover = coverpage[routePath];
                        path = cover === true ? "_coverpage" : cover;
                    }
                    const coverOnly = Boolean(path) && this.config.onlyCover;
                    if (path) {
                        path = this.router.getFile(root + path);
                        this.coverIsHTML = /\.html$/g.test(path);
                        get(path + stringifyQuery(query, [ "id" ]), false, requestHeaders).then((text => this._renderCover(text, coverOnly)));
                    } else {
                        this._renderCover(null, coverOnly);
                    }
                    return coverOnly;
                }
            }
            $fetch() {
                let cb = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : noop;
                let onNavigate = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : this.onNavigate.bind(this);
                const done = () => {
                    this.callHook("doneEach");
                    cb();
                };
                const onlyCover = this._fetchCover();
                if (onlyCover) {
                    done();
                } else {
                    this._fetch((() => {
                        onNavigate();
                        done();
                    }));
                }
            }
            _fetchFallbackPage(path, qs) {
                let cb = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : noop;
                const {requestHeaders: requestHeaders, fallbackLanguages: fallbackLanguages, loadSidebar: loadSidebar} = this.config;
                if (!fallbackLanguages) {
                    return false;
                }
                const local = path.split("/")[1];
                if (fallbackLanguages.indexOf(local) === -1) {
                    return false;
                }
                const newPath = this.router.getFile(path.replace(new RegExp(`^/${local}`), ""));
                const req = this.#request(newPath + qs, requestHeaders);
                req.then(((text, opt) => this._renderMain(text, opt, this._loadSideAndNav(path, qs, loadSidebar, cb))), (_error => this._fetch404(path, qs, cb)));
                return true;
            }
            _fetch404(path, qs) {
                let cb = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : noop;
                const {loadSidebar: loadSidebar, requestHeaders: requestHeaders, notFoundPage: notFoundPage} = this.config;
                const fnLoadSideAndNav = this._loadSideAndNav(path, qs, loadSidebar, cb);
                if (notFoundPage) {
                    const path404 = this.#get404Path(path, this.config);
                    this.#request(this.router.getFile(path404), requestHeaders).then(((text, opt) => this._renderMain(text, opt, fnLoadSideAndNav)), (_error => this._renderMain(null, {}, fnLoadSideAndNav)));
                    return true;
                }
                this._renderMain(null, {}, fnLoadSideAndNav);
                return false;
            }
            initFetch() {
                this.config;
                this.$fetch((_ => this.callHook("ready")));
            }
        };
    }
    function Events(Base) {
        return class Events extends Base {
            #intersectionObserver;
            #isScrolling;
            #title=$.title;
            initEvent() {
                const {topMargin: topMargin} = this.config;
                if (topMargin) {
                    document.documentElement.style.setProperty("scroll-padding-top", `${topMargin}px`);
                }
                this.#initCover();
                this.#initSkipToContent("#skip-to-content");
                this.#initSidebarCollapse(".sidebar");
                this.#initSidebarToggle("button.sidebar-toggle");
                this.#initKeyBindings();
            }
            #initCover() {
                const coverElm = find("section.cover");
                if (!coverElm) {
                    toggleClass(body, "add", "sticky");
                    return;
                }
                const observer = new IntersectionObserver((entries => {
                    const isIntersecting = entries[0].isIntersecting;
                    const op = isIntersecting ? "remove" : "add";
                    toggleClass(body, op, "sticky");
                }));
                observer.observe(coverElm);
            }
            #initHeadings() {
                const headingElms = findAll("#main :where(h1, h2, h3, h4, h5)");
                const headingsInView = new Set;
                this.#intersectionObserver?.disconnect();
                this.#intersectionObserver = new IntersectionObserver((entries => {
                    if (this.#isScrolling) {
                        return;
                    }
                    for (const entry of entries) {
                        const op = entry.isIntersecting ? "add" : "delete";
                        headingsInView[op](entry.target);
                    }
                    const activeHeading = headingsInView.size > 1 ? Array.from(headingsInView).sort(((a, b) => a.compareDocumentPosition(b) & Node.DOCUMENT_POSITION_FOLLOWING ? -1 : 1))[0] : headingsInView.values().next().value;
                    if (activeHeading) {
                        const id = activeHeading.getAttribute("id");
                        const href = this.router.toURL(this.router.getCurrentPath(), {
                            id: id
                        });
                        const newSidebarActiveElm = this.#markSidebarActiveElm(href);
                        newSidebarActiveElm?.scrollIntoView({
                            behavior: "instant",
                            block: "nearest",
                            inline: "nearest"
                        });
                    }
                }), {
                    rootMargin: "0% 0% -50% 0%"
                });
                headingElms.forEach((elm => {
                    this.#intersectionObserver.observe(elm);
                }));
            }
            #initKeyBindings() {
                const {keyBindings: keyBindings} = this.config;
                const modifierKeys = [ "alt", "ctrl", "meta", "shift" ];
                if (keyBindings && keyBindings.constructor === Object) {
                    Object.values(keyBindings || []).forEach((bindingConfig => {
                        const {bindings: bindings} = bindingConfig;
                        if (!bindings) {
                            return;
                        }
                        bindingConfig.bindings = Array.isArray(bindings) ? bindings : [ bindings ];
                        bindingConfig.bindings = bindingConfig.bindings.map((keys => {
                            const sortedKeys = [ [], [] ];
                            if (typeof keys === "string") {
                                keys = keys.split("+");
                            }
                            keys.forEach((key => {
                                const isModifierKey = modifierKeys.includes(key);
                                const targetArray = sortedKeys[isModifierKey ? 0 : 1];
                                const newKeyValue = key.trim().toLowerCase();
                                targetArray.push(newKeyValue);
                            }));
                            sortedKeys.forEach((arr => arr.sort()));
                            return sortedKeys.flat();
                        }));
                    }));
                    on("keydown", (e => {
                        const isTextEntry = document.activeElement.matches("input, select, textarea");
                        if (isTextEntry) {
                            return;
                        }
                        const bindingConfigs = Object.values(keyBindings || []);
                        const matchingConfigs = bindingConfigs.filter((_ref => {
                            let {bindings: bindings} = _ref;
                            return bindings && bindings.some((keys => keys.every((k => modifierKeys.includes(k) && e[k + "Key"] || e.key === k || e.code.toLowerCase() === k || e.code.toLowerCase() === `key${k}`))));
                        }));
                        matchingConfigs.forEach((_ref2 => {
                            let {callback: callback} = _ref2;
                            e.preventDefault();
                            callback(e);
                        }));
                    }));
                }
            }
            #initSidebarCollapse(elm) {
                elm = typeof elm === "string" ? document.querySelector(elm) : elm;
                if (!elm) {
                    return;
                }
                on(elm, "click", (_ref3 => {
                    let {target: target} = _ref3;
                    if (target.nodeName === "A" && target.nextSibling && target.nextSibling.classList && target.nextSibling.classList.contains("app-sub-sidebar")) {
                        toggleClass(target.parentNode, "collapse");
                    }
                }));
            }
            #initSidebarToggle(elm) {
                elm = typeof elm === "string" ? document.querySelector(elm) : elm;
                if (!elm) {
                    return;
                }
                const toggle = () => {
                    body.classList.toggle("close");
                    const isClosed = isMobile ? body.classList.contains("close") : !body.classList.contains("close");
                    elm.setAttribute("aria-expanded", isClosed);
                };
                on(elm, "click", (e => {
                    e.stopPropagation();
                    toggle();
                }));
                isMobile && on(body, "click", (() => body.classList.contains("close") && toggle()));
            }
            #initSkipToContent(elm) {
                elm = typeof elm === "string" ? document.querySelector(elm) : elm;
                if (!elm) {
                    return;
                }
                elm.addEventListener("click", (evt => {
                    evt.preventDefault();
                    find("main")?.scrollIntoView({
                        behavior: "smooth"
                    });
                    this.#focusContent();
                }));
            }
            onRender() {
                const currentPath = this.router.toURL(this.router.getCurrentPath());
                const currentTitle = find(`.sidebar a[href='${currentPath}']`)?.innerText;
                $.title = currentTitle || this.#title;
                this.#markAppNavActiveElm();
                this.#markSidebarCurrentPage();
                this.#initHeadings();
            }
            onNavigate(source) {
                const {auto2top: auto2top, topMargin: topMargin} = this.config;
                const {query: query} = this.route;
                this.#markSidebarActiveElm();
                if (source !== "history") {
                    if (query.id) {
                        const headingElm = find(`.markdown-section :where(h1, h2, h3, h4, h5)[id="${query.id}"]`);
                        if (headingElm) {
                            this.#watchNextScroll();
                            headingElm.scrollIntoView({
                                behavior: "smooth",
                                block: "start"
                            });
                        }
                    } else if (source === "navigate") {
                        if (auto2top) {
                            document.scrollingElement.scrollTop = topMargin ?? 0;
                        }
                    }
                }
                if (query.id || source === "navigate") {
                    this.#focusContent();
                }
            }
            #focusContent() {
                let options = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
                const settings = {
                    preventScroll: true,
                    ...options
                };
                const {query: query} = this.route;
                const focusEl = query.id ? find(`#${query.id}`) : find("#main :where(h1, h2, h3, h4, h5, h6)") || find("#main");
                focusEl?.focus(settings);
            }
            #markAppNavActiveElm() {
                const href = decodeURIComponent(this.router.toURL(this.route.path));
                const navElm = find("nav.app-nav");
                if (!navElm) {
                    return;
                }
                const newActive = findAll(navElm, "a").sort(((a, b) => b.href.length - a.href.length)).find((a => href.includes(a.getAttribute("href")) || href.includes(decodeURI(a.getAttribute("href")))));
                const oldActive = find(navElm, "li.active");
                if (newActive && newActive !== oldActive) {
                    oldActive?.classList.remove("active");
                    newActive.classList.add("active");
                }
                return newActive;
            }
            #markSidebarActiveElm(href) {
                href ??= this.router.toURL(this.router.getCurrentPath());
                const sidebar = find(".sidebar");
                if (!sidebar) {
                    return;
                }
                const oldActive = find(sidebar, "li.active");
                const newActive = find(sidebar, `a[href="${href}"], a[href="${decodeURIComponent(href)}"]`)?.closest("li");
                if (newActive && newActive !== oldActive) {
                    oldActive?.classList.remove("active");
                    newActive.classList.add("active");
                }
                return newActive;
            }
            #markSidebarCurrentPage(href) {
                href ??= this.router.toURL(this.route.path);
                const sidebar = find(".sidebar");
                if (!sidebar) {
                    return;
                }
                const path = href?.split("?")[0];
                const oldPage = find(sidebar, "li[aria-current]");
                const newPage = find(sidebar, `a[href="${path}"], a[href="${decodeURIComponent(path)}"]`)?.closest("li");
                if (newPage && newPage !== oldPage) {
                    oldPage?.removeAttribute("aria-current");
                    newPage.setAttribute("aria-current", "page");
                }
                return newPage;
            }
            #watchNextScroll() {
                document.addEventListener("scroll", (() => {
                    this.#isScrolling = true;
                    if ("onscrollend" in window) {
                        document.addEventListener("scrollend", (() => this.#isScrolling = false), {
                            once: true
                        });
                    } else {
                        const callback = () => {
                            clearTimeout(scrollTimer);
                            scrollTimer = setTimeout((() => {
                                document.removeEventListener("scroll", callback);
                                this.#isScrolling = false;
                            }), 100);
                        };
                        let scrollTimer;
                        document.addEventListener("scroll", callback, false);
                    }
                }), {
                    once: true
                });
            }
        };
    }
    function makeExactMatcher(matcher) {
        const matcherWithBeginningOfInput = matcher.startsWith("^") ? matcher : `^${matcher}`;
        const matcherWithBeginningAndEndOfInput = matcherWithBeginningOfInput.endsWith("$") ? matcherWithBeginningOfInput : `${matcherWithBeginningOfInput}$`;
        return matcherWithBeginningAndEndOfInput;
    }
    function createNextFunction() {
        let storedCb = () => null;
        function next(value) {
            storedCb(value);
        }
        function onNext(cb) {
            storedCb = cb;
        }
        return [ next, onNext ];
    }
    function VirtualRoutes(Base) {
        return class VirtualRoutes extends Base {
            routes() {
                return this.config.routes || {};
            }
            matchVirtualRoute(path) {
                const virtualRoutes = this.routes();
                const virtualRoutePaths = Object.keys(virtualRoutes);
                let done = () => null;
                function asyncMatchNextRoute() {
                    const virtualRoutePath = virtualRoutePaths.shift();
                    if (!virtualRoutePath) {
                        return done(null);
                    }
                    const matcher = makeExactMatcher(virtualRoutePath);
                    const matched = path.match(matcher);
                    if (!matched) {
                        return asyncMatchNextRoute();
                    }
                    const virtualRouteContentOrFn = virtualRoutes[virtualRoutePath];
                    if (typeof virtualRouteContentOrFn === "string") {
                        const contents = virtualRouteContentOrFn;
                        return done(contents);
                    }
                    if (typeof virtualRouteContentOrFn === "function") {
                        const fn = virtualRouteContentOrFn;
                        const [next, onNext] = createNextFunction();
                        onNext((contents => {
                            if (typeof contents === "string") {
                                return done(contents);
                            } else if (contents === false) {
                                return done(null);
                            } else {
                                return asyncMatchNextRoute();
                            }
                        }));
                        if (fn.length <= 2) {
                            const returnedValue = fn(path, matched);
                            return next(returnedValue);
                        } else {
                            return fn(path, matched, next);
                        }
                    }
                    return asyncMatchNextRoute();
                }
                return {
                    then(cb) {
                        done = cb;
                        asyncMatchNextRoute();
                    }
                };
            }
        };
    }
    const currentScript = document.currentScript;
    function config(vm) {
        const config = Object.assign({
            auto2top: false,
            autoHeader: false,
            basePath: "",
            catchPluginErrors: true,
            cornerExternalLinkTarget: "_blank",
            coverpage: "",
            el: "#app",
            executeScript: null,
            ext: ".md",
            externalLinkRel: "noopener",
            externalLinkTarget: "_blank",
            formatUpdated: "",
            ga: "",
            homepage: "README.md",
            loadNavbar: null,
            loadSidebar: null,
            maxLevel: 6,
            mergeNavbar: false,
            name: "",
            nameLink: window.location.pathname,
            nativeEmoji: false,
            noCompileLinks: [],
            noEmoji: false,
            notFoundPage: false,
            plugins: [],
            relativePath: false,
            repo: "",
            routes: {},
            routerMode: "hash",
            subMaxLevel: 0,
            topMargin: 0,
            __themeColor: "",
            get themeColor() {
                return this.__themeColor;
            },
            set themeColor(value) {
                if (value) {
                    this.__themeColor = value;
                    console.warn(stripIndent(`\n              $docsify.themeColor is deprecated. Use a --theme-color property in your style sheet. Example:\n              <style>\n                :root {\n                  --theme-color: deeppink;\n                }\n              </style>\n            `).trim());
                }
            }
        }, typeof window.$docsify === "function" ? window.$docsify(vm) : window.$docsify);
        if (config.keyBindings !== false) {
            config.keyBindings = Object.assign({
                toggleSidebar: {
                    bindings: [ "\\" ],
                    callback(e) {
                        const toggleElm = document.querySelector(".sidebar-toggle");
                        if (toggleElm) {
                            toggleElm.click();
                            toggleElm.focus();
                        }
                    }
                }
            }, config.keyBindings);
        }
        const script = currentScript || Array.from(document.getElementsByTagName("script")).filter((n => /docsify\./.test(n.src)))[0];
        if (script) {
            for (const prop of Object.keys(config)) {
                const val = script.getAttribute("data-" + hyphenate(prop));
                if (isPrimitive(val)) {
                    config[prop] = val === "" ? true : val;
                }
            }
        }
        if (config.loadSidebar === true) {
            config.loadSidebar = "_sidebar" + config.ext;
        }
        if (config.loadNavbar === true) {
            config.loadNavbar = "_navbar" + config.ext;
        }
        if (config.coverpage === true) {
            config.coverpage = "_coverpage" + config.ext;
        }
        if (config.repo === true) {
            config.repo = "";
        }
        if (config.name === true) {
            config.name = "";
        }
        window.$docsify = config;
        return config;
    }
    function Lifecycle(Base) {
        return class Lifecycle extends Base {
            _hooks={};
            _lifecycle={};
            initLifecycle() {
                const hooks = [ "init", "mounted", "beforeEach", "afterEach", "doneEach", "ready" ];
                hooks.forEach((hook => {
                    const arr = this._hooks[hook] = [];
                    this._lifecycle[hook] = fn => arr.push(fn);
                }));
            }
            callHook(hookName, data) {
                let next = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : noop;
                const queue = this._hooks[hookName];
                const catchPluginErrors = this.config.catchPluginErrors;
                const step = function(index) {
                    const hookFn = queue[index];
                    if (index >= queue.length) {
                        next(data);
                    } else if (typeof hookFn === "function") {
                        const errTitle = "Docsify plugin error";
                        if (hookFn.length === 2) {
                            try {
                                hookFn(data, (result => {
                                    data = result;
                                    step(index + 1);
                                }));
                            } catch (err) {
                                if (catchPluginErrors) {
                                    console.error(errTitle, err);
                                } else {
                                    throw err;
                                }
                                step(index + 1);
                            }
                        } else {
                            try {
                                const result = hookFn(data);
                                data = result === undefined ? data : result;
                                step(index + 1);
                            } catch (err) {
                                if (catchPluginErrors) {
                                    console.error(errTitle, err);
                                } else {
                                    throw err;
                                }
                                step(index + 1);
                            }
                        }
                    } else {
                        step(index + 1);
                    }
                };
                step(0);
            }
        };
    }
    class Docsify extends(Fetch(Events(Render(VirtualRoutes(Router(Lifecycle(Object))))))){
        config=config(this);
        constructor() {
            super();
            this.initLifecycle();
            this.initPlugin();
            this.callHook("init");
            this.initRouter();
            this.initRender();
            this.initEvent();
            this.initFetch();
            this.callHook("mounted");
        }
        initPlugin() {
            this.config.plugins.forEach((fn => {
                try {
                    isFn(fn) && fn(this._lifecycle, this);
                } catch (err) {
                    if (this.config.catchPluginErrors) {
                        const errTitle = "Docsify plugin error";
                        console.error(errTitle, err);
                    } else {
                        throw err;
                    }
                }
            }));
        }
    }
    var util = Object.freeze({
        __proto__: null,
        cached: cached$1,
        cleanPath: cleanPath,
        getParentPath: getParentPath,
        getPath: getPath,
        hyphenate: hyphenate,
        inBrowser: inBrowser,
        isAbsolutePath: isAbsolutePath,
        isExternal: isExternal,
        isFn: isFn,
        isMobile: isMobile,
        isPrimitive: isPrimitive,
        noop: noop,
        parseQuery: parseQuery,
        removeParams: removeParams,
        replaceSlug: replaceSlug,
        resolvePath: resolvePath,
        stringifyQuery: stringifyQuery
    });
    function initGlobalAPI() {
        window.Docsify = {
            util: util,
            dom: dom,
            get: get,
            slugify: slugify,
            version: "4.13.0"
        };
        window.DocsifyCompiler = Compiler;
        window.marked = marked;
        window.Prism = prism;
    }
    initGlobalAPI();
    documentReady((() => new Docsify));
})();
