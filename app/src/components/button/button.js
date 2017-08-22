class MyButton extends HTMLButtonElement {
  show () {
    alert(this.textContent)
  }
}

customElements.define('my-button', MyButton, {extends: 'button'})
