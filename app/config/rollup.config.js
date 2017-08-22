import json from 'rollup-plugin-json'
import resolve from 'rollup-plugin-node-resolve'
import babel from 'rollup-plugin-babel'

import bundleSize from './rollup-plugin-bundle-size'
import scssTask from './rollup.scss.config'
import uglifyTask from './rollup.uglify.config'

const distPath = 'dist/bundle'

const defaultConfig = {
  input: 'src/main.js',
  plugins: [
    json(),
    resolve({
      browser: true
    }),
    scssTask({
      distPath
    }),
    babel({
      exclude: 'node_modules/**'
    }),
    uglifyTask({
      distPath
    }),
    bundleSize()
  ],
  output: {
    file: `${distPath}.js`,
    format: 'iife',
    name: 'App',
    sourcemap: true
  },
  watch: {
    exclude: ['node_modules/**']
  }
}

export default defaultConfig
