
class FeedItem extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({mode: 'open'});
  }

  load() {
    const template = document.createElement('template');
    template.innerHTML = `
    <style>
    :root {
      width: 100%;
      overflow: scroll;
      overflow-x: auto;
    }
    * {
      max-width: 100%;
      height: auto;
    }
    table {
      width: 100%;
    }
    img {
      margin: auto auto;
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
