"use client";

import { useState } from "react";
import { FaClipboardCheck } from "react-icons/fa";
import { FaRegClipboard } from "react-icons/fa6";

export default function InstallCommand() {
  const [copied, setCopied] = useState<boolean>(false);
  
  const command = `bash -c "$(curl -sLo- https://commit-buddy.vercel.app/install.sh)"`
  return (
    <div className="mt-12 flex gap-2 flex-col items-start">
      <p className="font-bold">Get Started:</p>
      <div
        id="installcommand"
        className="w-full max-w-max gap-2 flex items-center justify-start px-4 border h-12 rounded-md"
      >
      
        <p className="text-sm">$ {command}</p>
        {!copied ? (
          <FaRegClipboard
            onClick={async () => {
              await navigator.clipboard.writeText(command);
              setCopied(true);
            }}
          />
        ) : (
          <FaClipboardCheck />
        )}
      </div>

      {/*<Button className="w-36">Hey</Button>*/}
    </div>
  );
}
