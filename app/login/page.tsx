// app/page.tsx
'use client'

import { useState } from 'react'
import { Form } from 'radix-ui'
import Link from 'next/link'

export default function SignInPage() {
  const [inputValue, setInputValue] = useState('')

  return (
    <div className='flex min-h-screen flex-col items-center justify-center bg-white p-4'>
      <div className='w-full max-w-md space-y-6'>
        {/* Twitter Logo */}
        <div className='flex justify-center'>
          <svg width='40' height='32' viewBox='0 0 24 24' fill='#1D9BF0'>
            <path d='M23.643 4.937c-.835.37-1.732.62-2.675.733.962-.576 1.7-1.49 2.048-2.578-.9.534-1.897.922-2.958 1.13-.85-.904-2.06-1.47-3.4-1.47-2.572 0-4.658 2.086-4.658 4.66 0 .364.042.718.12 1.06-3.873-.195-7.304-2.05-9.602-4.868-.4.69-.63 1.49-.63 2.342 0 1.616.823 3.043 2.072 3.878-.764-.025-1.482-.234-2.11-.583v.06c0 2.257 1.605 4.14 3.737 4.568-.392.106-.803.162-1.227.162-.3 0-.593-.028-.877-.082.593 1.85 2.313 3.198 4.352 3.234-1.595 1.25-3.604 1.995-5.786 1.995-.376 0-.747-.022-1.112-.065 2.062 1.323 4.51 2.093 7.14 2.093 8.57 0 13.255-7.098 13.255-13.254 0-.2-.005-.402-.014-.602.91-.658 1.7-1.477 2.323-2.41z' />
          </svg>
        </div>

        {/* Sign in Header */}
        <h1 className='text-center text-3xl font-bold'>Twitter にログイン</h1>

        <div className='space-y-4'>
          {/* Google Sign in Button */}
          <button className='flex w-full items-center justify-center gap-3 rounded-full border border-gray-300 py-3 font-medium transition hover:bg-gray-50'>
            <svg width='20' height='20' viewBox='0 0 48 48'>
              <path
                fill='#EA4335'
                d='M24 9.5c3.54 0 6.71 1.22 9.21 3.6l6.85-6.85C35.9 2.38 30.47 0 24 0 14.62 0 6.51 5.38 2.56 13.22l7.98 6.19C12.43 13.72 17.74 9.5 24 9.5z'
              />
              <path
                fill='#4285F4'
                d='M46.98 24.55c0-1.57-.15-3.09-.38-4.55H24v9.02h12.94c-.58 2.96-2.26 5.48-4.78 7.18l7.73 6c4.51-4.18 7.09-10.36 7.09-17.65z'
              />
              <path
                fill='#FBBC05'
                d='M10.53 28.59c-.48-1.45-.76-2.99-.76-4.59s.27-3.14.76-4.59l-7.98-6.19C.92 16.46 0 20.12 0 24c0 3.88.92 7.54 2.56 10.78l7.97-6.19z'
              />
              <path
                fill='#34A853'
                d='M24 48c6.48 0 11.93-2.13 15.89-5.81l-7.73-6c-2.15 1.45-4.92 2.3-8.16 2.3-6.26 0-11.57-4.22-13.47-9.91l-7.98 6.19C6.51 42.62 14.62 48 24 48z'
              />
            </svg>
            <span>Sign in with Google</span>
          </button>

          {/* Separator */}
          <div className='flex items-center justify-center'>
            <div className='h-px bg-gray-300 w-full'></div>
            <span className='px-4 text-gray-500 text-sm'>Or</span>
            <div className='h-px bg-gray-300 w-full'></div>
          </div>

          {/* Email/Phone/Username Input */}
          <Form.Root className='w-full'>
            <Form.Field name='identifier'>
              <Form.Control asChild>
                <input
                  className='w-full rounded-md border border-gray-300 px-4 py-3 placeholder-gray-500 focus:border-blue-500 focus:outline-none'
                  type='text'
                  placeholder='Email'
                  value={inputValue}
                  onChange={(e) => setInputValue(e.target.value)}
                />
              </Form.Control>
            </Form.Field>
          </Form.Root>

          {/* Next Button */}
          <button className='w-full rounded-full bg-black py-3 font-medium text-white transition hover:bg-gray-800'>
            Next
          </button>

          {/* Forgot Password Button */}
          <button className='w-full rounded-full border border-gray-300 py-3 font-medium text-gray-700 transition hover:bg-gray-50'>
            Forgot password?
          </button>
        </div>

        {/* Sign up prompt */}
        <div className='text-center'>
          <span className='text-gray-600'>Don't have an account? </span>
          <Link href='/signup' className='text-blue-500 hover:underline'>
            Sign up
          </Link>
        </div>
      </div>
    </div>
  )
}
