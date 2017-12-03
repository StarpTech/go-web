<app>
  <h1>{{ msg }}</h1>
  <script>
    this.on('mount', () => {
      this.update({
        msg: 'Modern Web Development with Go'
      })
    })
  </script>
  <style>
    h1 {
      color:aqua;
    }
  </style>
</app>
