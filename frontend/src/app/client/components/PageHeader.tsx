import { ArrowLeft } from "lucide-react";
import { Button } from "@/components/ui/button";

type Props = {
  title: string;
  showBackButton?: boolean;
};

const BackButton = () => {
  const onClick = () => {
    // TODO: Check if the history exists before going back.
    window.history.back();
  };
  return (
    <Button onClick={onClick} className='p-2'>
      <ArrowLeft />
    </Button>
  );
};

export default function PageHeader({ title, showBackButton }: Props) {
  return (
    <div className='grid grid-cols-12'>
      <div className='col-start-1 col-end-1'>
        {showBackButton ? <BackButton /> : null}
      </div>
      <div className='col-start-2 col-end-12 flex items-center justify-center'>
        <div className='text-3xl font-bold select-none'>{title}</div>
      </div>
    </div>
  );
}
