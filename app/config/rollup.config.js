import json from 'rollup-plugin-json'
import resolve from 'rollup-plugin-node-resolve'
import babel from 'rollup-plugin-babel'

import bundleSize from './rollup-plugin-bundle-size'
import scssTask from './rollup.scss.config'
import uglifyTask from './rollup.uglify.config'

const componentsDistPath = 'dist/components'
const userDistPath = 'dist/user'
const userDetailsDistPath = 'dist/user-details'

const componentConfig = {
  input: 'src/components.js',
  plugins: [
    json(),
    resolve({
      browser: true
    }),
    scssTask({
      distPath: componentsDistPath
    }),
    babel({
      exclude: 'node_modules/**'
    }),
    uglifyTask({
      distPath: componentsDistPath
    }),
    bundleSize()
  ],
  output: {
    file: `${componentsDistPath}.js`,
    format: 'iife',
    name: 'Components',
    sourcemap: true
  },
  watch: {
    exclude: ['node_modules/**']
  }
}

const userConfig = {
  input: 'src/user.js',
  plugins: [
    json(),
    resolve({
      browser: true
    }),
    scssTask({
      distPath: userDistPath
    }),
    babel({
      exclude: 'node_modules/**'
    }),
    uglifyTask({
      distPath: userDistPath
    }),
    bundleSize()
  ],
  output: {
    file: `${userDistPath}.js`,
    format: 'iife',
    name: 'User',
    sourcemap: true
  },
  watch: {
    exclude: ['node_modules/**']
  }
}

const userDetailsConfig = {
  input: 'src/user-details.js',
  plugins: [
    json(),
    resolve({
      browser: true
    }),
    scssTask({
      distPath: userDetailsDistPath
    }),
    babel({
      exclude: 'node_modules/**'
    }),
    uglifyTask({
      distPath: userDetailsDistPath
    }),
    bundleSize()
  ],
  output: {
    file: `${userDetailsDistPath}.js`,
    format: 'iife',
    name: 'UserDetails',
    sourcemap: true
  },
  watch: {
    exclude: ['node_modules/**']
  }
}

export default [componentConfig, userConfig, userDetailsConfig]
