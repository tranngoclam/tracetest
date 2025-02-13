import Span from 'models/Span.model';

export type TSpanFlatAttribute = {
  key: string;
  value: string;
};

export interface ISpanState {
  focusedSpan: string;
  matchedSpans: string[];
  selectedSpan?: Span;
}

export type TSpansResult = Record<
  string,
  {
    failed: number;
    passed: number;
  }
>;
