import { useState } from 'react'
import './App.css'

function App() {
  const [haiku, setHaiku] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const getHaiku = async () => {
    setLoading(true)
    setError(null)
    try {
      const res = await fetch(`${import.meta.env.VITE_API_URL}/haiku`)
      const data = await res.json()
      setHaiku(data.text)
    } catch (err) {
      setError('Failed to fetch haiku 😢')
    }
    setLoading(false)
  }

  return (
    <div className="container">
      <h1>DevOps Haiku Generator.</h1>
      <button onClick={getHaiku} disabled={loading}>
        {loading ? 'Generating...' : 'Get Haiku'}
      </button>
      {error && <p className="error">{error}</p>}
      {haiku && <pre className="haiku">{haiku}</pre>}
    </div>
  )
}

export default App
