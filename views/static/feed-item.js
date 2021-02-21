
class FeedItem extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({mode: 'open'});
  }

  connectedCallback() {
    const template = document.createElement('template');
    template.innerHTML = `
    <style>
    @font-face {
      font-family: "charter";
      src: url("https://glyph.medium.com/font/be78681/0-3j_4g_6bu_6c4_6c8_6c9_6cc_6cd_6ci_6cm/charter-400-normal.woff") format("woff");
      font-style: normal;
      font-weight: 400;
      unicode-range: U+0-7F, U+A0, U+200A, U+2014, U+2018, U+2019, U+201C, U+201D, U+2022, U+2026;
    }

    :host {
      width: 100% !important;
      overflow: scroll !important;
      overflow-x: auto !important;
      font-size: 18px;
    }
    * {
      max-width: 100% !important;
      height: auto !important;
      float: none !important;
    }
    table {
      width: 100% !important;
    }

    img {
      margin: auto auto !important;
    }
    h1, h2, h3, h4 {
      margin-top: 1.3em;
    }
    h1:first-of-type {
      margin-top: 0;
    }
    p, a {
      line-height: 1.2em;
    }
    p {
      font-family: charter, Georgia, "Times New Roman", Times, serif;
      font-style: normal;
      font-weight: 400;
      letter-spacing: -0.063px;
    }
    a {
      color: #333;
      font-weight: bold;
    }
    :host(.dark) a {
      color: #ccc;
    }
    a:hover, :host(.dark) a:hover {
      color: #ff2e88;
    }

    pre {
      overflow-x: scroll;
      padding: 8px;
      background: #62848463;
    }
    pre code {
      margin-right: 8px;
    }
    p code {
      background: #62848463;
      padding: 0 4px;
    }

    </style>
    `;

    fetch(`/api/item/${this.getAttribute('item-id')}`)
      .then(res => res.json())
      .then(item => {
        template.innerHTML += `<h1><a href="${item.URL}" target="_blank" rel="noopener">${item.Title}</a></h1>`;
        template.innerHTML += item.Content || item.Description;
        this.shadowRoot.appendChild(template.content.cloneNode(true));
        [...this.shadowRoot.querySelectorAll('a[href^=http]')].forEach(a => {
          a.setAttribute("target", "_blank");
          a.setAttribute("rel", "noopener");
        });
        [...this.shadowRoot.querySelectorAll('p')].forEach(p => {
          if (p.innerText.trim() == "") {
            p.remove();
          }
        });

        let url = new URL(item.URL);
        [...this.shadowRoot.querySelectorAll('img[src^="/"]')].forEach(i => {
          i.src = url.origin + i.getAttribute('src');
        });
        [...this.shadowRoot.querySelectorAll('a[href^="/"]')].forEach(a => {
          a.href = url.origin + a.getAttribute('src');
        });
        [...this.shadowRoot.querySelectorAll('img:not([src^=http])')].forEach(i => {
          i.src = url.origin +'/'+ i.getAttribute('src');
        });
        [...this.shadowRoot.querySelectorAll('a:not([href^=http])')].forEach(a => {
          a.href = url.origin +'/'+ a.getAttribute('src');
        });
      })

  }
}
customElements.define('feed-item', FeedItem);
