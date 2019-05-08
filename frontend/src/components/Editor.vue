<template>
  <div>
    <editor-menu-bar :editor="editor">
      <div slot-scope="{ commands, isActive }" class="menu-bar">

        <button :class="{ 'is-active': isActive.bold() }" @click="commands.bold">
          b
        </button>

        <button :class="{ 'is-active': isActive.italic() }" @click="commands.italic">
          i
        </button>

        <button @click="commands.underline">u</button>

        <button :class="{ 'is-active': true }" @click="commands.code_block">
          &lt;&gt;
        </button>

      </div>
    </editor-menu-bar>
    <editor-content :editor="editor" />
  </div>
</template>

<script>
import { Editor, EditorContent, EditorMenuBar } from 'tiptap'
import {
  Blockquote,
  CodeBlock,
  HardBreak,
  Heading,
  OrderedList,
  BulletList,
  ListItem,
  TodoItem,
  TodoList,
  Bold,
  Code,
  Italic,
  Link,
  Strike,
  Underline,
  History,
} from 'tiptap-extensions'

export default {
  components: {
    EditorMenuBar,
    EditorContent,
  },
  data() {
    return {
      editor: new Editor({
        extensions: [
          new Blockquote(),
          new CodeBlock(),
          new HardBreak(),
          new Heading({ levels: [1, 2, 3] }),
          new BulletList(),
          new OrderedList(),
          new ListItem(),
          new TodoItem(),
          new TodoList(),
          new Bold(),
          new Code(),
          new Italic(),
          new Link(),
          new Strike(),
          new Underline(),
          new History(),
        ],
        content: `
          <h1>Yay Headlines!</h1>
          <p>All these <strong>cool tags</strong> are working now.</p>
        `,
      }),
    }
  },
  beforeDestroy() {
    this.editor.destroy()
  },
}
</script>