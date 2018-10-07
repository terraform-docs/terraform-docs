class TerraformDocs < Formula
  desc "Tool to generate documentation from Terraform modules"
  homepage "https://github.com/segmentio/terraform-docs"
  url "https://github.com/segmentio/terraform-docs/archive/v0.4.5.tar.gz"
  sha256 "78b75d3ba2525b272ae93036465bf3a5137bb259d7f7c106d026529492df2e29"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    dir = buildpath/"src/github.com/segmentio/terraform-docs"
    dir.install buildpath.children

    cd dir do
      system "make", "build-darwin-amd64"
      bin.install "bin/darwin-amd64/terraform-docs"
      prefix.install_metafiles
    end
  end

  test do
    system "#{bin}/terraform-docs", "--version"
  end
end
