"use client";

import { Button } from "@workspace/ui/components/button";
import { Input } from "@workspace/ui/components/input";

export default function ReginsterPage() {
  return (
    <div className='flex flex-col items-center justify-center h-screen'>
      <h1>Welcome Back</h1>
      <Input placeholder='User Name' className='mb-4 w-80' />
      <Input placeholder='Login ID' className='mb-4 w-80' />
      <Input type='password' placeholder='Password' className='mb-4 w-80' />
      <Input
        type='password'
        placeholder='Confirm Password'
        className='mb-4 w-80'
      />
      <Button className='w-80'>Register</Button>
    </div>
  );
}
