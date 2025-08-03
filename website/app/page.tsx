import InstallCommand from "@/components/InstallCommand";
import { Button, buttonVariants } from "@/components/ui/button";
import Image from "next/image";
import Link from "next/link";
import { BiClipboard } from "react-icons/bi";
import { FaRegClipboard } from "react-icons/fa";
import { FaGithub } from "react-icons/fa6";

export default function Home() {
  return (
    <div className="h-screen flex flex-col items-center gap-4 px-4">
      <nav className="w-full flex h-10 justify-end p-2 ">
        <Link href={"https://github.com/winnerx0/commit-buddy"}>
          <FaGithub className="self-end" />
        </Link>
      </nav>
      <section className="mt-48 w-full gap-4 flex flex-col items-center justify-center">
        <h1 className="text-4xl md:text-6xl lg:text-7xl font-bold">
          Commit Buddy
        </h1>
        <p className="tracking-wider italic text-center">
          Create detailed commit based on your git changes
        </p>
        <Link
          href={"https://github.com/winnerx0/commit-buddy"}
          className={buttonVariants({ className: "h-12 rounded-2xl mt-4" })}
        >
          Give A Star On Github
        </Link>
      </section>
      <InstallCommand />

      <div className="rounded-2xl flex flex-col gap-2 mt-24">
        <h1 className="font-bold text-2xl">Preview</h1>
        <Image
          src={"/image1.png"}
          width={1000}
          height={1000}
          alt="Image1"
          className="rounded-2xl"
        />
      </div>
    </div>
  );
}
