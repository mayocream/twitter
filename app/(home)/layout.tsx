'use client'

import { supabase } from '@/lib/supabase'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'

export default function Home() {
  const [loading, setLoading] = useState(true)
  const router = useRouter()

  const isAuthenticated = async () => {
    const {
      data: { session },
      error,
    } = await supabase.auth.getSession()
    if (error || !session) {
      return false
    }
    return true
  }

  useEffect(() => {
    isAuthenticated().then((isAuthenticated) => {
      if (!isAuthenticated) {
        console.log('User is not authenticated')
        router.push('/login')
      } else {
        setLoading(false)
      }
    })
  }, [])

  if (loading) {
    return null
  }

  return <></>
}
