export default function checkoutPage() {
  return (
    <div className='flex flex-col h-screen space-y-4'>
      <p className='text-2xl font-extrabold dark:text-white'>Checkout</p>
      <p className='text-3xl font-bold dark:text-white'>
        Thank you for your order!
      </p>
      <p>Please proceed to the cashier to complete your payment.</p>
    </div>
  );
}
