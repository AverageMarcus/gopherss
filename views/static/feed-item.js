
class FeedItem extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({mode: 'open'});
  }

  load() {
    const template = document.createElement('template');
    template.innerHTML = `
    <style>
    :host {
      width: 100% !important;
      overflow: scroll !important;
      overflow-x: auto !important;
    }
    * {
      max-width: 100% !important;
      height: auto !important;
    }
    table {
      width: 100% !important;
    }
    img {
      margin: auto auto !important;
    }
    p {
      font-family: 'Roboto', sans-serif;
      font-size: 14px;
      line-height: 20px;
      letter-spacing: 0em;
      font-weight: 500;
    }
    a {
      color: #333;
      font-weight: bold;
    }
    :host(.dark) a {
      color: #ccc;
    }

    </style>
    `;

    if (!this.loaded) {
      return fetch(`/api/item/${this.getAttribute('item-id')}`)
        .then(res => res.json())
        .then(item => {
          template.innerHTML += item.Content || item.Description;
          this.shadowRoot.appendChild(template.content.cloneNode(true));
        })
        .then(() => this.loaded = true);
    }
  }
}
customElements.define('feed-item', FeedItem);
