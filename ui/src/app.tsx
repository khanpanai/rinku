import { useState } from 'preact/hooks'
import preactLogo from './assets/preact.svg'
import './app.css'

export function App() {
  const [count, setCount] = useState(0)

  return (
    <>
        <h1 className="text-3xl font-bold underline">
            Hello world!
        </h1>
    </>
  )
}
