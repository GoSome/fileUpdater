module.exports = {
  chainWebpack: config => {
    const oneOfsMap = config.module.rule('scss').oneOfs.store
    oneOfsMap.forEach(item => {
      item
          .use('sass-resources-loader')
          .loader('sass-resources-loader')
          .options({
            resources: ['./src/scss/_base.scss']
          })
          .end()
    })
  },
  devServer: {
    port: 8088,
    proxy: {
      '/api/': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  }
}
