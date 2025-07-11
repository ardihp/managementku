"use client";

import TextStyle from "@tiptap/extension-text-style";
import { EditorProvider, useCurrentEditor, useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Placeholder from "@tiptap/extension-placeholder";
import Link from "@tiptap/extension-link";

import React, { useCallback, useEffect } from "react";
import { ToggleGroup, ToggleGroupItem } from "./toggle-group";
import {
  Bold,
  Code,
  Italic,
  Link2,
  List,
  ListOrdered,
  Pilcrow,
  Redo2,
  RemoveFormatting,
  Strikethrough,
  Undo2,
  Unlink2,
} from "lucide-react";
import { cn } from "@/lib/utils";

const MenuBar = ({
  refetch = 0,
  value = "",
}: {
  refetch?: number;
  value?: string;
}) => {
  const { editor } = useCurrentEditor();

  useEffect(() => {
    if (refetch >= 1) {
      if (value === "") {
        editor?.destroy();
      } else {
        editor?.commands.clearContent();
        editor?.commands.setContent(value);
      }
    }
  }, [refetch]);

  if (!editor) {
    return null;
  }

  let activeToggle: string[] = [];

  editor.isActive("bold") && activeToggle.push("bold");
  editor.isActive("italic") && activeToggle.push("italic");
  editor.isActive("strike") && activeToggle.push("strike");
  editor.isActive("code") && activeToggle.push("code");
  editor.isActive("bulletList") && activeToggle.push("bullet-list");
  editor.isActive("orderedList") && activeToggle.push("ordered-list");
  editor.isActive("link") && activeToggle.push("link");

  const handleLink = useCallback(() => {
    const previousUrl = editor.getAttributes("link").href;
    const url = window.prompt("URL", previousUrl);

    // cancelled
    if (url === null) {
      return;
    }

    // empty
    if (url === "") {
      editor.chain().focus().extendMarkRange("link").unsetLink().run();

      return;
    }

    // update link
    try {
      editor
        .chain()
        .focus()
        .extendMarkRange("link")
        .setLink({ href: url })
        .run();
    } catch (error: any) {
      alert(error.message);
    }
  }, [editor]);

  return (
    <div className="flex items-center gap-4">
      <ToggleGroup type="single" variant="outline">
        <ToggleGroupItem
          value="undo"
          onClick={() => editor.chain().focus().undo().run()}
          disabled={!editor.can().chain().focus().undo().run()}
        >
          <Undo2 className="size-4" />
        </ToggleGroupItem>
        <ToggleGroupItem
          value="redo"
          onClick={() => editor.chain().focus().redo().run()}
          disabled={!editor.can().chain().focus().redo().run()}
        >
          <Redo2 className="size-4" />
        </ToggleGroupItem>
      </ToggleGroup>

      <ToggleGroup type="multiple" variant="outline" value={activeToggle}>
        <ToggleGroupItem
          value="bold"
          onClick={() => editor.chain().focus().toggleBold().run()}
          disabled={!editor.can().chain().focus().toggleBold().run()}
        >
          <Bold className="size-4" />
        </ToggleGroupItem>
        <ToggleGroupItem
          value="italic"
          onClick={() => editor.chain().focus().toggleItalic().run()}
          disabled={!editor.can().chain().focus().toggleItalic().run()}
        >
          <Italic className="size-4" />
        </ToggleGroupItem>
        <ToggleGroupItem
          value="strike"
          onClick={() => editor.chain().focus().toggleStrike().run()}
          disabled={!editor.can().chain().focus().toggleStrike().run()}
        >
          <Strikethrough className="size-4" />
        </ToggleGroupItem>
        <ToggleGroupItem
          value="code"
          onClick={() => editor.chain().focus().toggleCode().run()}
          disabled={!editor.can().chain().focus().toggleCode().run()}
        >
          <Code className="size-4" />
        </ToggleGroupItem>
        <ToggleGroupItem
          value="clear-mark"
          onClick={() => editor.chain().focus().unsetAllMarks().run()}
        >
          <RemoveFormatting />
        </ToggleGroupItem>
        <ToggleGroupItem
          value="paragraph"
          onClick={() => editor.chain().focus().setParagraph().run()}
        >
          <Pilcrow className="size-4" />
        </ToggleGroupItem>
        <ToggleGroupItem
          value="bullet-list"
          onClick={() => editor.chain().focus().toggleBulletList().run()}
        >
          <List className="size-4" />
        </ToggleGroupItem>
        <ToggleGroupItem
          value="ordered-list"
          onClick={() => editor.chain().focus().toggleOrderedList().run()}
        >
          <ListOrdered className="size-4" />
        </ToggleGroupItem>
        <ToggleGroupItem value="link" onClick={handleLink}>
          <Link2 className="size-4" />
        </ToggleGroupItem>
        {editor.isActive("link") && (
          <ToggleGroupItem
            value="unlink"
            onClick={() => editor.chain().focus().unsetLink().run()}
          >
            <Unlink2 className="size-4" />
          </ToggleGroupItem>
        )}
      </ToggleGroup>
    </div>
  );
};

const extensions = [
  TextStyle.configure({}),
  StarterKit.configure({
    bulletList: {
      keepMarks: true,
      keepAttributes: false, // TODO : Making this as `false` becase marks are not preserved when I try to preserve attrs, awaiting a bit of help
    },
    orderedList: {
      keepMarks: true,
      keepAttributes: false, // TODO : Making this as `false` becase marks are not preserved when I try to preserve attrs, awaiting a bit of help
    },
  }),
  Link.configure({
    openOnClick: false,
    autolink: true,
    defaultProtocol: "https",
    protocols: ["http", "https"],
    isAllowedUri: (url, ctx) => {
      try {
        // construct URL
        const parsedUrl = url.includes(":")
          ? new URL(url)
          : new URL(`${ctx.defaultProtocol}://${url}`);

        // use default validation
        if (!ctx.defaultValidate(parsedUrl.href)) {
          return false;
        }

        // disallowed protocols
        const disallowedProtocols = ["ftp", "file", "mailto"];
        const protocol = parsedUrl.protocol.replace(":", "");

        if (disallowedProtocols.includes(protocol)) {
          return false;
        }

        // only allow protocols specified in ctx.protocols
        const allowedProtocols = ctx.protocols.map((p) =>
          typeof p === "string" ? p : p.scheme
        );

        if (!allowedProtocols.includes(protocol)) {
          return false;
        }

        // disallowed domains
        const disallowedDomains = [
          "example-phishing.com",
          "malicious-site.net",
        ];
        const domain = parsedUrl.hostname;

        if (disallowedDomains.includes(domain)) {
          return false;
        }

        // all checks have passed
        return true;
      } catch {
        return false;
      }
    },
    shouldAutoLink: (url) => {
      try {
        // construct URL
        const parsedUrl = url.includes(":")
          ? new URL(url)
          : new URL(`https://${url}`);

        // only auto-link if the domain is not in the disallowed list
        const disallowedDomains = [
          "example-no-autolink.com",
          "another-no-autolink.com",
        ];
        const domain = parsedUrl.hostname;

        return !disallowedDomains.includes(domain);
      } catch {
        return false;
      }
    },
  }),
];

interface TextEditorProps {
  value: string;
  onValueChange: (val: string) => void;
  placeholder?: string;
  disabled?: boolean;
  editorClassName?: string;
  refetch?: number;
}

export default function TextEditor({
  value,
  onValueChange,
  placeholder,
  disabled = false,
  editorClassName,
  refetch,
}: TextEditorProps) {
  return (
    <EditorProvider
      editable={!disabled}
      editorContainerProps={{
        className: cn(
          "border rounded-lg px-3 py-2 shadow-xs focus-within:border-ring focus-within:ring-ring/50 focus-within:ring-[3px] mt-1 transition-[color,box-shadow]",
          editorClassName
        ),
      }}
      content={value}
      slotBefore={<MenuBar refetch={refetch} value={value} />}
      extensions={[
        ...extensions,
        Placeholder.configure({
          placeholder: placeholder,
        }),
      ]}
      onUpdate={({ editor }) => onValueChange(editor.getHTML())}
      immediatelyRender
    />
  );
}
