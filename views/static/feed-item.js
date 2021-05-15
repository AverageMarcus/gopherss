
class FeedItem extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({mode: 'open'});
  }

  connectedCallback() {
    const template = document.createElement('template');
    template.innerHTML = `
    <style>

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
      font-family: "Atkinson Hyperlegible Bold";
      margin-top: 1.3em;
    }
    :root > h1 {
      margin-top: 0;
    }
    p, a {
      line-height: 1.2em;
    }
    p, li, div {
      font-family: "Atkinson Hyperlegible Regular";
      font-style: normal;
      font-weight: 400;
      letter-spacing: 0.05em
    }
    em {
      font-family: "Atkinson Hyperlegible Italic";
      font-style: normal;
    }
    strong {
      font-weight: 500;
      font-family: "Atkinson Hyperlegible Bold";
    }
    em strong, strong em {
      font-family: "Atkinson Hyperlegible BoldItalic";
    }
    li {
      margin: 0.6em 0;
    }
    a {
      color: #333;
      font-family: "Atkinson Hyperlegible Bold";
      font-weight: 500;
      letter-spacing: 0.05em
    }
    :host(.dark) a {
      color: #eee;
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

    iframe {
      display: block;
      width: 100%;
      min-height: 600px;
      border: none;
    }
    </style>
    `;

    fetch(`/api/item/${this.getAttribute('item-id')}`)
      .then(res => res.json())
      .then(item => {
        template.innerHTML += `<h1><a href="${item.URL}" target="_blank" rel="noopener">${item.Title}</a></h1>`;
        template.innerHTML += `<div class="feedContent">${item.Content || item.Description}</div>`;
        template.innerHTML += `<iframe style="display: none;" data-src="${item.URL}"></iframe>`
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
          a.href = url.origin + a.getAttribute('href');
        });
        [...this.shadowRoot.querySelectorAll('img:not([src^=http])')].forEach(i => {
          i.src = url.origin +'/'+ i.getAttribute('src');
        });
        [...this.shadowRoot.querySelectorAll('a:not([href^=http])')].forEach(a => {
          a.href = url.origin +'/'+ a.getAttribute('href');
        });
      })
  }

  showIframe() {
    if (this.shadowRoot.querySelector(".feedContent").style.display != "none") {
      this.shadowRoot.querySelector(".feedContent").style.display = "none";
      this.shadowRoot.querySelector("iframe").src = this.shadowRoot.querySelector("iframe").dataset.src;
      this.shadowRoot.querySelector("iframe").style.display = "block";
    } else {
      this.shadowRoot.querySelector(".feedContent").style.display = "block";
      this.shadowRoot.querySelector("iframe").style.display = "none";
    }
  }
}
customElements.define('feed-item', FeedItem);
