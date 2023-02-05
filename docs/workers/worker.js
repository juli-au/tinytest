importScripts('./wasm_exec.js')

let cancelFunc

async function init() {
  const go = new Go()
  const source = await fetch('tinytest.wasm')
  const obj = await WebAssembly.instantiateStreaming(source, go.importObject)
  //const wasm = obj.instance

  self.onmessage = async (event) => {
    if (event.data === 'main') {
      console.log('starting main...')
      //setTimeout(() => go.run(obj.instance), 0)
      go.run(obj.instance)
      console.log('main started')
      return
    }
    if (event.data === 'cancel') {
      console.log('cancelling....')
      cancelFunc
      console.log('cancelled')
      return
    }
  }
}
init()
