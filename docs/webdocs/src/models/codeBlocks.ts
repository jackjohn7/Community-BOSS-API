export type CodeDataTab = {
  lang: string,

}

export class CodeBlockLanguage {
  url: string;
  lang: string;
  constructor(lang: string, url: string) {
    this.lang = lang;
    this.url = url;
  }

  public formattedCode(uri: string) {
    return this.url.replace("[REQUEST_URI]", uri);
  }
};

