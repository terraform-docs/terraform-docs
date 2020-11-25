## terraform-docs xml

Generate XML of inputs and outputs

### Synopsis

Generate XML of inputs and outputs

```
terraform-docs xml [PATH] [flags]
```

### Options

```
  -h, --help   help for xml
```

### Options inherited from parent commands

```
  -c, --config string               config file name (default ".terraform-docs.yml")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [header, inputs, outputs, providers, requirements]
      --hide-all                    hide all sections (default false)
      --output-values               inject output values into outputs (default false)
      --output-values-from string   inject output values from file into outputs (default "")
      --show strings                show section [header, inputs, outputs, providers, requirements]
      --show-all                    show all sections (default true)
      --sort                        sort items (default true)
      --sort-by-required            sort items by name and print required ones first (default false)
      --sort-by-type                sort items by type of them (default false)
```

### Example

Given the [`examples`](/examples/) module:

```shell
terraform-docs xml ./examples/
```

generates the following output:

    <module>
      <header>Usage:&#xA;&#xA;Example of &#39;foo_bar&#39; module in `foo_bar.tf`.&#xA;&#xA;- list item 1&#xA;- list item 2&#xA;&#xA;Even inline **formatting** in _here_ is possible.&#xA;and some [link](https://domain.com/)&#xA;&#xA;* list item 3&#xA;* list item 4&#xA;&#xA;```hcl&#xA;module &#34;foo_bar&#34; {&#xA;  source = &#34;github.com/foo/bar&#34;&#xA;&#xA;  id   = &#34;1234567890&#34;&#xA;  name = &#34;baz&#34;&#xA;&#xA;  zones = [&#34;us-east-1&#34;, &#34;us-west-1&#34;]&#xA;&#xA;  tags = {&#xA;    Name         = &#34;baz&#34;&#xA;    Created-By   = &#34;first.last@email.com&#34;&#xA;    Date-Created = &#34;20180101&#34;&#xA;  }&#xA;}&#xA;```&#xA;&#xA;Here is some trailing text after code block,&#xA;followed by another line of text.&#xA;&#xA;| Name | Description     |&#xA;|------|-----------------|&#xA;| Foo  | Foo description |&#xA;| Bar  | Bar description |</header>
      <inputs>
        <input>
          <name>bool-1</name>
          <type>bool</type>
          <description>It&#39;s bool number one.</description>
          <default>true</default>
          <required>false</required>
        </input>
        <input>
          <name>bool-2</name>
          <type>bool</type>
          <description>It&#39;s bool number two.</description>
          <default>false</default>
          <required>false</required>
        </input>
        <input>
          <name>bool-3</name>
          <type>bool</type>
          <description xsi:nil="true"></description>
          <default>true</default>
          <required>false</required>
        </input>
        <input>
          <name>bool_default_false</name>
          <type>bool</type>
          <description xsi:nil="true"></description>
          <default>false</default>
          <required>false</required>
        </input>
        <input>
          <name>input-with-code-block</name>
          <type>list</type>
          <description>This is a complicated one. We need a newline.  &#xA;And an example in a code block&#xA;```&#xA;default     = [&#xA;  &#34;machine rack01:neptune&#34;&#xA;]&#xA;```&#xA;</description>
          <default>
            <item>name rack:location</item>
          </default>
          <required>false</required>
        </input>
        <input>
          <name>input-with-pipe</name>
          <type>string</type>
          <description>It includes v1 | v2 | v3</description>
          <default>v1</default>
          <required>false</required>
        </input>
        <input>
          <name>input_with_underscores</name>
          <type>any</type>
          <description>A variable with underscores.</description>
          <default xsi:nil="true"></default>
          <required>true</required>
        </input>
        <input>
          <name>list-1</name>
          <type>list</type>
          <description>It&#39;s list number one.</description>
          <default>
            <item>a</item>
            <item>b</item>
            <item>c</item>
          </default>
          <required>false</required>
        </input>
        <input>
          <name>list-2</name>
          <type>list</type>
          <description>It&#39;s list number two.</description>
          <default xsi:nil="true"></default>
          <required>true</required>
        </input>
        <input>
          <name>list-3</name>
          <type>list</type>
          <description xsi:nil="true"></description>
          <default></default>
          <required>false</required>
        </input>
        <input>
          <name>list_default_empty</name>
          <type>list(string)</type>
          <description xsi:nil="true"></description>
          <default></default>
          <required>false</required>
        </input>
        <input>
          <name>long_type</name>
          <type>object({&#xA;    name = string,&#xA;    foo  = object({ foo = string, bar = string }),&#xA;    bar  = object({ foo = string, bar = string }),&#xA;    fizz = list(string),&#xA;    buzz = list(string)&#xA;  })</type>
          <description>This description is itself markdown.&#xA;&#xA;It spans over multiple lines.&#xA;</description>
          <default>
            <bar>
              <bar>bar</bar>
              <foo>bar</foo>
            </bar>
            <buzz>
              <item>fizz</item>
              <item>buzz</item>
            </buzz>
            <fizz></fizz>
            <foo>
              <bar>foo</bar>
              <foo>foo</foo>
            </foo>
            <name>hello</name>
          </default>
          <required>false</required>
        </input>
        <input>
          <name>map-1</name>
          <type>map</type>
          <description>It&#39;s map number one.</description>
          <default>
            <a>1</a>
            <b>2</b>
            <c>3</c>
          </default>
          <required>false</required>
        </input>
        <input>
          <name>map-2</name>
          <type>map</type>
          <description>It&#39;s map number two.</description>
          <default xsi:nil="true"></default>
          <required>true</required>
        </input>
        <input>
          <name>map-3</name>
          <type>map</type>
          <description xsi:nil="true"></description>
          <default></default>
          <required>false</required>
        </input>
        <input>
          <name>no-escape-default-value</name>
          <type>string</type>
          <description>The description contains `something_with_underscore`. Defaults to &#39;VALUE_WITH_UNDERSCORE&#39;.</description>
          <default>VALUE_WITH_UNDERSCORE</default>
          <required>false</required>
        </input>
        <input>
          <name>number-1</name>
          <type>number</type>
          <description>It&#39;s number number one.</description>
          <default>42</default>
          <required>false</required>
        </input>
        <input>
          <name>number-2</name>
          <type>number</type>
          <description>It&#39;s number number two.</description>
          <default xsi:nil="true"></default>
          <required>true</required>
        </input>
        <input>
          <name>number-3</name>
          <type>number</type>
          <description xsi:nil="true"></description>
          <default>19</default>
          <required>false</required>
        </input>
        <input>
          <name>number-4</name>
          <type>number</type>
          <description xsi:nil="true"></description>
          <default>15.75</default>
          <required>false</required>
        </input>
        <input>
          <name>number_default_zero</name>
          <type>number</type>
          <description xsi:nil="true"></description>
          <default>0</default>
          <required>false</required>
        </input>
        <input>
          <name>object_default_empty</name>
          <type>object({})</type>
          <description xsi:nil="true"></description>
          <default></default>
          <required>false</required>
        </input>
        <input>
          <name>string-1</name>
          <type>string</type>
          <description>It&#39;s string number one.</description>
          <default>bar</default>
          <required>false</required>
        </input>
        <input>
          <name>string-2</name>
          <type>string</type>
          <description>It&#39;s string number two.</description>
          <default xsi:nil="true"></default>
          <required>true</required>
        </input>
        <input>
          <name>string-3</name>
          <type>string</type>
          <description xsi:nil="true"></description>
          <default></default>
          <required>false</required>
        </input>
        <input>
          <name>string-special-chars</name>
          <type>string</type>
          <description xsi:nil="true"></description>
          <default>\.&lt;&gt;[]{}_-</default>
          <required>false</required>
        </input>
        <input>
          <name>string_default_empty</name>
          <type>string</type>
          <description xsi:nil="true"></description>
          <default></default>
          <required>false</required>
        </input>
        <input>
          <name>string_default_null</name>
          <type>string</type>
          <description xsi:nil="true"></description>
          <default xsi:nil="true"></default>
          <required>false</required>
        </input>
        <input>
          <name>string_no_default</name>
          <type>string</type>
          <description xsi:nil="true"></description>
          <default xsi:nil="true"></default>
          <required>true</required>
        </input>
        <input>
          <name>unquoted</name>
          <type>any</type>
          <description xsi:nil="true"></description>
          <default xsi:nil="true"></default>
          <required>true</required>
        </input>
        <input>
          <name>with-url</name>
          <type>string</type>
          <description>The description contains url. https://www.domain.com/foo/bar_baz.html</description>
          <default></default>
          <required>false</required>
        </input>
      </inputs>
      <outputs>
        <output>
          <name>output-0.12</name>
          <description>terraform 0.12 only</description>
        </output>
        <output>
          <name>output-1</name>
          <description>It&#39;s output number one.</description>
        </output>
        <output>
          <name>output-2</name>
          <description>It&#39;s output number two.</description>
        </output>
        <output>
          <name>unquoted</name>
          <description>It&#39;s unquoted output.</description>
        </output>
      </outputs>
      <providers>
        <provider>
          <name>aws</name>
          <alias xsi:nil="true"></alias>
          <version>&gt;= 2.15.0</version>
        </provider>
        <provider>
          <name>aws</name>
          <alias>ident</alias>
          <version>&gt;= 2.15.0</version>
        </provider>
        <provider>
          <name>null</name>
          <alias xsi:nil="true"></alias>
          <version xsi:nil="true"></version>
        </provider>
        <provider>
          <name>tls</name>
          <alias xsi:nil="true"></alias>
          <version xsi:nil="true"></version>
        </provider>
      </providers>
      <requirements>
        <requirement>
          <name>terraform</name>
          <version>&gt;= 0.12</version>
        </requirement>
        <requirement>
          <name>aws</name>
          <version>&gt;= 2.15.0</version>
        </requirement>
        <requirement>
          <name>random</name>
          <version>&gt;= 2.2.0</version>
        </requirement>
      </requirements>
    </module>

###### Auto generated by spf13/cobra on 25-Nov-2020
