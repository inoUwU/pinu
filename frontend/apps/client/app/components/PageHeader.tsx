import { ChevronLeft } from "lucide-react";
import { Button } from "@workspace/ui/components/button";

type Props = {
  title: string;
  showBackButton?: boolean;
};

const BackButton = () => {
  const onClick = () => {
    window.history.back();
  };
  return (
    <Button onClick={onClick} className="p-6 w-20 relative" variant="ghost">
      <ChevronLeft className="absolute left-3 top-1/2 transform -translate-y-1/2" />
    </Button>
  );
};

export default function PageHeader({ title, showBackButton }: Props) {
  return (
    <div className="grid grid-cols-12 items-center p-2">
      <div className="col-start-1 col-end-1">
        {showBackButton ? <BackButton /> : null}
      </div>
      <div className="col-start-2 col-end-12 flex items-center justify-center">
        <div className="text-3xl font-bold select-none">{title}</div>
      </div>
    </div>
  );
}
