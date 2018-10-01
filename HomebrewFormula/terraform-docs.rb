class TerraformDocs < Formula
    desc "Tool to generate documentation from Terraform modules"
    homepage "https://github.com/segmentio/terraform-docs"
    url "https://github.com/segmentio/terraform-docs", :tag => "v0.4.0"
    sha256 "dc52e1701e2858ed6dd19bd7ef308241c0bb3dd3060b7da056b9e2c64883f96f"
  
    depends_on "go" => :build
  
    def install
      ENV["GOPATH"] = buildpath
      dir = buildpath/"src/github.com/segmentio/terraform-docs"
      dir.install buildpath.children
  
      cd dir do
        system "make", "build-darwin-amd64"
        bin.install "build/darwin-amd64/terraform-docs"
        prefix.install_metafiles
      end
    end
  
    test do
      system "#{bin}/terraform-docs", "--version"
    end
  end
  