import PageHeader from "@/app/components/PageHeader";

const MainSection = () => {
  return (
    <div className='flex flex-col items-center justify-center h-full w-full text-center'>
      <p className='text-2xl  font-bold'>ご利用ありがとうございました</p>
      <p className='text-base'>お支払いのため、レジへお進みください</p>
    </div>
  );
};

export default function checkoutPage() {
  return (
    <div className='flex flex-col h-full'>
      <div className='basis-1/12 flex-shrink-0 flex align-middle justify-center'>
        <PageHeader title='チェックアウト' />
      </div>
      <div className='basis-11/12 grow overflow-auto bg-gray-50 pt-2'>
        <MainSection />
      </div>
    </div>
  );
}
