/*
Copyright 2023 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import styled from 'styled-components';

import teleport from 'design/assets/images/icons/teleport.png';

export const TooltipLogosSpacer = styled.div`
  padding: 0 8px;
`;

export const TeleportIcon = styled.div`
  background: url(${teleport}) no-repeat;
  width: 30px;
  height: 30px;
  background-size: contain;
`;

export const TooltipLogos = styled.div`
  display: flex;
  align-items: center;
`;

export const TooltipFooter = styled.div`
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

export const Tooltip = styled.div`
  position: absolute;
  z-index: 100;
  top: 150px;
  left: 210px;
  background: ${({ theme }) => theme.colors.levels.popout};
  border-radius: 5px;
  width: 270px;
  font-size: 15px;
  padding: 20px 20px 15px;
  display: flex;
  flex-direction: column;

  &:after {
    content: '';
    position: absolute;
    width: 0;
    height: 0;
    border-style: solid;
    border-width: 10px 10px 10px 0;
    border-color: transparent ${({ theme }) => theme.colors.levels.popout}
      transparent transparent;
    left: -10px;
    top: 20px;
  }
`;

export const TooltipTitle = styled.div`
  font-size: 18px;
  font-weight: bold;
  border-radius: 5px;
  margin-bottom: 15px;
`;

export const TooltipTitleBackground = styled.span`
  background: linear-gradient(-45deg, #ee7752, #e73c7e);
  padding: 5px;
  border-radius: 5px;
  color: white;
`;

export const TooltipButton = styled.div`
  cursor: pointer;
  display: inline-flex;
  border: 1px solid ${({ theme }) => theme.colors.text.slightlyMuted};
  border-radius: 5px;
  padding: 8px 15px;
`;
